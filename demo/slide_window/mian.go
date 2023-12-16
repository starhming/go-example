package main

import (
	"fmt"
	"time"
)

// SlideWindowSplit 按滑动窗口将数据切分
func SlideWindowSplit(data []int, windowSize, step int) [][]int {
	var result [][]int

	// 确保数据不为空且窗口大小大于0
	if len(data) == 0 || windowSize <= 0 {
		return result
	}

	// 遍历数据
	var i int
	for i = 0; i <= len(data)-windowSize; i += step {
		// 根据当前位置切割窗口
		window := data[i : i+windowSize]

		// 将窗口添加到结果中
		result = append(result, window)
	}

	// 处理剩余的数据，确保最后一个窗口也被加入
	if i < len(data)-1 {
		lastWindow := data[i:]
		result = append(result, lastWindow)
	}

	// 处理数据长度小于窗口大小的情况
	if len(data) < windowSize {
		result = append(result, data)
	}

	return result
}

func main() {
	// 示例数据
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//data := []int{1}

	// 定义滑动窗口大小为3，步幅为2
	windowSize := 3
	step := 2

	// 切分数据
	result := SlideWindowSplit(data, windowSize, step)

	// 打印结果
	for i, window := range result {
		fmt.Printf("Window %d: %v\n", i+1, window)
	}

	fmt.Println(int(5 * time.Minute / time.Millisecond))
}
