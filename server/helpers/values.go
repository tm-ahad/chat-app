package helpers

import "strings"

func Values(s string) []string {
	var values []string
	split := strings.Split(s, "\n")
	
	for _, line := range split {
		if len(line) == 0 {continue}

		kv := strings.Split(line, ":")
		values = append(values, kv[1])
	}

	return values
}
