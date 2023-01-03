package util

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
)

func PromptSomething(prompt string) (int, bool) {
	//DO something
	return 5, false
}

func WriteFile(path string, content string) error {
	var error customError
	error.err = "hata var hataaaa"
	return error
}

type customError struct {
	err string
}

func (e customError) Error() string {
	return e.err
}

func SomeMethod(a interface{}) {
	fmt.Println(a)
}

func ReadRandomBytes(length int) ([]byte, error) {
	buffer := make([]byte, length)
	_, err := io.ReadFull(rand.Reader, buffer)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return buffer, nil
}
