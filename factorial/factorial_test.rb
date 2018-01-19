require 'minitest/autorun'
require 'minitest/unit'

require_relative 'factorial'

class FactorialTest < MiniTest::Unit::TestCase
  def test_factorial
    assert_equal(120, factorial(5))
    assert_equal(1, factorial(0))
  end

  def test_factorial_bench
    10000.times do |i|
      factorial(i)
    end
  end
end
