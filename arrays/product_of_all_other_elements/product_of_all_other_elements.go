package productofallotherelements

// ProductOfAllOtherElements takes A returns B, where B[i] = product of all terms except i
// It should be in O(n) without using division
func ProductOfAllOtherElements(list []int) []int {
	//? Approach-1: time - O(n^2) | space - O(n)

	//* Result is used to store the result because if list is altered, then the values passed in the loops change
	//* Looping through the first time to access each index,
	//* then a second time to get the product, if the index is equal it is not considered

	// result := make([]int, len(list))

	// for i := 0; i < len(list); i++ {
	// 	product := 1
	// 	for j := 0; j < len(list); j++ {
	// 		if j == i {
	// 			continue
	// 		}
	// 		product *= list[j]
	// 	}
	// 	result[i] = product
	// }

	// return result

	//? Approach-2: time - O(n) | space - O(1)

	//* Product is used to store the product of all the terms
	//* Looping through the list to replace each index with the product of all terms except itself

	product := 1

	for _, v := range list {
		product *= v
	}

	for i := 0; i < len(list); i++ {
		list[i] = product/list[i]
	}

	return list
}