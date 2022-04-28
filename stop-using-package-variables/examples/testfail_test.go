package main

import (
	"net/http"
	"testing"
	"time"
)

// START EXAMPLE1 OMIT
func Test_doGet(t *testing.T) {
	// set a short timeout to test failures
	client.Timeout = time.Millisecond

	// testing not shown
}

// END EXAMPLE1 OMIT

// START EXAMPLE2 OMIT
func Test_doGet(t *testing.T) {
	// make a copy of the client and restore it after the test
	c := client
	defer func() {
		client = c
	}()

	// set a short timeout to test failures
	client.Timeout = time.Millisecond

	// testing not shown
}

// END EXAMPLE2 OMIT

// START EXAMPLE3 OMIT
func Test_doGet(t *testing.T) {
	// defer a reset of the client after the test
	c := *client
	defer func() {
		*client = c
	}()

	// set a short timeout to test failures
	client.Timeout = time.Millisecond

	// testing not shown
}

// END EXAMPLE3 OMIT

// START EXAMPLE4 OMIT
func Test_doGet(t *testing.T) {
	// defer a reset of the client after the test
	defer func() {
		client = &http.Client{}
	}()

	// set a short timeout to test failures
	client.Timeout = time.Millisecond

	// testing not shown
}

// END EXAMPLE4 OMIT
