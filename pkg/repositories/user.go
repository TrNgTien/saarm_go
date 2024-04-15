package repository

import (
	"gorm.io/gorm"

	"saarm/pkg/models"
)

type UserRepository interface {
	FindAllUsers() []models.User
	FindUserByID(id string) models.User
	FindUserByEmail(email string) models.User
	UpdateUserByID(id string, user models.User) models.User
	DeleteUserByID(id string)
}

type userRepository struct {
	Db *gorm.DB
}

func UserRepo(db *gorm.DB) UserRepository {
	return &userRepository{
		Db: db,
	}
}

func (r *userRepository) FindAllUsers() []models.User {
	var users []models.User
	r.Db.Find(&users)
	return users
}

func (r *userRepository) FindUserByID(id string) models.User {
	var user models.User
	r.Db.Find(&user, "id = ?", id)
	return user
}

func (r *userRepository) FindUserByEmail(email string) models.User {
	var user models.User
	r.Db.Find(&user, "email = ?", email)
	return user
}

func (r *userRepository) UpdateUserByID(id string, user models.User) models.User {
	var oldData models.User
	r.Db.Find(&oldData, "id = ?", id)
	if oldData.ID == user.ID {
		r.Db.Save(&user)
	}
	return user
}

func (r *userRepository) DeleteUserByID(id string) {
	var user models.User
	r.Db.Delete(&user, "id = ?", id)
}
