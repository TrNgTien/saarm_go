package services

import (
	"errors"
	"saarm/modules/pg"
	"saarm/pkg/common"
	"saarm/pkg/models"
	modelRequests "saarm/pkg/models/request"
	modelReponses "saarm/pkg/models/response"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateApartments(apartment modelRequests.NewApartment) (modelReponses.AparmentResponse, error) {
	newApartment := models.Apartment{
		Name:          apartment.Name,
		LocationUrl:   apartment.LocationUrl,
		Address:       apartment.Address,
		TotalRoom:     apartment.TotalRoom,
		RoomAvailable: apartment.RoomAvailable,
	}

	err := pg.DB.Create(&newApartment).Error

	if err != nil {
		return modelReponses.AparmentResponse{}, errors.New(err.Error())
	}

	return modelReponses.AparmentResponse{
		ID:            newApartment.ID,
		Name:          newApartment.Name,
		LocationUrl:   newApartment.LocationUrl,
		Address:       newApartment.Address,
		TotalRoom:     newApartment.TotalRoom,
		RoomAvailable: newApartment.RoomAvailable,
	}, nil
}

func LinkUserApartment(linkReq modelRequests.LinkUser) error {

	newUserAparment := models.UserApartment{
		UserID:      linkReq.UserID,
		ApartmentID: linkReq.ApartmentID,
	}

	err := pg.DB.Create(&newUserAparment).Error

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func GetApartments(query common.PaginationQuery) ([]modelReponses.AparmentResponse, error) {
	var apartments []modelReponses.AparmentResponse

	rows, err := pg.DB.Raw("SELECT id, name, location_url, address, total_room, room_available FROM apartments ORDER BY created_at DESC LIMIT ? OFFSET ?", query.Limit, query.Offset).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var apartment modelReponses.AparmentResponse

		err := rows.Scan(&apartment.ID, &apartment.Name, &apartment.LocationUrl, &apartment.Address, &apartment.TotalRoom, &apartment.RoomAvailable)

		if err != nil {
			return nil, err // Propagate errors from Scan
		}

		apartments = append(apartments, apartment)
	}
	return apartments, nil

}

func GetAparmentByID(id uuid.UUID) (modelReponses.AparmentResponse, error) {
	var apartment modelReponses.AparmentResponse

	pg.DB.Raw("SELECT id, name, location_url, address, total_room, room_available FROM apartments WHERE id = ?", id).Scan(&apartment)

	return modelReponses.AparmentResponse{ID: apartment.ID, Name: apartment.Name, Address: apartment.Address, LocationUrl: apartment.LocationUrl, TotalRoom: apartment.TotalRoom, RoomAvailable: apartment.RoomAvailable}, nil
}

func PatchAparment(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
		"data":    "users",
	})
}

func PutAparments(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
		"data":    "users",
	})
}

func DeleteAparmentByID(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
		"data":    "users",
	})
}

func DeleteAparments(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
		"data":    "users",
	})
}
