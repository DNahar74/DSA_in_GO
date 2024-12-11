package rotateksteps

// RotateKSteps rotates the given list k steps
func RotateKSteps(list []int, k int) {
	//? Approach-1: time - O(n^2) | space - O(1)

	//* Create a loop that runs k times
	//* create a temp := last element in list
	//* Run a loop in reverse till the second element where: list[j] = list[j-1]
	//* Once, this ends, assign list[0] = temp

	// for i := 0; i < k; i++ {
	// 	temp := list[len(list)-1]
	// 	for j := len(list)-1; j > 0; j-- {
	// 		list[j] = list[j-1]
	// 	}
	// 	list[0] = temp
	// }

	//? Approach-2: time - O(n^2) | space - O(1)

	//* Just looks like O(n), it runs for n*k times
	//* create a temp := last element in list
	//* Run a loop where i goes in reverse across the list, until k == 0
	//* In each iteration, swap list[i] and list[i-1]
	//* If i == 0 then,
	//* list[i] = temp; reassign temp = last element; k--;
	//* And, i = len(list) not len(list) - 1 because there is a decrement in i value after each loop, log to check if needed

	// // To remove the case of empty array
	// if len(list) < 1 {
	// 	return
	// }

	// temp := list[len(list)-1]

	// for i := len(list) - 1; k > 0; i-- {
	// 	if i == 0 {
	// 		list[i] = temp
	// 		temp = list[len(list)-1]
	// 		i = len(list)
	// 		k--
	// 	} else {
	// 		list[i] = list[i-1]
	// 	}
	// }

	//? Approach-3: time - O(n) | space - O(n)

	//* Calculate the position for each index
	//* NewPosition = (i + k)%len(list)
	//* The array can be divided into 2 subarrays, 1 that don't change on %n and 2 that needs to circle back
	//* There will be k elements that need to be circled back list[len(list)-k:]

	if len(list) == 0 {
		return
	}

	k %= len(list)
	circleBack := append([]int{}, list[len(list)-k:]...)
	for i := len(list) - k - 1; i >= 0; i-- {
		list[i+k] = list[i]
	}

	// This is reassignment, so it is local
	// list = append(circleBack, list[k:]...)

	copy(list[:k], circleBack)
}
