package response

import (
	"strings"
	"fmt"
)

const (
	MessageNotFound = "Message not found."
	MessageRemoved 	= "Message removed."
	MessageUpdated  = "Message updated."
	MessageSaved 	= "Message saved successfully."
	UserNotFound    = "User not found."
	AccessDenied 	= "Access denied."
	UserCreated		= "User created."
	UserRemoved		= "User removed."
)

func Marshal(action string, params []string) string {
	return fmt.Sprintf("%s:%s", action, strings.Join(params, ":"))
}

