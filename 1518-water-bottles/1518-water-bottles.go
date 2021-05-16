/*
  Problems: https://leetcode.com/problems/water-bottles/

  Input: numBottles = 9, numExchange = 3
  Output: 13

  Input: numBottles = 5, numExchange = 5
  Output: 6

  Input: numBottles = 2, numExchange = 3
  Output: 2

  Input: numBottles = 15, numExchange = 4
  Output: 19
*/
package main

import (
	"log"
)

func numWaterBottles(numBottles int, numExchange int) int {
	if numBottles < numExchange {
		return numBottles
	}

	// Tổng số chai đã uống
	numTotalDrinkBottles := numBottles

	// Tổng số trai rỗng
	numEmptyBottles := numBottles

	for {
		// Tính số chai có nước sẽ đổi được
		numWaterBottlesCanExchange := numEmptyBottles / numExchange

		// Tiến hành đổi số chai rỗng lấy chai có nước
		// Giảm số chai nước rỗng xuống sau khi đã đổi
		numEmptyBottles = numEmptyBottles - (numWaterBottlesCanExchange * numExchange)

		// Sau đó lấy nước uống
		// Cập nhật số chai đã uống được lên
		numTotalDrinkBottles += numWaterBottlesCanExchange

		// Uống hết nước các chai đã đổi
		numEmptyBottles += numWaterBottlesCanExchange

		if numEmptyBottles < numExchange {
			return numTotalDrinkBottles
		}
	}
}

func main() {
	// log.Println("numWaterBottles(9, 3)", numWaterBottles(9, 3))
	// log.Println("numWaterBottles(5, 5)", numWaterBottles(5, 5))
	// log.Println("numWaterBottles(2, 3)", numWaterBottles(2, 3))
	log.Println("numWaterBottles(15, 4)", numWaterBottles(15, 4))
}
