package main

func findFirstOccurrence(nums []int32, target int32) int32 {
	// Write your code here
	if len(nums) < 1 {
		return -1
	}
	var pos int32 = -1
	var low int32 = 0
	var high int32 = int32(len(nums) - 1)
	for {
		var avg = (low + high) / 2
		if target == nums[avg] {
			if avg < pos || pos == -1 {
				pos = avg
			}
		}
		if target > nums[avg] {
			low = avg + 1
		} else {
			high = avg
		}
		if low >= high {
			break
		}
	}

	return pos
}
