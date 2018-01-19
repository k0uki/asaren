def bubble_sort(list)
  p list
  (list.length - 1).times do |i|
    (list.length - i - 1).times do |j|
      if list[j] > list[j + 1]
        list[j], list[j + 1] = list[j + 1], list[j]
      end
      p list
    end
  end
  list
end
