def fib(n)
  case n
  when 0
    return 0
  when 1
    return 1
  else
    # F(n + 2) = F(n) + F(n + 1)
    # F = F(n - 2) + F(n - 1)
    return fib(n - 2) + fib(n - 1)
  end
end

def fib_dp(n)
  memo = [0, 1]
  i = 2
  while(n >= i) do
    memo[i] = memo[i - 1] + memo[i - 2]
    i = i + 1
  end
  memo[n]
end
