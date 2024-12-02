package main

import (
	"testing"
)

func TestCalculateSumAndAverage(t *testing.T) {
	// 接受一个 *testing.T 类型的参数 t，这个参数是 Go 语言测试框架提供的，用于报告测试失败和其他测试相关的信息。
	// 定义测试用例
	testCases := []struct {
		nums        []int
		expectedSum int
		expectedAvg float64
	}{
		{[]int{1, 2, 3, 4, 5}, 15, 3.0},
		{[]int{10, 20, 30, 40, 50}, 150, 30.0},
		{[]int{-1, -2, -3, -4, -5}, -15, -3.0},
	}

	// 遍历测试用例
	for _, tc := range testCases {
		// 调用被测试的函数
		sum, avg := calculateSumAndAverage(tc.nums)
		// 断言结果
		if sum != tc.expectedSum {
			t.Errorf("Expected sum %d, but got %d", tc.expectedSum, sum)
		}
		if avg != tc.expectedAvg {
			t.Errorf("Expected average %f, but got %f", tc.expectedAvg, avg)
		}
	}
}

// 假设的被测试函数
func calculateSumAndAverage(nums []int) (int, float64) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	avg := float64(sum) / float64(len(nums))
	return sum, avg
}
