package main

import "fmt"

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	result := missingNumber(nums)
	fmt.Println(result)
}

func missingNumber(nums []int) int {
	n := len(nums)
	return n*(n+1)/2 - sum(nums) // арифметическая прогрессия
}

func sum(nums []int) (result int) {
	for i := range nums {
		result += nums[i]
	}
	return
}
