require 'minitest/autorun'
require 'minitest/unit'
require 'benchmark'
require_relative 'dp_fibonacci'

class DpFibTest < MiniTest::Unit::TestCase
  def test_fib
    assert_equal(0, fib(0))
    assert_equal(1, fib(1))
    assert_equal(3, fib(4))
    assert_equal(5, fib(5))
    assert_equal(21, fib_dp(8))
  end

  def test_fib_bench
    Benchmark.bm(7) do |b|
      b.report('fib(7)') do
        fib(35)
      end
      b.report('fib_dp(7)') do
        fib_dp(35)
      end
    end
  end
end
