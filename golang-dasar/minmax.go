package main

import (
	"fmt"
	"sort"
)

//
//func miniMaxSum(arr []int32) {
//	var min int32
//	var max int32
//
//	for i := 0; i < len(arr); i++ {
//		var sum int32
//
//		for j := 0; j < len(arr); j++ {
//			if arr[i] == arr[j] {
//				continue
//			}
//			fmt.Println(arr[j])
//			sum += arr[j]
//		}
//
//		fmt.Println("sum", sum)
//		fmt.Println("min", min)
//		fmt.Println("max", max)
//		fmt.Println("=====")
//
//		if sum < min {
//			fmt.Println("min =")
//			min = sum
//		}
//		if sum > max {
//			max = sum
//		}
//	}
//
//	fmt.Println(min, max)
//}

func miniMaxSum(s []int) {
	fmt.Println("MiniMaxSum")
	sort.Ints(s)
	// fmt.Println(s)
	var maxSum int
	var minSum int
	var sLen = len(s)
	// fmt.Println(sLen)
	for _, v := range s[:sLen-1] {
		// fmt.Println(v)
		minSum += v
	}
	fmt.Println(minSum)

	for _, v := range s[1:] {
		// fmt.Println(v)
		maxSum += v
	}
	fmt.Println(maxSum)

}

func main() {

	var arr = []int{7, 69, 2, 221, 8974}
	miniMaxSum(arr)
}
