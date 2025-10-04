package main

func countAffordablePairs(prices []int32, budget int32) int32 {
	// Write your code here
	var count int32
	for k, vk := range prices {
		for _, vj := range prices[k+1:] {
			sum := vk + vj
			if sum <= budget {
				count++
			}
		}
	}
	return count
}
