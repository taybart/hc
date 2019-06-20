package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	p := os.Getenv("PORT")
	if p == "" {
		p = "80"
	}
	_, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/hc", p))
	if err != nil {
		os.Exit(1)
	}
}
