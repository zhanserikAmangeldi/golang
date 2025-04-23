package hello

import "strings"

func Say(names []string) string {
	if len(names) == 0 {
		names = append(names, "World")
	}

	return "Hello, " + strings.Join(names, ", ") + "!"
}
