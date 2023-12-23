package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	var qwerty int
	r, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Printf("%+v", r)
	io.Copy(os.Stdout, r.Body)
}
