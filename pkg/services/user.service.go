package services

import (
	"fmt"
	"saarm/modules/pg"
	"saarm/pkg/helpers"
	modelRequest "saarm/pkg/models/request"
	"saarm/pkg/utilities"

	"errors"
	"saarm/pkg/common"
	"saarm/pkg/models"
	modelResponse "saarm/pkg/models/response"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

func IsExistedUser(user modelRequest.SignUpRequest) bool {
	var count int
	pg.DB.Raw("select count(*) from users where username = ?", user.Username).Scan(&count)

	return count > 0
}

func CreateUser(user modelRequest.SignUpRequest) (modelResponse.SignUpResponse, error) {
	tx := pg.DB.Begin()

	newUser := models.User{Email: user.Email, Password: helpers.HashPassword(user.Password), Username: user.Username}

	result := tx.Clauses(clause.Returning{Columns: []clause.Column{{Name: "id"}}}).Create(&newUser)

	if result.Error != nil {
		tx.Rollback()

		return modelResponse.SignUpResponse{}, errors.New(result.Error.Error())
	}

	var roleID string

	pg.DB.Raw("SELECT r.id FROM roles r WHERE r.name = ?", common.GUEST_ROLE).Scan(&roleID)

	//---------- Assign Role for user------------
	assignRoleUser := models.UserRole{UserID: newUser.ID, RoleID: utilities.ParseStringToUuid(roleID)}

	assignRoleUserErr := tx.Create(&assignRoleUser).Error

	if assignRoleUserErr != nil {
		tx.Rollback()

		return modelResponse.SignUpResponse{}, assignRoleUserErr
	}

	tx.Commit()

	return modelResponse.SignUpResponse{LastLoginAt: time.Now(), UserID: newUser.ID}, nil
}

func GetUsers(c echo.Context) error {
	return utilities.R200(c, "users")
}

func GetUserByID(id uuid.UUID) (modelResponse.UserResponse, error) {

	var user modelResponse.UserResponse

	q := fmt.Sprintf(`
  SELECT u.id, u.last_login_at, u.status, a.name "ApartmentName", a.address
  FROM users u
  INNER JOIN apartments a ON a.user_id = u.id and u.id = '%s'`, id)

	err := pg.DB.Raw(q).Scan(&user)

	if err.Error != nil {
		return modelResponse.UserResponse{}, err.Error
	}

	return user, nil
}

func PatchUser(userID uuid.UUID, req modelRequest.UpdateUserRequest) error {
	err := pg.DB.Exec("UPDATE user_roles SET role_id = ? WHERE user_id = ?", req.RoleName, userID)

	return err.Error
}

func PutUsers(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
		"data":    "users",
	})
}

func DeleteUserByID(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
		"data":    "users",
	})
}

func DeleteUsers(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
		"data":    "users",
	})
}
