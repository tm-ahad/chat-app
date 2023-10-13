package helpers

import "strings"

func KeysAndValues(s string) ([]string, []string) {
	values := make([]string, 2)
	keys   := make([]string, 2)
	split  := strings.Split(s, "\n")
	
	for _, line := range split {
		if len(line) == 0 {continue}

		i := strings.Index(line, ":")

		key     := line[:i]
		values   = append(values, line[i+1:])
		keys     = append(keys, key)
	}

	return keys, values
}
