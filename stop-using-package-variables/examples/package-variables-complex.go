package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
)

// START OMIT
var (
	url = ""
)

func main() {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "You did a "+r.Method+" on "+r.URL.Path)
	}))
	defer s.Close()
	url = s.URL

	getRoot()
}

func getRoot() {
	r, _ := http.Get(url)
	io.Copy(os.Stdout, r.Body)
	fmt.Println()
}

// END OMIT
