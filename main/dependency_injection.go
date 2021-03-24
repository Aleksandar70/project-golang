package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

//fmt.Fprintf is like fmt.Printf
//but instead takes a Writer to send the string to, whereas fmt.Printf defaults to stdout.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}
