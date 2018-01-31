package main

import (
  "fmt"
)

func fib(n int) int{
  switch n {
  case 0: return 0
  case 1: return 1
  default: return fib(n-2) + fib(n-1)
  }
}

func fib_dp(n int) int{
  memo := []int{0, 1}
  i := 2
  for n >= i {
    memo = append(memo, memo[i-1] + memo[i -2])
    i++
  }

  return memo[n]
}

func main() {
  fmt.Println("hello from dp")
}
