package main

import (
	"sort"
)

func main(nums []int) [][]int {
	var ret [][]int
	sort.Ints(nums)

	for first := 0; first < len(nums)-2; first++ {
		firstNum := nums[first]

		if first > 0 && nums[first-1] == firstNum {
			continue
		}
		second, last := first+1, len(nums)-1

		for second < last {
			secondNum, lastNums := nums[second], nums[last]
			sum := firstNum + secondNum + lastNums

			if second > first+1 && secondNum == nums[second-1] {
				second++
				continue
			}

			if sum == 0 {
				ret = append(ret, []int{firstNum, secondNum, lastNums})
				second++
				last--
			} else if sum < 0 {
				second++
			} else {
				last--
			}
		}
	}
	return ret
}
