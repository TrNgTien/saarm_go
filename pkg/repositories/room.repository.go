package repositories

import (
	"gorm.io/gorm"

	"saarm/pkg/common"

	"saarm/pkg/models"
)

type RoomRepository interface {
	FindAllUsers(paginationQuery common.PaginationQuery) []models.UserResponse
	FindUserByID(id int) models.UserResponse
	FindUserByEmail(email string) models.User
	UpdateUserByID(id string, user models.User) models.User
	DeleteUserByID(id string)
}

func RoomRepo(db *gorm.DB) UserRepository {
	return &userRepository{
		Db: db,
	}
}

func (r *userRepository) Find(paginationQuery common.PaginationQuery) []models.UserResponse {
	var users []models.UserResponse

	limit, offset, page := paginationQuery.Limit, paginationQuery.Offset, paginationQuery.Page

	switch {
	case limit == 0:
		{
			r.Db.Limit(50).Offset(0).Find(&users)
		}

	case page > 0:
		{
			r.Db.Limit(limit).Offset(page * limit).Find(&users)
		}

	default:
		{
			r.Db.Model(&models.User{}).Limit(limit).Offset(offset).Find(&users)
		}
	}

	return users

}
