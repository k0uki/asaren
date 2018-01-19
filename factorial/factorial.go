package Factorial

import "fmt"

func Factorial(n int) int{
  if n > 0 {
    return n * Factorial(n-1)
  } else {
    return 1
  }
}

func main() {
  fmt.Println(Factorial(5))
}
