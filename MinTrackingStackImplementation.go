package main

import (
	"strconv"
	"strings"
)

func processCouponStackOperations(operations []string) []int32 {
	// Write your code here
	stack := []int32{}
	result := []int32{}
	var min int32 = 0
	findMin := func(arr []int32) int32 {
		var min int32 = 0
		if len(arr) > 0 {
			min = arr[0]
		}
		for _, val := range arr {
			if val < min {
				min = val
			}
		}
		return min
	}
	for _, operation := range operations {
		args := strings.SplitN(operation, " ", 2)
		switch args[0] {
		case "push":
			num, _ := strconv.Atoi(args[1])
			if len(stack) < 1 {
				min = int32(num)
			}
			if int32(num) < min {
				min = int32(num)
			}
			stack = append(stack, int32(num))
		case "pop":
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if top == min {
					min = findMin(stack)
				}
			}
		case "top":
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				result = append(result, top)
			}
		case "getMin":
			result = append(result, min)
		}

	}
	return result
}
