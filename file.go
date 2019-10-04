package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("log.txt"); err == nil {
		fmt.Println("File found")
	} else {
		fmt.Println("File not found")
	}
}
