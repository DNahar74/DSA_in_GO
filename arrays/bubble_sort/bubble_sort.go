package bubblesort

// BubbleSort sorts the input array in ascending order using the bubble sort algorithm.
func BubbleSort(input []int) {
	//? Approach: time - O(n^2) | space - O(1)

	//* The algorithm repeatedly steps through the list, compares adjacent elements and swaps them if they are in the wrong order.
	//* Basically, in each loop the largest element is bubbled to the end 
	//* and the next loop's size is reduced by one so that it is not a part of it again
	//* The pass through the list is repeated until the list is sorted.
	//* Best case scenario: O(n) time complexity.
	//* Average case scenario: O(n^2) time complexity.

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input)-1-i; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
}
