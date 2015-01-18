package halt

import (
	"fmt"
)

func As(code int, msg ...interface{}) {
	panic(fmt.Sprintln("halt:", code, msg))
}
