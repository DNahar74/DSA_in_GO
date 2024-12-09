package zerosumtriplets

// ZeroSumTriplets returns triplets where sum = 0
func ZeroSumTriplets(list []int) [][]int {
	//? Approach-1 (make 3 nested for loops)

	// result := make([][]int, 0)

	// for i:=0; i<len(list); i++ {
	// 	for j:=i+1; j<len(list); j++ {
	// 		if list[i] > list[j] {
	// 			list[i], list[j] = list[j], list[i]
	// 		}
	// 	}
	// }

	// for i := 0; i < len(list); i++ {
	// 	for j := i + 1; j < len(list); j++ {
	// 		for k := j + 1; k < len(list); k++ {
	// 			if list[i]+list[j]+list[k] == 0 {
	// 				e := 0
	// 				triplet := []int{list[i], list[j], list[k]}
	// 				for _, exists := range result {
	// 					if exists[0] == triplet[0] && exists[1] == triplet[1] && exists[2] == triplet[2] {
	// 						e++
	// 					}
	// 				}
	// 				if e != 0 {
	// 					continue
	// 				}
	// 				result = append(result, []int{list[i], list[j], list[k]})
	// 			}
	// 		}
	// 	}
	// }

	// return result

	//? Approach-2 (Implement a QuickSort-like approach where, each time you take a pivot element and check each case of left, right pairings)

	result := make([][]int, 0)

	// Sorting the list

	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}

	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			sum := list[i] + list[j]
			for k := j + 1; k < len(list); k++ {
				if sum+list[k] == 0 {
					// For any pair of elements, the sum=0 only for one number
					result = append(result, []int{list[i], list[j], list[k]})
					break
				}
				if sum+list[k] > 0 {
					// Since, the list is sorted, list[k++] is always > list[k]
					break
				}
			}
		}
	}

	// Check if there are any triplets

	if len(result) == 0 {
		return result
	}

	// Check for duplicates using hashmap

	triplets := [][]int{result[0]}

	for i := 0; i < len(result); i++ {
		c := 0
		for j := 0; j < len(triplets); j++ {
			if result[i][0] == triplets[j][0] && result[i][1] == triplets[j][1] && result[i][2] == triplets[j][2] {
				c++
				break
			}
		}
		if c == 0 {
			triplets = append(triplets, result[i])
		}
	}

	return triplets
}
