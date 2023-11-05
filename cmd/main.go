package main

import (
	"fmt"
	"github.com/nikaNGC224/golang-sort-lib/pkg/customSort"
)

func main() {
	nums := []int{1, 3, 2, 5, 10}
	fmt.Println(customSort.BubbleSort(nums))
}
