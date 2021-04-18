package cmd

import (
	"fmt"
	"strings"
)

func GenUserId(users []UserResponse) string {
	if len(users) == 0 {
		return fmt.Sprintf("%0x", 0)
	} else {
		return fmt.Sprintf("%0x", len(users))
	}
}

func IsValidUsername(username string) bool {
	// TODO: Check to see if username is unique

	if len(username) > 80 || len(username) <= 0 {
		return false
	}

	return true
}

func IsValidEmail(email string) bool {
	// TODO: Improve validation logic

	if len(email) > 80 || len(email) <= 0 {
		return false
	} else if !strings.Contains(email, "@") {
		return false
	} else if !strings.Contains(email, ".com") && !strings.Contains(email, ".co.uk") {
		return false
	}

	return true
}

func IsValidPassword(password string) bool {
	// TODO: Require capital and special character

	if len(password) > 80 || len(password) < 8 {
		return false
	}

	return true
}
