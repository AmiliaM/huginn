package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://api.github.com/")
	fmt.Println(resp)
}
