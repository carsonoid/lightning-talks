package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

// START MAIN OMIT
var (
	url = ""
)

func main() {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "You called "+r.URL.Path)
	}))
	defer s.Close()
	url = s.URL

	go startHealthProbes()

	doHello()
	time.Sleep(time.Second)
	doHello()
}

// END MAIN OMIT

func doHello() {
	r, _ := http.Get(url)
	io.Copy(os.Stdout, r.Body)
	fmt.Println()
}

// START HEALTH OMIT
func startHealthProbes() {
	for range time.Tick(time.Second / 2) {
		doHeathProbe()
	}
}

func doHeathProbe() {
	// START BUG1 OMIT
	// the health probe neeeds a timeout (default is 0 which means no timeout)
	http.DefaultClient.Timeout = time.Second * 30
	// END BUG1 OMIT

	// START BUG2 OMIT
	// the health probe needs to use a different url
	url = url + "/health"
	// END BUG2 OMIT

	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	// read until EOF
	io.Copy(io.Discard, r.Body)
}

// END HEALTH OMIT
