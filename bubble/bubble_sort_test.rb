require 'minitest/autorun'
require 'minitest/unit'

require_relative 'bubble_sort'

class BubbleSortTest < MiniTest::Unit::TestCase
  def test_bubble_sort
    list = []
    300.times do
      list << rand(1..100)
    end
    assert_equal(list.sort, bubble_sort(list))
  end
end
