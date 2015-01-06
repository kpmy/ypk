package halt

import (
	"fmt"
)

func As(code int) {
	panic(fmt.Sprintln("halt:", code))
}
