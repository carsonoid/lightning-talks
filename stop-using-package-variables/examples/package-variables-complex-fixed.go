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
func main() {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "You called "+r.URL.Path)
	}))
	defer s.Close()

	go startHealthProbes(s.URL)

	c := &http.Client{}
	doHello(c, s.URL)
	time.Sleep(time.Second)
	doHello(c, s.URL)
}

// END MAIN OMIT

// START GET OMIT
func doHello(c *http.Client, url string) {
	r, _ := c.Get(url)
	io.Copy(os.Stdout, r.Body)
	fmt.Println()
}

// END GET OMIT

// START HEALTH OMIT
func startHealthProbes(url string) {
	// the health probe makes its own client
	// with a longer timeout for reasons
	client := &http.Client{Timeout: time.Second * 30}

	// the health probe needs to use a different url
	url = url + "/health"

	for range time.Tick(time.Second / 2) {
		doHealthProbe(client, url)
	}
}

func doHealthProbe(c *http.Client, url string) {
	r, err := c.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	// read until EOF
	io.Copy(io.Discard, r.Body)
}

// END HEALTH OMIT
