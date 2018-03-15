def e_furui(n)
  prime = []
  target = Math.sqrt(n)
  list = (2..n).to_a
  while list[0] < target do
    prime << list.slice!(0)
    list.reject! {|n| n % prime.last == 0}
  end
  prime.concat(list)
end
