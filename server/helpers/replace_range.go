package helpers

import "chat-app/structs"

func ReplaceRange(input string, rng structs.Range, rw string) string {
	s, e := rng.Start(), rng.End()

    return input[:s] + rw + input[e:]
}
