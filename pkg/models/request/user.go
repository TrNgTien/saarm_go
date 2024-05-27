package modelRequest

type (
	NewUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Status   string `json:"status"`
	}

	SignInRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	SignUpRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	UpdateUserRequest struct {
		RoleName string `json:"roleName"`
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Status   string `json:"status"`
	}
)
