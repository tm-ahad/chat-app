package response

import (
	"strings"
	"fmt"
)

const (
	MessageNotFound = "Message not found."
	MessageRemoved 	= "Message removed."
	MessageUpdated  = "Message updated."
	InvalidAction   = "Invalid Action."
	MessageSaved 	= "Message saved successfully."
	UserNotFound    = "User not found."
	AccessDenied 	= "Access denied."
	UserCreated		= "User created."
	UserRemoved		= "User removed."
	UserExists      = "User exists."
	UserSaved 		= "User saved."
	Empty           = ""
)

func Marshal(action string, params []string) string {
	return fmt.Sprintf("%s:%s", action, strings.Join(params, ":"))
}

