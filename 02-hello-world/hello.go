package hello

import (
	"strings"
)

// func <func name (argName argType) returnType
func Say(names []string) string {
	if len(names) == 0 {
		names = []string{"world"}
	}
		return "Hello, " + strings.Join(names, ", ") + "!"
}