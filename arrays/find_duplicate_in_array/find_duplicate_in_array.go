package findduplicateinarray

// FindDuplicate finds the duplicate element in an array
func FindDuplicate(list []int) int {
	//? Approach-1 (make 2 for loops to check each pair of elements in the array and return as soon as a duplicate is found)

	// for i := 0; i < len(list); i++ {
	// 	for j := i+1; j < len(list); j++ {
	// 		if list[i] == list[j] {
	// 			return list[i]
	// 		}
	// 	}
	// }

	// return -1

	//? Approach-2 (make a count array and return as soon as a duplicate is found)

	// // to check for an empty array
	// if len(list)==0 {
	// 	return -1
	// }

	// max := list[0]

	// for _, value := range list {
	// 	if value > max {
	// 		max = value
	// 	}
	// }

	// // count array is initialized with all values = 0
	// count := make([]int, max+1)

	// for _, value := range list {
	// 	if count[value] > 0 {
	// 		return value
	// 	}
	// 	count[value]++
	// }

	// return -1

	//? Approach-3 using a hashmap

	if len(list) == 0 {
		return -1
	}

	// Initialize the array
	hashmap := make(map[int]int)

	for _, value := range list {
		_, ok := hashmap[value]
		// If value is in the array, return
		if ok == true {
			return value
		}
		// Add the field
		hashmap[value] = 1
	}

	return -1
}
