package main

import "fmt"

func maxArea(heights []int) int {
	start := 0
	end := len(heights) - 1
	ans := 0
	for start < end {
		width := end - start
		height := min(heights[start], heights[end])
		current := width * height
		ans = max(ans, current)
		if heights[start] < heights[end] {
			start++
		} else {
			end--
		}
	}
	return ans
}

func main() {
	arr := []int{1, 7, 2, 5, 4, 7, 3, 6}
	fmt.Println(maxArea(arr))
}
