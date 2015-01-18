package assert

import (
	"fmt"
)

func For(cond bool, code int, msg ...interface{}) {
	e := fmt.Sprint(code)
	if !cond {
		switch {
		case (code >= 20) && (code < 40):
			e = fmt.Sprintln(code, "precondition violated", fmt.Sprint(msg...))
		case (code >= 40) && (code < 60):
			e = fmt.Sprintln(code, "subcondition violated", fmt.Sprint(msg...))
		case (code >= 60) && (code < 80):
			e = fmt.Sprintln(code, "postcondition violated", fmt.Sprint(msg...))
		default:
		}
		panic(e)
	}
}
