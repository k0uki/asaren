package main

import "fmt"

func BubbleSort(list []int) []int{
  for i := range list {
    for j := 0; j < len(list) -i -1; j++ {
      if list[j] > list[j+1] {
        list[j], list[j + 1] = list[j + 1], list[j]
      }
    }
  }
  return list
}

func main() {
  fmt.Println(BubbleSort([]int{3,2,1}))
}
