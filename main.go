package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("request failed")

func main() {
	urls := []string{
		"https://www.www.google.com",
		"https://www.amazon.com",
		"https://www.flipkart.com",
		"https://www.reddit.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://www.airbnb.com",
		"https://www.github.com",
		"https://www.linkedin.com",
	}

	for _, url := range urls {
		hitURL(url)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking -> ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
