package main

import "fmt"

func main() {
	array := []int{2, 7, 11, 15}
	target := 9

	result := twoSum1(array, target)
	fmt.Println(result)
    result = twoSum2(array, target)
	fmt.Println(result)
	result = twoSum3(array, target)
	fmt.Println(result)
}

func twoSum1(nums []int, target int) []int {
	for i, first := range nums {
		for j, second := range nums {
			if (first+second) == target && i != j {
				return []int{i, j}
			}
		}
	}

	return []int{0, 0}
}

/*
twoSum1 цикл внутри цикла,  O(n2). Очень долго
*/

func twoSum2(nums []int, target int) []int {
	hashMap := make(map[int]int) // инициализация хэш таблицы

	for i := range nums {
		hashMap[nums[i]] = i // сразу заполняю таблицу (map)
	}

	for i := range nums {
		remainder := target - nums[i]     // остаток
		if key, ok := hashMap[remainder]; ok && i != key { // возвращает сам элемент и true/false
			return []int{i, key}
		}
	}
	return []int{0, 0}
}

/*
twoSum2 реализовано с помощью хэш таблицы, не плохо
*/

func twoSum3(nums []int, target int) []int {
	hashMap := make(map[int]int) // инициализация хэш таблицы

	for i := range nums {
		remainder := target - nums[i] // остаток
		if key, ok := hashMap[remainder]; ok && i != key { // возвращает сам элемент и true/false
			return []int{i, key}
		}

		hashMap[nums[i]] = i // заполняю в итерации
	}
	return []int{0, 0}
}

/*
twoSum3 реализовано с помощью хэш таблицы, оптимальное решение O(n)
*/