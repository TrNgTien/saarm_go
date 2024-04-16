package repositories

import (
	"gorm.io/gorm"

	"saarm/pkg/models"
)

type UserRepository interface {
  FindAllUsers(paginationQuery PaginationQuery) []models.UserResponse
  FindUserByID(id int) models.UserResponse
  FindUserByEmail(email string) models.User
  UpdateUserByID(id string, user models.User) models.User
  DeleteUserByID(id string)
}

type PaginationQuery struct {
  Limit int
  Offset int
  Page int
}

type userRepository struct {
  Db *gorm.DB
}

func UserRepo(db *gorm.DB) UserRepository {
  return &userRepository{
    Db: db,
  }
}

func (r *userRepository) FindAllUsers(paginationQuery PaginationQuery) []models.UserResponse {
  var users []models.UserResponse

  limit, offset, page := paginationQuery.Limit, paginationQuery.Offset, paginationQuery.Page

  switch {
    case limit == 0: {
      r.Db.Limit(50).Offset(0).Find(&users)
    }

    case page > 0: {
      r.Db.Limit(limit).Offset(page * limit).Find(&users)
    }

    default: {
      r.Db.Model(&models.User{}).Limit(limit).Offset(offset).Find(&users)
    }
  }

  return users

}

func (r *userRepository) FindUserByID(id int) models.UserResponse {
  var user models.UserResponse
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
