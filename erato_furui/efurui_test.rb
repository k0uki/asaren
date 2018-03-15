require 'minitest/autorun'
require 'minitest/unit'
require_relative 'efurui'

class EfuruiTest < MiniTest::Unit::TestCase
  def test_furui
    expected = [2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97,101,103,107,109,113]
    actual = e_furui(120)
    assert_equal(expected, actual)
  end
end