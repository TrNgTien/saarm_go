package services

import (
	"errors"
	"fmt"
	"saarm/modules/pg"
	"saarm/pkg/models"
	modelRequest "saarm/pkg/models/request"
	modelReponses "saarm/pkg/models/response"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateApartments(apartment modelRequest.NewApartment) (modelReponses.AparmentResponse, error) {
	newApartment := models.Apartment{
		Name:          apartment.Name,
		LocationUrl:   apartment.LocationUrl,
		Address:       apartment.Address,
		TotalRoom:     apartment.TotalRoom,
		RoomAvailable: apartment.RoomAvailable,
		UserID:        apartment.UserID,
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

func GetApartments() ([]modelReponses.AparmentResponse, error) {
	var apartments []modelReponses.AparmentResponse

	q := "SELECT a.id, a.name, a.address, a.total_room, a.room_available FROM  apartments a"

	rows, err := pg.DB.Raw(q).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var apartment modelReponses.AparmentResponse

		err := rows.Scan(&apartment.ID, &apartment.Name, &apartment.Address, &apartment.TotalRoom, &apartment.RoomAvailable)

		if err != nil {
			return nil, err
		}

		apartments = append(apartments, apartment)
	}
	return apartments, nil

}

func GetApartmentsByUserID(userID uuid.UUID) ([]modelReponses.AparmentResponse, error) {
	var apartments []modelReponses.AparmentResponse

	q := fmt.Sprintf(`SELECT a.id, a.name, a.address, a.total_room, a.room_available
  FROM apartments a WHERE a.user_id = '%s' ORDER BY a.created_at asc`, userID)

	rows, err := pg.DB.Raw(q).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var apartment modelReponses.AparmentResponse

		err := rows.Scan(&apartment.ID, &apartment.Name, &apartment.Address, &apartment.TotalRoom, &apartment.RoomAvailable)

		if err != nil {
			return nil, err
		}

		apartments = append(apartments, apartment)
	}
	return apartments, nil

}

func GetRoomsByApartmentID(id uuid.UUID) ([]modelReponses.RoomsResponseByApartment, error) {
	var roomsByApartment []modelReponses.RoomsResponseByApartment

	q := "SELECT r.id, r.name, r.status, r.room_price, r.max_people, r.current_people FROM rooms r INNER JOIN apartments a ON r.apartment_id = a.id AND a.id = ? ORDER BY r.name ASC"

	rows, err := pg.DB.Raw(q, id).Rows()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var room modelReponses.RoomsResponseByApartment

		err := rows.Scan(&room.ID, &room.Name, &room.Status, &room.RoomPrice, &room.MaxPeople, &room.CurrentPeople)

		if err != nil {
			return nil, err
		}

		roomsByApartment = append(roomsByApartment, room)
	}

	return roomsByApartment, nil
}

func GetApartmentByID(id uuid.UUID) (modelReponses.AparmentResponse, error) {
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
