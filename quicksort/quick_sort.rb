def quick_sort(list, left = nil, right = nil)
  left ||= 0
  right ||= list.length - 1
  return list if left > right
  i, j = left, right
  pivot = find_pivot(list[left..right])
  p "left #{left}, right #{right}, pivot: #{pivot}, target: #{list[left..right]}}"

  while true
    while(list[i] < pivot) do
      i = i + 1
    end
    while(pivot < list[j] ) do
      j = j - 1
    end
    break if i >= j

    list[i], list[j] = list[j], list[i]
    i = i + 1
    j = j - 1
  end

  list = quick_sort(list, left, i - 1)
  list = quick_sort(list, j + 1, right)
end


def find_pivot(list)
  # p "find:#{list}"
  list[rand(0..list.length-1)]
end
