package helpers

import "strings"

func KeysAndValues(s string) ([]string, []string) {
	values := make([]string, 0)
	keys   := make([]string, 0)
	split  := strings.Split(s, "\n")
	
	for _, line := range split {
		if len(line) == 0 {continue}

		kv := strings.Split(line, ":")

		key     := kv[0]
		values  = append(values, kv[1])
		keys    = append(keys, key)
	}

	return keys, values
}
