package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// GetInput downloads the Input for a specific Day of the Advent of Code
func GetInput(adventOfCodeYear int, adventOfCodeDay int) string {
	sessionCookie := os.Getenv("ADVENT_OF_CODE_COOKIE")

	if sessionCookie == "" {
		log.Fatal("ADVENT_OF_CODE_COOKIE env var is unset")
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", adventOfCodeYear, adventOfCodeDay), nil)
	if err != nil {
		log.Fatalf("error creating HTTP request: %s", err)
	}

	req.Header.Add("Cookie", fmt.Sprintf("session=%s", sessionCookie))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("error executing HTTP request: %s", err)
	}
	defer resp.Body.Close()

	var bodyBytes []byte
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("error reading HTTP response body: %s", err)
		}
	} else {
		log.Fatalf("HTTP response status code: %d", resp.StatusCode)
	}

	return string(bodyBytes)
}
