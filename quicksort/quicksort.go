package QuickSort

import "fmt"

func QuickSort(list []int, left, right int) {
	fmt.Println("left: right", left, right)
	if left > right {
		return
	}

	i, j := left, right
	pivot := list[j]
	for {
		for list[i] < pivot {
			i++
		}
		for pivot < list[j] {
			j--
		}
		if i >= j {
			break
		}
		if i < j {
			list[i], list[j] = list[j], list[i]
		}
		i++
		j--
	}
	QuickSort(list, left, i-1)
	QuickSort(list, j+1, right)
}

func main() {
	// list := []int{3,4,1,7,9,10,12}
	// return list
	// fmt.Println(QuickSort(list, 0, len(list) -1))
}
