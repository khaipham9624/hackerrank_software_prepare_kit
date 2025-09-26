package main

func countResponseTimeRegressions(responseTimes []int32) int32 {
	// Write your code here
	var count int32 = 0
	var avg int64 = 0
	var sum int64 = 0
	if len(responseTimes) <= 1 {
		return count
	}
	for k, v := range responseTimes {
		if k < 1 {
			sum += int64(v)
			continue
		}
		avg = sum / int64(k)
		sum += int64(v)
		if int64(v) > avg {
			count++
		}
	}
	return count
}
