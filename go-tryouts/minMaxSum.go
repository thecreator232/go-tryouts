package main

import (
	"fmt"
)

// problem statement: https://www.hackerrank.com/challenges/mini-max-sum/problem

func main() {

	arr := []int32{12, 45, 23, 56, 67}

	minMaxSum(arr)
}

//bruteforce
func minMaxSum(arr []int32) {
	max := arr[0]
	min := arr[0]
	sum := int64(arr[0])

	for i := 1; i < len(arr); i++ {
		sum += int64(arr[i])
		temp := arr[i]
		if temp > max {
			max = temp
		}
		if temp < min {
			min = temp

		}
		if min == 0 {
			min = temp
		}
		fmt.Println("\n", min, max, temp, i, a, arr)
	}

}
