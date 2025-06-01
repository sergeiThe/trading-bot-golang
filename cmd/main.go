package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil { panic(err) }
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response:", string(body))
}
