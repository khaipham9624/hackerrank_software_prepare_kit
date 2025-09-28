package main

import "sort"

func maximizeNonOverlappingMeetings(meetings [][]int32) int32 {
	// Write your code here
	if len(meetings) < 1 {
		return 0
	}
	sort.Slice(meetings, func(i int, j int) bool {
		return meetings[i][1] < meetings[j][1]
	})
	var count int32
	var high int32
	for _, meeting := range meetings {
		if meeting[0] < high {
			continue
		}
		high = meeting[1]
		count++
	}
	return count
}
