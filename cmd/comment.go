package cmd

import (
	"fmt"
)

func GenCommentId(comments []CommentResponse) string {
	if len(comments) == 0 {
		return fmt.Sprintf("%0x", 0)
	} else {
		return fmt.Sprintf("%0x", len(comments))
	}
}

func IsValidCommentContent(content string) bool {
	if len(content) > 0 && len(content) < 1000 {
		return true
	} else {
		return false
	}
}

func IsValidCommentAuthor(currentSession UserSessionResponse) bool {
	// TODO: Make sure there is a logged in user
	if currentSession.LoginStatus {
		return true
	} else {
		return false
	}
}

func IsValidCommentPostId(posts []PostResponse, id string) bool {
	for i := 0; i < len(posts); i++ {
		if posts[i].Id == id {
			return true
		}
	}

	return false
}
