package main

import (
	"fmt"
	"io"
	"net/http"
)

// START OMIT
var client = &http.Client{}

func main() {
	err := doGet()
	if err != nil {
		panic(err)
	}
}

func doGet() error {
	r, err := client.Get("https://www.google.com")
	if err != nil {
		return err
	}
	defer r.Body.Close()
	io.Copy(io.Discard, r.Body)
	fmt.Println("GET is done")
	return nil
}

// END OMIT
