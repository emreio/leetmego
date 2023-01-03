package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net/http"
	"time"

	"emrekantar.com.tr/leetmego/v1/types"
)

var httpServer types.IHttpServer

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in if ", r)
		}
	}()

	bytess := new(bytes.Buffer)

	var vl uint64 = 24

	go func(bytess *bytes.Buffer, vl uint64) {
		binary.Write(bytess, binary.LittleEndian, vl)
		fmt.Println("in go routine")
	}(bytess, vl)

	fmt.Printf("buffer length: %+v\n", bytess.Len())

	// if b, err := bytess.ReadByte(); b == 255 && err == nil {
	// 	fmt.Println("correct")
	// } else {
	// 	fmt.Println(b)
	// 	fmt.Println("incorrect")
	// }
	for i := 0; i < 100; i++ {
		fmt.Printf("buffer length: %+v\n", bytess.Len())
		time.Sleep(200)
	}

	// return

	// httpServer = sandbox.NewHttpServer(4444, getHandler, postHandler)

	// httpServer.(*sandbox.Httpserver).Port = 2323

	// fmt.Printf("%p  %T %#v %+v\n", httpServer, httpServer, httpServer, httpServer)

	// httpServer.StartServer()
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello go server, i am overrided"))
}
func postHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello go server, i am overrided"))
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
