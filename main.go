package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	wr, err := io.Copy(os.Stdout, resp.Body)
	if err != nil {
		return
	}
	fmt.Println(wr)
}
