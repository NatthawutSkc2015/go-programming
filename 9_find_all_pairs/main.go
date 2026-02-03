package main

import "fmt"

// TwoSumAllPairs finds all pairs of indices where nums[i] + nums[j] = target
// Returns all unique pairs (no duplicate pairs, no same index used twice)
// Time Complexity: O(n), Space Complexity: O(n)
func TwoSumAllPairs(nums []int, target int) [][]int {
	// Result to store all pairs
	result := [][]int{}

	// Map to store: value -> list of indices
	numMap := make(map[int][]int)

	// Set to track used pairs (to avoid duplicates)
	usedPairs := make(map[string]bool)

	for i, num := range nums {
		// Calculate the complement needed
		complement := target - num

		// Check if complement exists in map
		if indices, found := numMap[complement]; found {
			// Found complement, create pairs with all its indices
			for _, j := range indices {
				// Create a unique key for this pair (sorted indices)
				var pairKey string
				if j < i {
					pairKey = fmt.Sprintf("%d-%d", j, i)
				} else {
					pairKey = fmt.Sprintf("%d-%d", i, j)
				}

				// Add pair if not already used
				if !usedPairs[pairKey] {
					result = append(result, []int{j, i})
					usedPairs[pairKey] = true
				}
			}
		}

		// Store current number and its index
		numMap[num] = append(numMap[num], i)
	}

	return result
}

func main() {
	fmt.Println("=== Two Sum All Pairs Problem ===")

	// Test case 1: Single pair
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	result1 := TwoSumAllPairs(nums1, target1)
	fmt.Printf("Input: nums = %v, target = %d\n", nums1, target1)
	fmt.Printf("Output: %v\n", result1)
	if len(result1) > 0 {
		for _, pair := range result1 {
			fmt.Printf("  nums[%d] + nums[%d] = %d + %d = %d\n",
				pair[0], pair[1], nums1[pair[0]], nums1[pair[1]], target1)
		}
	}
	fmt.Println()

	// Test case 2: Multiple pairs
	nums2 := []int{1, 5, 3, 2, 4, 6}
	target2 := 7
	result2 := TwoSumAllPairs(nums2, target2)
	fmt.Printf("Input: nums = %v, target = %d\n", nums2, target2)
	fmt.Printf("Output: %v\n", result2)
	if len(result2) > 0 {
		for _, pair := range result2 {
			fmt.Printf("  nums[%d] + nums[%d] = %d + %d = %d\n",
				pair[0], pair[1], nums2[pair[0]], nums2[pair[1]], target2)
		}
	}
	fmt.Println()

	// Test case 3: Duplicate values
	nums3 := []int{3, 3, 3}
	target3 := 6
	result3 := TwoSumAllPairs(nums3, target3)
	fmt.Printf("Input: nums = %v, target = %d\n", nums3, target3)
	fmt.Printf("Output: %v\n", result3)
	if len(result3) > 0 {
		for _, pair := range result3 {
			fmt.Printf("  nums[%d] + nums[%d] = %d + %d = %d\n",
				pair[0], pair[1], nums3[pair[0]], nums3[pair[1]], target3)
		}
	}
	fmt.Println()

	// Test case 4: No pairs found
	nums4 := []int{1, 2, 3}
	target4 := 10
	result4 := TwoSumAllPairs(nums4, target4)
	fmt.Printf("Input: nums = %v, target = %d\n", nums4, target4)
	fmt.Printf("Output: %v\n", result4)
	if len(result4) == 0 {
		fmt.Println("  No pairs found")
	}
	fmt.Println()

	// Test case 5: Many duplicate values forming multiple pairs
	nums5 := []int{2, 4, 2, 4, 2}
	target5 := 6
	result5 := TwoSumAllPairs(nums5, target5)
	fmt.Printf("Input: nums = %v, target = %d\n", nums5, target5)
	fmt.Printf("Output: %v\n", result5)
	if len(result5) > 0 {
		for _, pair := range result5 {
			fmt.Printf("  nums[%d] + nums[%d] = %d + %d = %d\n",
				pair[0], pair[1], nums5[pair[0]], nums5[pair[1]], target5)
		}
	}
	fmt.Println()

	// Test case 6: Array with negative numbers
	nums6 := []int{-1, 0, 1, 2, -1, -4, 3}
	target6 := 0
	result6 := TwoSumAllPairs(nums6, target6)
	fmt.Printf("Input: nums = %v, target = %d\n", nums6, target6)
	fmt.Printf("Output: %v\n", result6)
	if len(result6) > 0 {
		for _, pair := range result6 {
			fmt.Printf("  nums[%d] + nums[%d] = %d + %d = %d\n",
				pair[0], pair[1], nums6[pair[0]], nums6[pair[1]], target6)
		}
	}
}
