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

// TEMP BEARER TOKEN

var envErr error = godotenv.Load()
var bearerToken string = os.Getenv("BEARER_TOKEN")

// Users
var users []rem.UserResponse = nil

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

func newUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if isRequestAuthorised(&w, r) {
		if r.Method == http.MethodPost {
			r.ParseForm()

			salt := rem.GenSalt()
			passwordHash := rem.HashPassword(r.Form.Get("password"), salt)

			newUser := rem.UserResponse{
				Id:       rem.GenUserId(users),
				Username: r.Form.Get("username"),
				Email:    r.Form.Get("email"),
				Hash:     passwordHash,
				Salt:     string(salt[:]),
			}

			userJson, err := json.Marshal(newUser)

			if err != nil {
				panic(err)
			}

			users = append(users, newUser)
			fmt.Fprintf(w, "%s", userJson[:])
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
 * MAIN FUNCTION
 */

func main() {
	if envErr != nil {
		panic(envErr)
	}

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/user/new", newUserHandler)
	http.HandleFunc("/user/auth", authUserHandler)

	fmt.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
