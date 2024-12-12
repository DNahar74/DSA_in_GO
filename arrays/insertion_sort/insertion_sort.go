package insertionsort

// InsertionSort sorts the given array in ascending order using insertion sort algorithm.
func InsertionSort(input []int) {
	//? Approach: time - O(n^2) | space - O(1)

	//* Loop through for each element, assigning a min index value=i
	//* Loop through again comparing the index i+1 with all previous indices.
	//* If the current element is smaller than the previous element, swap them
  //* If not, break the loop and continue with the next element, because all elemnts before it must be smaller
  //* Continue until the loop completes.
  //* After each iteration, the sorted part of the array increases by one element.
  //* The sorted part is checked with the remaining unsorted part.

	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if input[j] < input[j-1] {
        input[j], input[j-1] = input[j-1], input[j]
      } else {
        break
      }
		}
	}
}