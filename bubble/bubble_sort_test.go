package main

import (
  "testing"
  "reflect"
)

func TestBubble(t *testing.T){
  target := []int{10,5,8,45,20,60,15,1}
  actual := BubbleSort(target)
  expected := []int{1,5,8,10,15,20,45,60}

  if !reflect.DeepEqual(actual,expected) {
    t.Errorf("got %v\nwant %v", actual, expected)
  }
}

// func Benchmark_Factorial(b *testing.B){
//   for i := 0; i < b.N; i++ {
//     Factorial(i)
//   }
// }
