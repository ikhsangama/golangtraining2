package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWritter struct {
}

func (lw logWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	lw := logWritter{}

	wr, err := io.Copy(lw, resp.Body)
	if err != nil {
		return
	}
	fmt.Println(wr)
}
