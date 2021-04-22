package cmd

import (
	"fmt"
)

func GenPostId(posts []PostResponse) string {
	if len(posts) == 0 {
		return fmt.Sprintf("%0x", 0)
	} else {
		return fmt.Sprintf("%0x", len(posts))
	}
}

func IsValidPostTitle(title string) bool {
	if len(title) > 0 && len(title) < 255 {
		return true
	} else {
		return false
	}
}

func IsValidPostAuthor(currentSession UserSessionResponse) bool {
	// TODO: Make sure there is a logged in user
	if currentSession.LoginStatus {
		return true
	} else {
		return false
	}
}

func IsValidPostContent(content string) bool {
	// Abritary max val.
	// TODO: Content should be in Markdown
	if len(content) > 0 && len(content) < 10000 {
		return true
	} else {
		return false
	}
}
