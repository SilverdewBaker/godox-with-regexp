package godoxwithregexp

import (
	"fmt"
	"go/token"
)

// Issue is a linting issue.
type Issue struct {
	// Message - the message to display to the user.
	Message string
	// Pos - the position of the issue.
	Pos token.Position
}

func (i *Issue) String() string {
	return fmt.Sprintf("%s: %s", i.Pos, i.Message)
}

func filter[T any](xs []T, f func(T) bool) []T {
	var ys []T
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}
