package services

import (
	"errors"
	"fmt"
	"saarm/modules/pg"
	"saarm/pkg/common"
	"saarm/pkg/dto"
	"saarm/pkg/helpers"
	modelRequest "saarm/pkg/models/request"
	modelResponse "saarm/pkg/models/response"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func comparePassword(reqPass string, hashedPass string) bool {
	return helpers.ValidatePassword(reqPass, hashedPass)
}

func SignIn(user modelRequest.SignInRequest) (modelResponse.AuthResponse, error) {
	var userData dto.UserDtoData
	var userRole string

	if user.Username == "" {
		return modelResponse.AuthResponse{}, errors.New("[SignIn] Username cannot be empty")
	}

	pg.DB.Raw("SELECT id, username, password, last_login_at FROM users WHERE username = ?", user.Username).Scan(&userData)

	if userData.Username == "" {
		return modelResponse.AuthResponse{}, errors.New("[SignIn] Error fetching user")
	}

	isMatchPass := comparePassword(user.Password, userData.Password)

	if !isMatchPass {
		return modelResponse.AuthResponse{}, errors.New("[SignIn] Incorrect password")
	}

	pg.DB.Exec("UPDATE users SET last_login_at = ? WHERE id = ?", time.Now(), userData.ID)

	pg.DB.Raw("SELECT r.name FROM user_roles ur INNER JOIN roles r ON r.id = ur.role_id AND ur.user_id = ?", userData.ID).Scan(&userRole)

	token, err := helpers.GenerateToken(userData.ID, userRole)

	if err != nil {
		return modelResponse.AuthResponse{}, err
	}

	return modelResponse.AuthResponse{Type: common.JwtBearer, Value: token, LastLoginAt: userData.LastLoginAt}, nil
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
