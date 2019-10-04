package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gookit/color"
)

func main() {
	green := color.FgGreen.Render
	red := color.FgRed.Render
	url, status := statusCode("https://google.com/")
	if status != 200 {
		fmt.Printf("URL: %s\nStatus: %s\n", url, red(status))
	} else {
		fmt.Printf("URL: %s\nStatus: %s\n", url, green(status))

	}
}

func statusCode(url string) (string, int) {
	red := color.FgRed.Render
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Your shit's all fucked up:\n --> %s", red(err))
	}
	defer resp.Body.Close()
	return url, resp.StatusCode
}
