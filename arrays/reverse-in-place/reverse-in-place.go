package reverseinplace

// ReverseInPlace reverses the elements from start to end in place
func ReverseInPlace(list []int, start int, end int) {

	//? Approach-1 (Works, but too many lines so, O(7n))
	//? How to do in O(n)

	// for index := range list {

	// 	// base cases
	// 	if index < start {
	// 		continue
	// 	} else if start >= end {
	// 		break
	// 	}

	// 	// reversing elements
	// 	list[start], list[end] = list[end], list[start]

	// 	// increment start and decrement end
	// 	start++
	// 	end--
	// }

	//? Approach-2 (number of iterations cut down)

	for i := start; i < end; i++ {
		list[i], list[end] = list[end], list[i]
		end--
	}
}
