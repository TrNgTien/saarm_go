package services

import (
	"errors"
	"fmt"
	"saarm/modules/pg"
	"saarm/pkg/common"
	"saarm/pkg/helpers"
	"time"

	"saarm/pkg/models"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type SignUpResponse struct {
	id       uuid.UUID
	email    string
	username string
}

func IsExistedUser(user common.UserDTO) bool {
	var count int
	row := pg.DB.Raw("select count(*) from users where username = ?", user.Username).Row()

	row.Scan(&count)

	return count > 0
}

func comparePassword(reqPass string, hashedPass string) bool {

	return helpers.ValidatePassword(reqPass, hashedPass)

}

type UserData struct {
	Id          uuid.UUID
	Password    string
	Username    string
	Email       string
	LastLoginAt time.Time
}

func SignIn(user common.UserDTO) (models.UserResponse, error) {
	var userData UserData

	if user.Username == "" {
		return models.UserResponse{}, errors.New("[SignIn] Username cannot be empty")
	}

	pg.DB.Raw("SELECT id, username, email, password, last_login_at FROM users WHERE username = ? LIMIT 1", user.Username).Scan(&userData)

	fmt.Println("[SignIn]", userData)

	if userData.Username == "" {
		return models.UserResponse{}, errors.New("[SignIn] Error fetching user")
	}

	isMatchPass := comparePassword(user.Password, userData.Password)

	if !isMatchPass {
		return models.UserResponse{}, errors.New("[SignIn] Incorrect password")
	}

	pg.DB.Exec("UPDATE users SET last_login_at = ? WHERE id = ?", time.Now(), userData.Id).Row()

	token, err := helpers.GenerateToken(userData.Id)

	if err != nil {

		return models.UserResponse{}, err
	}

	return models.UserResponse{Type: "Bearer", Value: token}, nil
}

func SignUp(user common.UserDTO) (models.UserResponse, error) {

	newUser := models.User{Email: user.Email, Password: helpers.HashPassword(user.Password), Username: user.Username}

	err := pg.DB.Create(&newUser).Error

	if err != nil {
		return models.UserResponse{}, errors.New(err.Error())
	}

	token, err := helpers.GenerateToken(newUser.ID)

	if err != nil {

		return models.UserResponse{}, err
	}

	return models.UserResponse{Type: "Bearer", Value: token}, nil
}

func SignUpWithGoogle() {

	// Developer Console (https://console.developers.google.com).
	conf := &oauth2.Config{
		ClientID:     "YOUR_CLIENT_ID",
		ClientSecret: "YOUR_CLIENT_SECRET",
		RedirectURL:  "YOUR_REDIRECT_URL",
		Scopes: []string{
			"https://www.googleapis.com/auth/bigquery",
			"https://www.googleapis.com/auth/blogger",
		},
		Endpoint: google.Endpoint,
	}
	// Redirect user to Google's consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state")
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// Handle the exchange code to initiate a transport.
	// tok, err := conf.Exchange(oauth2.NoContext, "authorization-code")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// client := conf.Client(oauth2.To, tok)
	// client.Get("...")
}
