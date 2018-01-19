require 'minitest/autorun'
require 'minitest/unit'

require_relative 'binary_search'

class BinarySearchTest < MiniTest::Unit::TestCase
  def test_binary_search_simple
    list = [1,3,5,9,10,15,25,30,45]

    assert_equal(7, binary_search(list, 30))
    assert_equal(false, binary_search(list, 0))
    assert_equal(false, binary_search(list, 9999))
  end

  def test_binary_search_bench
    list = []
    length = 10000
    length.times do
      list << rand(1..length)
    end

    list.sort!.uniq!
    length = list.length
    target = list[rand(0..(length-1))]

    assert_equal(list.find_index(target), binary_search(list, target))
  end
end
