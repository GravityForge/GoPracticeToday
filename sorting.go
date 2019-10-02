package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{1, 5, 3, 6, 2, 10, 8}
	tobeSorted := sort.IntSlice(numbers)
	sort.Sort(sort.Reverse(tobeSorted))
	fmt.Println(tobeSorted)
	//opposite
	gonnabeSorted := sort.IntSlice(numbers)
	sort.Sort(gonnabeSorted)
	fmt.Println(gonnabeSorted)
}