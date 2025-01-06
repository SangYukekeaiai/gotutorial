package main

import (
	"fmt"
)

var nums = []int{11, 2, 7, 15}
var target int = 9

func two_sum(nums []int, target int) []int{
	num_map := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		index, exists := num_map[complement]
		if exists {
			return []int{index, i}
		}
		num_map[nums[i]] = i
	}
	return nil
}

func main()  {
	fmt.Println(two_sum(nums, target))
}