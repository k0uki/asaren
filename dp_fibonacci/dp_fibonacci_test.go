package main

import (
  "testing"
)

func TestFib(t *testing.T){
  actual := fib(8)
  expected := 21

  if actual != expected {
    t.Errorf("got %v\nwant %v", actual, expected)
  }
}

func TestFibDP(t *testing.T){
  actual := fib_dp(8)
  expected := 21

  if actual != expected {
    t.Errorf("got %v\nwant %v", actual, expected)
  }
}
