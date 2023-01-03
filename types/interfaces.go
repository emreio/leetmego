package types

import "fmt"

type IHttpServer interface {
	StartServer()
}

type MyString float32

func (m MyString) Describe() {
	fmt.Printf("i am  %v", m)
}
