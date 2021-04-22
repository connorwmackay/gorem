package main

import (
	"encoding/json"
	"fmt"
	rem "gorem/cmd"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var (
	envErr      error                   = godotenv.Load()
	bearerToken string                  = os.Getenv("BEARER_TOKEN")
	users       []rem.UserResponse      = nil
	posts       []rem.PostResponse      = nil
	session     rem.UserSessionResponse = rem.UserSessionResponse{
		Id:          "",
		LoginStatus: false,
	}
)

/*
 * HANDLERS
 */

func isRequestAuthorised(w *http.ResponseWriter, r *http.Request) bool {
	if (*r).Header.Get("Authorization") != bearerToken {
		response := rem.APIAuthResponse{RequestStatus: "denied"}
		responseJson, err := json.Marshal(response)

		if err != nil {
			panic(err)
		}

		fmt.Fprintf((*w), string(responseJson))

		return false
	} else {
		return true
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if isRequestAuthorised(&w, r) {
		fmt.Fprintf(w, "{}")
	}
}

/*
 * User Handlers
 */

func newUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if isRequestAuthorised(&w, r) {
		if r.Method == http.MethodPost {
			r.ParseForm()

			userCreationStatus := rem.UserCreationResponse{
				UsernameStatus: rem.IsValidUsername(r.Form.Get("username")),
				EmailStatus:    rem.IsValidEmail(r.Form.Get("email")),
				PasswordStatus: rem.IsValidPassword(r.Form.Get("password")),
			}

			isValidUserCreation := userCreationStatus.UsernameStatus &&
				userCreationStatus.EmailStatus &&
				userCreationStatus.PasswordStatus

			if isValidUserCreation {
				salt := rem.GenSalt()
				passwordHash := rem.HashPassword(r.Form.Get("password"), salt)

				newUser := rem.UserResponse{
					Id:       rem.GenUserId(users),
					Username: r.Form.Get("username"),
					Email:    r.Form.Get("email"),
					Hash:     passwordHash,
					Salt:     string(salt[:]),
				}

				users = append(users, newUser)
			}

			userCreationJson, err := json.Marshal(userCreationStatus)

			if err != nil {
				panic(err)
			}

			fmt.Fprintf(w, string(userCreationJson))
		}
	}
}

func authUserHandler(w http.ResponseWriter, r *http.Request) {
	if isRequestAuthorised(&w, r) {
		if r.Method == http.MethodPost {
			r.ParseForm()

			isAuth := false

			for i := 0; i < len(users); i++ {
				if users[i].Username == r.Form.Get("username") {
					if rem.CheckHashedPasswords(r.Form.Get("password"), users[i].Hash, []byte(users[i].Salt)) {
						isAuth = true
						session.Id = users[i].Id
						session.LoginStatus = true
					}
				}
			}

			response := rem.UserAuthResponse{IsAuth: isAuth}
			responseJson, err := json.Marshal(response)
			if err != nil {
				panic(err)
			}

			fmt.Fprintf(w, string(responseJson))
		}
	}
}

/*
 * Post Handlers
 */
func newPostHandler(w http.ResponseWriter, r *http.Request) {
	if isRequestAuthorised(&w, r) {
		if r.Method == http.MethodPost {
			r.ParseForm()

			postCreationStatus := rem.PostCreationResponse{
				TitleStatus:      rem.IsValidPostTitle(r.Form.Get("title")),
				ContentStatus:    rem.IsValidPostContent(r.Form.Get("content")),
				AuthorAuthStatus: rem.IsValidPostAuthor(session),
			}

			isValidPost := postCreationStatus.TitleStatus &&
				postCreationStatus.ContentStatus &&
				postCreationStatus.AuthorAuthStatus

			if isValidPost {
				newPost := rem.PostResponse{
					Id:       rem.GenPostId(posts),
					AuthorId: session.Id,
					Title:    r.Form.Get("title"),
					Content:  r.Form.Get("content"),
					Comments: []rem.CommentResponse{},
				}

				posts = append(posts, newPost)

				fmt.Println("=New Post=\nPost ID: " + newPost.Id)
			}

			postCreationJson, err := json.Marshal(postCreationStatus)

			if err != nil {
				panic(err)
			}

			fmt.Fprintln(w, string(postCreationJson))
		}
	}
}

// TODO: Add many filter options, allow to get multiple posts...
func getPostHandler(w http.ResponseWriter, r *http.Request) {
	if isRequestAuthorised(&w, r) {
		if r.Method == http.MethodGet {
			query := r.URL.Query()
			postId := query["id"]
			isValidId := false

			for i := 0; i < len(posts); i++ {
				if posts[i].Id == postId[0] {
					postJson, err := json.Marshal(posts[i])
					if err != nil {
						panic(err)
					}

					isValidId = true
					fmt.Fprintf(w, string(postJson))
				}
			}

			if len(postId) == 0 || !isValidId {
				fmt.Fprintf(w, "{}")
			}
		}
	}
}

/*
 * MAIN FUNCTION
 */

func main() {
	if envErr != nil {
		panic(envErr)
	}

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/user/new", newUserHandler)
	http.HandleFunc("/user/auth", authUserHandler)
	http.HandleFunc("/post/new", newPostHandler)
	http.HandleFunc("/post/get", getPostHandler)

	fmt.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
