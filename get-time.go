package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()
	fmt.Println(current.String())
	fmt.Println("MM-DD-YYYY:", current.Format("01-02-2006"))
}
