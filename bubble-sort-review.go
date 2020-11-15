package main

import "fmt"

// BubbleSort sorts the given numbers in place
func BubbleSort2(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				Swap2(nums, j)
			}
		}
	}
}

// Swap nums[i], nums[i+1]
func Swap2(nums []int, i int) {
	//nums[i], nums[i+1] = nums[i+1], nums[i]
	temp := nums[i]
	nums[i] = nums[i+1]
	nums[i+1] = temp
}

func main() {
	// user enters upto 10 int
	fmt.Print("Enter input count (less than or equal to 10): ")
	var count int
	var num int
	nums := make([]int, 0, 10)

	fmt.Scanln(&count)
	fmt.Println()

	for i := 0; i < count; i++ {
		fmt.Print("Enter number and press 'Enter' key: ")
		fmt.Scanln(&num)
		nums = append(nums, num)
	}

	BubbleSort2(nums)
	fmt.Println(nums)
}
