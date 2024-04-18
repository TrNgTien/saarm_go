package repositories

import (
	"gorm.io/gorm"

	"saarm/pkg/common"
	"saarm/pkg/models"
)

type AparmentRepository interface {
	FindAllAparments(paginationQuery common.PaginationQuery) []models.AparmentResponse
	FindAparmentByID(id int) models.AparmentResponse
	FindAparmentByEmail(email string) models.Apartment
	UpdateAparmentByID(id string, user models.Apartment) models.Apartment
	DeleteAparmentByID(id string)
}

func AparmentRepo(db *gorm.DB) AparmentRepository {
	return &userRepository{
		Db: db,
	}
}

func (r *userRepository) FindAllAparments(paginationQuery common.PaginationQuery) []models.AparmentResponse {
	var users []models.AparmentResponse

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
			r.Db.Model(&models.Apartment{}).Limit(limit).Offset(offset).Find(&users)
		}
	}

	return users

}

func (r *userRepository) FindAparmentByID(id int) models.AparmentResponse {
	var user models.AparmentResponse
	r.Db.Find(&user, "id = ?", id)
	return user
}

func (r *userRepository) FindAparmentByEmail(email string) models.Apartment {
	var user models.Apartment
	r.Db.Find(&user, "email = ?", email)
	return user
}

func (r *userRepository) UpdateAparmentByID(id string, user models.Apartment) models.Apartment {
	var oldData models.Apartment
	r.Db.Find(&oldData, "id = ?", id)
	if oldData.ID == user.ID {
		r.Db.Save(&user)
	}
	return user
}

func (r *userRepository) DeleteAparmentByID(id string) {
	var user models.Apartment
	r.Db.Delete(&user, "id = ?", id)
}
