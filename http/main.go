package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	lw := logWriter{}

	// io.Copy(os.Stdout, resp.Body)

	io.Copy(lw, resp.Body)
}

// (logWriter) is the receiver with no variable since it won't be used
// (bs []byte) is the parameter that will be passed an argument of type byte slice
// (int, error) are the two required return types from this function
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
