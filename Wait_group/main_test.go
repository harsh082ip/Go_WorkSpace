package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {

	sout := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	var wg sync.WaitGroup

	wg.Add(1)

	go printSomething("hey", &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = sout

	if !strings.Contains(output, "hey") {
		t.Errorf("Expected to find hey, but it was not there")
	}

}
