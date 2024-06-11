package services

import (
	"encoding/base64"
	"fmt"
	"os"
	"saarm/modules/pg"
	"saarm/pkg/common"
	"saarm/pkg/helpers"
	"saarm/pkg/models"
	"saarm/pkg/utilities"
	"strings"

	modelRequest "saarm/pkg/models/request"
	modelResponse "saarm/pkg/models/response"

	"github.com/google/uuid"
)

func saveFileSystem(file, roomID string, fileChan chan string) {
	baseData := file[strings.IndexByte(file, ',')+1:]
	var outputFileName string

	decodedBase64, err := base64.StdEncoding.DecodeString(baseData)

	if err != nil {
		fileChan <- outputFileName
	}

	parts := strings.SplitN(file, ";", 2)

	var fileType string

	if len(parts) != 2 {
		fileChan <- outputFileName
	}

	mimeType := parts[0]
	fileType = strings.Split(mimeType, "/")[1]

	timestamp := helpers.GetCurrentTimestampString()
	outputFileName = "room-" + roomID + "-" + timestamp + "." + fileType

	pathData := utilities.GetFilePath(common.WATER_METER_PATH, outputFileName)

	f, err := os.Create(pathData)

	if err != nil {
		fileChan <- outputFileName
	}

	defer f.Close()

	if _, err := f.Write(decodedBase64); err != nil {
		fileChan <- outputFileName
	}

	if err := f.Sync(); err != nil {
		fileChan <- outputFileName
	}

	fmt.Println("[storeWaterMeterFile] | Created system file")

	fileChan <- outputFileName
}

func DetectWaterMeter(file common.UploadWaterMeter, roomID string) ([]string, error) {
	var numbersDetected []string
	fileChan := make(chan string)

	go saveFileSystem(file.CroppedFile, roomID, fileChan)
	go saveFileSystem(file.OriginalFile, roomID, fileChan)

	fileCropped, fileOriginal := <-fileChan, <-fileChan

	close(fileChan)

	IMAGE_WATER_METER_PATH := utilities.GetFilePath(common.WATER_METER_PATH, fileCropped)
	ORIGINAL_WATER_METER_PATH := utilities.GetFilePath(common.WATER_METER_PATH, fileOriginal)

	info, err := UploadObject(common.MINIO_BUCKET_CROPPED, fileCropped, IMAGE_WATER_METER_PATH)

	if err != nil {
		fmt.Println("Failed to upload cropped image Minio", err.Error())
		return numbersDetected, err
	}

	infoOriginal, err := UploadObject(common.MINIO_BUCKET_ORIGINAL, fileOriginal, ORIGINAL_WATER_METER_PATH)

	if err != nil {
		fmt.Println("Failed to upload original image", err.Error())
		return numbersDetected, err
	}

	fmt.Println("Upload image cropped, originalFile success", info, infoOriginal)

	numbersDetected, err = GetTextDetection(common.WATER_METER_PATH, fileCropped)

	if err != nil {
		fmt.Println("RUNNING detect water meter failed: ", err.Error())
		return numbersDetected, err
	}

	return numbersDetected, nil
}

func CreateRoom(room modelRequest.NewRoom) (modelResponse.RoomResponse, error) {
	tx := pg.DB.Begin()

	newRoom := models.Room{
		Name:          room.Name,
		Password:      helpers.HashPassword(room.Password),
		Username:      room.Username,
		RoomPrice:     room.RoomPrice,
		MaxPeople:     room.MaxPeople,
		CurrentPeople: 0,
		ApartmentID:   room.ApartmentID,
	}

	newRoomErr := tx.Create(&newRoom).Error

	if newRoomErr != nil {
		tx.Rollback()
		return modelResponse.RoomResponse{}, newRoomErr
	}

	tx.Commit()
	return modelResponse.RoomResponse{ID: newRoom.ID, Name: newRoom.Name}, nil
}

func GetRooms() error {
	return nil
}

func GetRoomByID(roomID uuid.UUID) (modelResponse.RoomResponse, error) {
	var room modelResponse.RoomResponse

	err := pg.DB.Raw("SELECT r.id, r.name, r.room_price, r.status, a.name \"ApartmentName\", a.address FROM rooms r INNER JOIN apartments a ON a.id = r.apartment_id AND r.id = ?", roomID).Scan(&room)

	if err.Error != nil {
		return modelResponse.RoomResponse{}, err.Error
	}

	return room, nil
}

func GetBillByRoom(roomID uuid.UUID, monthReq string) (modelResponse.BillByRoomResponse, error) {
	var billRoom modelResponse.BillByRoomResponse

	q := fmt.Sprintf(`SELECT m.id, m.created_at, water_consume, electricity_consume, extra_fee, r.room_price
     FROM monthly_bill_logs as m
     INNER JOIN rooms as r on r.id = m.room_id AND m.room_id = '%s'
     AND m.created_at >= date_trunc('month', timestamp with time zone '%s')
     AND m.created_at < date_trunc('month', timestamp with time zone '%s' + interval '1 month')
     LIMIT 1`,
		roomID, monthReq, monthReq)

	err := pg.DB.Raw(q).Scan(&billRoom)

	if err.Error != nil {
		return modelResponse.BillByRoomResponse{}, err.Error
	}

	return billRoom, nil
}

func GetBills() (modelResponse.RoomResponse, error) {
	return modelResponse.RoomResponse{}, nil
}

func DuplicateRoom(roomID uuid.UUID) (modelResponse.DuplicateRoomResponse, error) {
	tx := pg.DB.Begin()
	room := models.Room{}

	tx.Raw("SELECT max_people, room_price, apartment_id, current_people FROM rooms WHERE id = ?", roomID).Scan(&room)

	duplicatedRoomErr := tx.Create(&room).Error

	if duplicatedRoomErr != nil {

		tx.Rollback()
		return modelResponse.DuplicateRoomResponse{}, duplicatedRoomErr
	}

	tx.Commit()
	return modelResponse.DuplicateRoomResponse{ID: room.ID}, nil
}

func CheckSubmittedWaterMeter(roomID uuid.UUID) (bool, error) {
	var isSubmitted int8

	isSubmittedErr := pg.DB.Raw("SELECT COUNT(*) FROM monthly_bill_logs WHERE room_id = ? AND created_at >= DATE_TRUNC('month', CURRENT_DATE) AND created_at < DATE_TRUNC('month', CURRENT_DATE + interval '1 month')", roomID).Scan(&isSubmitted)

	if isSubmittedErr.Error != nil {

		return false, isSubmittedErr.Error
	}

	return isSubmitted > 0, nil
}

func ConfirmWaterMeter(roomID uuid.UUID, waterMeterNumber string) error {
	var waterNumberLatest string

	err := pg.DB.Raw("SELECT water_number from monthly_bill_logs WHERE room_id = ? ORDER BY created_at DESC LIMIT 1", roomID).Scan(&waterNumberLatest).Error

	if err != nil {
		return err
	}

	oldWater := "0"
	var diffConsume int

	if waterNumberLatest != "" {
		oldWater = waterNumberLatest[:4]
	}

	newWater := waterMeterNumber[:4]

	oldWaterMeter, err := utilities.GetIntValue(oldWater)

	if err != nil {
		return err
	}

	newWaterMeter, err := utilities.GetIntValue(newWater)

	if err != nil {
		return err
	}

	diffConsume = newWaterMeter - oldWaterMeter

	monthlyLogs := models.MonthlyBillLogs{
		RoomID:       roomID,
		WaterNumber:  waterMeterNumber,
		WaterConsume: diffConsume,
	}

	monthlyLogErr := pg.DB.Create(&monthlyLogs).Error

	if monthlyLogErr != nil {
		return monthlyLogErr
	}

	return nil
}

func GetHistorySubmitted(roomID uuid.UUID) ([]modelResponse.HistorySubmitResponse, error) {
	var histories []modelResponse.HistorySubmitResponse

	rows, err := pg.DB.Raw("SELECT id, created_at, water_number, water_consume from monthly_bill_logs where room_id = ? ORDER BY created_at DESC LIMIT 50 OFFSET 0", roomID).Rows()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var log modelResponse.HistorySubmitResponse

		err := rows.Scan(&log.ID, &log.CreatedAt, &log.WaterNumber, &log.WaterConsume)

		if err != nil {
			return nil, err
		}

		histories = append(histories, log)
	}

	return histories, nil
}
