package QuickSort

import "fmt"

func QuickSort(list []int, left, right int) []int{
  fmt.Println("left: right", left, right)
  if left > right { return list }

  target := make([]int, len(list))
  copy(target, list)
  i , j := left , right
  pivot := target[j]
  for {
    for target[i] < pivot { i++ }
    for pivot < target[j] { j--}
    if i >= j { break }
    if i < j {
      target[i], target[j] = target[j], target[i]
    }
    i++;j--;
  }
  target = QuickSort(target, left, i -1)
  target = QuickSort(target, j + 1, right)
  return target
}

func main() {
  list := []int{3,4,1,7,9,10,12}
  fmt.Println(QuickSort(list, 0, len(list) -1))
}
