package array

import "sort"

//ThreeSum 三数之和
func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	results := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		two := 0 - nums[i]
		for j := 1; j < len(nums); j++ {
			one := two - nums[j]
			for k := 2; k < len(nums); k++ {
				if nums[k] == one {
					results = append(results, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}

	return results
}

//TreeSum1 不重复的三位数之和的数组
func TreeSum1(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	results := make([][]int, 0)
	sort.Ints(nums)
	for first := 0; first < len(nums); first++ {
		//保证了没有重复的数组的出现
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := len(nums) - 1
		target := -1 * nums[first]
		for second := first + 1; second < len(nums); second++ {
			//make sure no repeat
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				results = append(results, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return results
}
