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

func SubmitWaterMeter(file common.UploadWaterMeter, roomID string) ([]string, error) {
	var numbersDetected []string
	fileChan := make(chan string)

	go saveFileSystem(file.CroppedFile, roomID, fileChan)
	go saveFileSystem(file.OriginalFile, roomID, fileChan)

	fileCropped, fileOriginal := <-fileChan, <-fileChan

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

func GetRoomByID(roomID uuid.UUID) error {
	var room modelResponse.AparmentResponse

	pg.DB.Raw("SELECT * FROM rooms WHERE id = ?", roomID).Scan(&room)
	return nil
}

func GetBills() error {
	return nil
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

// func SubmitWaterMeterRecord () error{
//   var waterMeterNumber string
//   return nil
// }
