package main

import "fmt"

func twoSum(nums []int, target int) []int {
	tempMap := make(map[int]int) // val:index

	for i, item := range nums {
		if val, ok := tempMap[target-item]; ok {
			return []int{val, i}
		}
		tempMap[item] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target)) // Output: [0 1]
}
