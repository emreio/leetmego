package sandbox

import (
	"fmt"
)

func Do() {
	var someType interface{}

	switch someType.(type) {
	case typeInterface:
		fmt.Println("some interface")
	case float32:
		fmt.Println("this is human")
	case int:
		fmt.Println("this is an integer")
	default:
		fmt.Println("unknown type ")
	}
}

type typeInterface interface {
	dowork()
}
