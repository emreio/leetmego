package sandbox

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Httpserver struct {
	Port        int
	GetHandler  func(w http.ResponseWriter, r *http.Request)
	PostHandler func(w http.ResponseWriter, r *http.Request)
}

func (httpserver *Httpserver) StartServer() {

	if httpserver.GetHandler == nil {
		fmt.Println("get handler is nil, assigning default get handler")
		http.HandleFunc("/get", getHello)
	} else {
		http.HandleFunc("/get", httpserver.GetHandler)
	}

	http.HandleFunc("/post", httpserver.PostHandler)
	fmt.Printf("server started at port: %v\n", httpserver.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", httpserver.Port), nil))
}

func NewHttpServer(port int, getHandler func(w http.ResponseWriter, r *http.Request), postHandler func(w http.ResponseWriter, r *http.Request)) *Httpserver {
	return &Httpserver{Port: port, GetHandler: getHandler, PostHandler: postHandler}
}

func getHello(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(401)
		w.Write([]byte("GET NOT ALLOWED"))
		return
	}

	bodyBuffer := make([]byte, 1)

	for {
		_, err := io.ReadFull(r.Body, bodyBuffer)
		//fmt.Printf("%d  %e", count, err)
		if err == io.EOF {
			break
		}
		//fmt.Println(string(bodyBuffer))
	}

	for {
		io.CopyN(w, rand.Reader, 256)
		if _, ok := w.(http.Flusher); ok {
			//f.Flush()
			fmt.Println("response buffer flushed..")
		}
	}

}
