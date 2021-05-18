/*
  Problem: https://leetcode.com/problems/number-of-good-pairs/

  Desc:
    Given an array of integers nums.
    A pair (i,j) is called good if nums[i] == nums[j] and i < j.
    Return the number of good pairs.

  Ex1:
    Input: nums = [1,2,3,1,1,3]
    Output: 4
    Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.

  Ex2:
    Input: nums = [1,1,1,1]
    Output: 6
    Explanation: Each pair in the array are good.

  Ex3:
    Input: nums = [1,2,3]
    Output: 0
*/
package main

import "log"

/*
  Cách 1:
    Cách giải đơn giản là for qua và vét cạn tất cả
    các case có thể so sánh và + vào result.
    Tuy nhiên cách này sẽ làm độ phức tạp thuật toán
    tang cao
  Độ phức tạp thuật toán: O(n^2)
  Độ phức tạp sử dụng bộ nhớ: O(1)
*/
func numIdenticalPairs(nums []int) int {
	result := 0

	/*
	   Giải thích:
	     Phải `len(nums) - 1` vì chạy đến cuối chả có so với ai hết
	     không cần thiết phải chạy
	*/
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				result++
			}
		}
	}

	return result
}

func main() {
	log.Println("numIdenticalPairs([]int{1,2,3,1,1,3})", numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))

	log.Println("numIdenticalPairs([]int{1,1,1,1})", numIdenticalPairs([]int{1, 1, 1, 1}))

	log.Println("numIdenticalPairs([]int{1,2,3})", numIdenticalPairs([]int{1, 2, 3}))
}
