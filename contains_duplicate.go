package main

import "fmt"

func main() {
	array := []int{5, 6, 5, 8}
	result := containsDuplicate1(array) // 1 вариант
	fmt.Println(result)

	result = containsDuplicate2(array) // 2 вариант
	fmt.Println(result)
}

func containsDuplicate1(nums []int) bool {
	result := 0
	for i := range nums {
		result = unique(nums, nums[i])
		if result >= 2 {
			return true // duplicate
		}
	}
	return false
}

func unique(nums []int, item int) int {
	result := 0
	for i := range nums {
		if nums[i] == item {
			result = result + 1
		}
	}
	return result
}

/*
1 вариант не самое лучшее решение для этой задачи, потому что цикл внутри цикла + долго
*/

func containsDuplicate2(nums []int) bool {
	twice := make(map[int]bool, 0)

	for i := range nums {
		if twice[nums[i]] == true {
			return true
		} else {
			twice[nums[i]] = true
		}
	}
	return false
}

/*
2 вариант оптимальное решение для этой задачи, реализовано с помощью хэш таблицы (map + make)
*/
