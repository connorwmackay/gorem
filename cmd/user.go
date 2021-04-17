package cmd

import (
	"fmt"
)

func GenUserId(users []UserResponse) string {
	if len(users) == 0 {
		return fmt.Sprintf("%0x", 0)
	} else {
		return fmt.Sprintf("%0x", len(users))
	}
}
