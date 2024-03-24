package main

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	if len(nums) < k {
		return -1
	}

	pivot := nums[0]
	var left []int
	var right []int

	for _, num := range nums[1:] {
		if num > pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	if len(left) == k-1 {
		return pivot
	} else if len(left) >= k {
		return findKthLargest(left, k)
	} else {
		return findKthLargest(right, k-len(left)-1)
	}
}

func main() {
	nums1 := []int{3, 2, 1, 5, 6, 4}
	k1 := 2
	fmt.Printf("Input: nums = %v, k = %d\nOutput: %d\n", nums1, k1, findKthLargest(nums1, k1))

	nums2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k2 := 4
	fmt.Printf("Input: nums = %v, k = %d\nOutput: %d\n", nums2, k2, findKthLargest(nums2, k2))
}
