package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	bs := make([]byte, 99999)
	_, _ = resp.Body.Read(bs)
	//if err != nil {
	//	fmt.Println("Error2:", err)
	//	return
	//}
	fmt.Println(string(bs))
}
