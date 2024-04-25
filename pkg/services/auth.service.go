package services

import (
	"errors"
	"fmt"
	"saarm/modules/pg"
	"saarm/pkg/common"
	"saarm/pkg/dto"
	"saarm/pkg/helpers"
	"saarm/pkg/models"
	modelRequests "saarm/pkg/models/request"
	modelReponses "saarm/pkg/models/response"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func IsExistedUser(user modelRequests.SignUpRequest) bool {
	var count int
	pg.DB.Raw("select count(*) from users where username = ?", user.Username).Scan(&count)

	return count > 0
}

func comparePassword(reqPass string, hashedPass string) bool {
	return helpers.ValidatePassword(reqPass, hashedPass)
}

func SignIn(user modelRequests.SignInRequest) (modelReponses.AuthResponse, error) {
	var userData dto.UserData

	if user.Username == "" {
		return modelReponses.AuthResponse{}, errors.New("[SignIn] Username cannot be empty")
	}

	pg.DB.Raw("SELECT id, username, password, last_login_at FROM users WHERE username = ?", user.Username).Scan(&userData)

	if userData.Username == "" {
		return modelReponses.AuthResponse{}, errors.New("[SignIn] Error fetching user")
	}

	isMatchPass := comparePassword(user.Password, userData.Password)

	if !isMatchPass {
		return modelReponses.AuthResponse{}, errors.New("[SignIn] Incorrect password")
	}

	pg.DB.Exec("UPDATE users SET last_login_at = ? WHERE id = ?", time.Now(), userData.ID)

	token, err := helpers.GenerateToken(userData.ID)

	if err != nil {

		return modelReponses.AuthResponse{}, err
	}

	return modelReponses.AuthResponse{Type: common.JwtBearer, Value: token, LastLoginAt: userData.LastLoginAt}, nil
}

func SignUp(user modelRequests.SignUpRequest) (modelReponses.AuthResponse, error) {

	newUser := models.User{Email: user.Email, Password: helpers.HashPassword(user.Password), Username: user.Username}

	err := pg.DB.Create(&newUser).Error

	if err != nil {
		return modelReponses.AuthResponse{}, errors.New(err.Error())
	}

	token, err := helpers.GenerateToken(newUser.ID)

	if err != nil {

		return modelReponses.AuthResponse{}, err
	}

  return modelReponses.AuthResponse{Type: common.JwtBearer, Value: token, LastLoginAt: newUser.LastLoginAt}, nil
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
