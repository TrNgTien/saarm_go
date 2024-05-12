package services

import (
	"saarm/modules/pg"
	"saarm/pkg/models"
	modelRequest "saarm/pkg/models/request"
)

func CreateRole(user modelRequest.NewRole) error {

	newRole := models.Role{Name: user.Name}

	if err := pg.DB.Create(&newRole); err != nil {
		return err.Error
	}

	return nil
}

func GetRoles() error {
	return nil
}
