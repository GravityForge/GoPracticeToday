package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	contentBytes, err := ioutil.ReadFile("names.txt")
	if err == nil {
		var contentStr string = string(contentBytes)
		fmt.Println(contentStr)
	}
}
