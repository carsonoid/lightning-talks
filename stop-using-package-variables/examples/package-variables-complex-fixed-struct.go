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

	hp := &HealthProber{
		client: &http.Client{Timeout: time.Second * 30},
		url:    s.URL + "/health",
	}
	go hp.Start()

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

type HealthProber struct {
	url    string
	client *http.Client
}

func (hp *HealthProber) Start() {
	for range time.Tick(time.Second / 2) {
		hp.Probe()
	}
}

func (hp *HealthProber) Probe() {
	r, err := hp.client.Get(hp.url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	// read until EOF
	io.Copy(io.Discard, r.Body)
}

// END HEALTH OMIT
