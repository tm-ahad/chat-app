package sa

import "strings"

func Marshal(sa []string) string {
	return strings.Join(sa, "\n")
}
