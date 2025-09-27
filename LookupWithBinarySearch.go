package main

func binarySearch(nums []int32, target int32) int32 {
	// Write your code here
	if len(nums) < 1 {
		return -1
	}
	low := 0
	high := len(nums) - 1
	for {
		if target == nums[(low+high)/2] {
			return int32(low+high) / 2
		}
		if low == high {
			break
		}
		if target > nums[(low+high)/2] {
			low = (low+high)/2 + 1
		} else {
			high = (low + high) / 2
		}

	}
	return -1
}
