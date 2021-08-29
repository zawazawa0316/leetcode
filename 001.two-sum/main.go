package main

import "fmt"

func twoSum(nums []int, target int) []int {
	indices := []int{0, 1}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				indices[0] = i
				indices[1] = j
				return indices
			}
		}
	}
	return indices
}

func main() {
	input := []int{2, 7, 11, 15}
	target := 18
	output := twoSum(input, target)
	fmt.Println(output)
}
