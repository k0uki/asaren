package QuickSort

import (
  "testing"
  "reflect"
  "sort"
)

func TestQuickSort(t *testing.T){
  list := []int{3,4,5,99,1,2,3000,5}
  actual := QuickSort(list, 0, len(list) - 1)
  
  sort.Sort(sort.IntSlice(list))
  expected := list

  if !reflect.DeepEqual(actual,expected) {
    t.Errorf("got %v\nwant %v", actual, expected)
  }
}

// func Benchmark_Factorial(b *testing.B){
//   for i := 0; i < b.N; i++ {
//     Factorial(i)
//   }
// }
