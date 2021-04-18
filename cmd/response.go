package cmd

/*
 * JSON Responses
 */

type UserResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Hash     string `json:"hash"`
	Salt     string `json:"salt"`
}

type UserCreationResponse struct {
	UsernameStatus bool `json:"isValidUsermame"`
	EmailStatus    bool `json:"isValidEmail"`
	PasswordStatus bool `json:"isValidPassword"`
}

type UserAuthResponse struct {
	IsAuth bool `json:"isAuthenticated"`
}

type APIAuthResponse struct {
	RequestStatus string `json:"requestStatus"`
}
