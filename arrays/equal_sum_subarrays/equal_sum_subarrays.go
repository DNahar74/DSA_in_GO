package equalsumsubarrays

// EqualSubArrays returns consecutive subarrays with equal sum
func EqualSubArrays(list []int) [][]int {
	//? Approach-1: time - O(?) | space - O(?)

	//* A merge-sort like approach dividing and adding the subarrays until a match is found
	//* Does not work for the given testcases

	// result := make([][]int, 0)

	// if len(list) < 2 {
	// 	return result
	// }

	// mergeAndCheck()

	// return result

	//? Approach-2: time - O(n) | space - O(n)

	//* A two-pointer approach, two sub-arrays made to store them
	//* subarr1 initialized with first element, subarr2 initialized with last element
	//* each with a sum: sum1 = sum of all elements in subarr1
	//* start = 0, end = len(list)-1
	//* while start < end {
	//* 	if (sum1 > sum2) end--; sum2+=list[end]; subarr2 = append(list[end])
	//*		similarly for sum1 < sum2
	//* }

	result := make([][]int, 0)

	if len(list) < 2 {
		return result
	}

	start, end := 0, len(list)-1

	sum1, sum2 := list[start], list[end]

	subArray1 := []int{list[start]}
	subArray2 := []int{list[end]}

	for start < end-1 {
		if sum1 == sum2 {
			// Add an element on either side
			// If this was 2, anelment gets repeated, take {1,2,2,3} for example and dry run
			if end-start >= 3 {
				end--
				sum2 += list[end]
				subArray2 = append([]int{list[end]}, subArray2...)
				start++
				sum1 += list[start]
				subArray1 = append(subArray1, list[start])
			} else {
				start++
				sum1 += list[start]
				subArray1 = append(subArray1, list[start])
				break
			}
		} else if sum1 > sum2 {
			end--
			sum2 += list[end]
			subArray2 = append([]int{list[end]}, subArray2...)
		} else {
			start++
			sum1 += list[start]
			subArray1 = append(subArray1, list[start])
		}
	}

	if sum1 == sum2 {
		result = append(result, subArray1)
		result = append(result, subArray2)
	}

	return result
}

// func mergeAndCheck(arr []int) []int {
// 	if len(arr)<=1 {
// 		return nil
// 	}
// 	mid := int(len(arr)/2)
// 	mergeAndCheck(arr[:mid])
// 	mergeAndCheck(arr[mid+1:])
// 	if merge(arr, mid) {
// 		return arr
// 	} else {
// 		return nil
// 	}
// }

// func merge(arr []int, mid int) bool {
// 	sum1, sum2 := 0, 0
// 	for index, value := range arr {
// 		if index <= mid {
// 			sum1 += value
// 		} else {
// 			sum2 += value
// 		}
// 	}
// 	if (sum1 == sum2) {
// 		return true
// 	}
// 	return false
// }
