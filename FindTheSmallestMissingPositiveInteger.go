package main

func findSmallestMissingPositive(orderNumbers []int32) int32 {
	// Write your code here
	if len(orderNumbers) < 1 {
		return 1
	}
	for k := range orderNumbers {
		for orderNumbers[k] != int32(k+1) {
			if orderNumbers[k] <= 0 || orderNumbers[k] > int32(len(orderNumbers)) {
				break
			}
			if orderNumbers[k] == orderNumbers[orderNumbers[k]-1] {
				break
			}
			temp := orderNumbers[k]
			orderNumbers[k] = orderNumbers[orderNumbers[k]-1]
			orderNumbers[temp-1] = temp
		}
	}

	for k := range orderNumbers {
		if orderNumbers[k] != int32(k+1) {
			return int32(k + 1)
		}
	}
	return int32(len(orderNumbers) + 1)
}
