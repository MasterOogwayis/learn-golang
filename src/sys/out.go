package sys

import "fmt"

func Printf(str string, a ...interface{}) (n int, err error) {
	return fmt.Printf(str, a)
}
