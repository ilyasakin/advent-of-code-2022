package main

import (
	"fmt"
	"strconv"
	"strings"
)

func max(numbers []int) int {
	max := numbers[0]
	for _, v := range numbers {
		if v > max {
			max = v
		}
	}
	return max
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}

	return result
}

func main() {
	totalInventory := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	var separatedInventorySlice [][]int

	separatedInventory := strings.Split(totalInventory, "\n\n")

	for _, element := range separatedInventory {
		var inventory []int

		for _, count := range strings.Split(element, "\n") {
			intValue, _ := strconv.Atoi(count)
			inventory = append(inventory, intValue)
		}

		separatedInventorySlice = append(separatedInventorySlice, inventory)
	}

	var inventorySums []int

	for _, intSlice := range separatedInventorySlice {
		inventorySums = append(inventorySums, sum(intSlice))
	}

	resultString := "The elf that carries the most calories is carrying :count calories"

	fmt.Print(strings.Replace(resultString, ":count", strconv.Itoa(max(inventorySums)), -1))
}
