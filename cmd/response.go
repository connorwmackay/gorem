package cmd

/*
 * JSON Responses
 */

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Hash     string `json:"hash"`
	Salt     string `json:"salt"`
}

type UserAuthResponse struct {
	IsAuth bool `json:"isAuthenticated"`
}
