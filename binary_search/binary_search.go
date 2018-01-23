package main

import (
	"fmt"
	"sort"
)

func BinarySearch(list []int, target int, imin int, imax int) int {
	if imax < imin {
		return imax + 1
	} else {
		mid := imin + (imax-imin)/2
		if list[mid] > target {
			return BinarySearch(list, target, imin, mid-1)
		} else if list[mid] < target {
			return BinarySearch(list, target, mid+1, imax)
		} else {
			return mid
		}
	}
}

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15}
	fmt.Println(sort.SearchInts(list, 0))
	fmt.Println(BinarySearch(list, 5, 0, len(list)-1))
}
