package cmd

/*
 * User Responses
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

type UserSessionResponse struct {
	Id          string `json:"id"`
	LoginStatus bool   `json:"isLoggedIn"`
}

/*
 * Post Responses
 */

type PostCreationResponse struct {
	TitleStatus      bool `json:"isValidTitle"`
	ContentStatus    bool `json:"isValidContent"`
	AuthorAuthStatus bool `json:"isValidAuthor"`
}

type PostResponse struct {
	Id       string            `json:"id"`
	AuthorId string            `json:"authorId"`
	Title    string            `json:"title"`
	Content  string            `json:"content"`
	Comments []CommentResponse `json:"comments"`
}

/*
 * Comment Responses
 */

type CommentCreationResponse struct {
	AuthorAuthStatus bool `json:"isValidAuthor"`
	PostIdStatus     bool `json:"isValidPost"`
	ContentStatus    bool `json:"isValidContent"`
}

type CommentResponse struct {
	Id       string `json:"id"`
	PostId   string `json:"postId"`
	AuthorId string `json:"authorId"`
	Content  string `json:"content"`
}

/*
 * System Responses
 */

type APIAuthResponse struct {
	RequestStatus string `json:"requestStatus"`
}
