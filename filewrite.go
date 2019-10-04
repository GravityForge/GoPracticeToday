package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	hello := "Hello,World"
	err := ioutil.WriteFile("hello_world", []byte(hello), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
