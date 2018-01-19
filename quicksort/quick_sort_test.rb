require 'minitest/autorun'
require 'minitest/unit'

require_relative 'quick_sort'

class QuickSortTest < MiniTest::Unit::TestCase
  def test_quick_sort
    list = []
    1000.times do
      list << rand(1..100000)
    end
    sorted = quick_sort(list)
    assert_equal(list.sort, sorted)
    p sorted
  end
end
