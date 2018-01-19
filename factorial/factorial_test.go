package Factorial

import "testing"

func TestFactorial(t *testing.T){
  actual := Factorial(5)
  expected := 120

  if actual != expected {
    t.Errorf("got %v\nwant %v", actual, expected)
  }
}

func Benchmark_Factorial(b *testing.B){
  for i := 0; i < b.N; i++ {
    Factorial(i)
  }
}
