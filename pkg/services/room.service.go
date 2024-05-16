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

func saveFileSystem(file, roomID string) (string, error) {
	baseData := file[strings.IndexByte(file, ',')+1:]
	var outputFileName string

	decodedBase64, err := base64.StdEncoding.DecodeString(baseData)

	if err != nil {
		return outputFileName, err
	}

	parts := strings.SplitN(file, ";", 2)
	var fileType string

	if len(parts) != 2 {
		return outputFileName, err
	}

	mimeType := parts[0]
	fileType = strings.Split(mimeType, "/")[1]

	timestamp := helpers.GetCurrentTimestampString()
	outputFileName = "room-" + roomID + "-" + timestamp + "." + fileType

	pathData := utilities.GetFilePath(common.WATER_METER_PATH, outputFileName)

	f, err := os.Create(pathData)

	if err != nil {
		return outputFileName, err
	}

	defer f.Close()

	if _, err := f.Write(decodedBase64); err != nil {
		return outputFileName, err
	}

	if err := f.Sync(); err != nil {
		return outputFileName, err
	}

	fmt.Println("[storeWaterMeterFile] | Created system file")

	return outputFileName, nil
}

func SubmitWaterMeter(file common.UploadWaterMeter, roomID string) ([]string, error) {
	var numbersDetected []string

	fileCropped, err := saveFileSystem(file.CroppedFile, roomID)

	if err != nil {
    fmt.Println("[SubmitWaterMeter] cropped ", err.Error())
		return numbersDetected, err
	}

	fileOriginal, err := saveFileSystem(file.OriginalFile, roomID)

	if err != nil {
		return numbersDetected, err
	}

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

func GetCurrentBill() error {
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
