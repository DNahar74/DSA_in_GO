package addtwonumbers

// AddTwoNumbers adds two numbers represented as arrays of digits.
func AddTwoNumbers(num1 []int, num2 []int) []int {
	//? Approach-1

	// // Instantiate sum array
	// var sum, small []int

	// // Check the number of bits
	// if len(num1) >= len(num2) {
	// 	sum = num1
	// 	small = num2
	// } else {
	// 	sum = num2
	// 	small = num1
	// }

	// // Initialize carry
	// carry := 0

	// // Iterate over digits from right to left
	// for i := 0; i < len(sum); i++ {
	// 	if i > len(small) {

	// 	} else if i == len(small) {
	// 		if (sum[i] + carry) >= 10 {
	// 			sum[i] = (sum[i]  + carry) % 10
	// 			carry = 1
	// 		} else {
	// 			sum[i] = sum[i] + carry
	// 			carry = 0
	// 		}
	// 		sum[i] = (sum[i]+carry) % 10
	// 	}else {

	// 		if (sum[i] + small[i] + carry) >= 10 {
	// 			sum[i] = (sum[i] + small[i] + carry) % 10
	// 			carry = 1
	// 		} else {
	// 			sum[i] = sum[i] + small[i] + carry
	// 			carry = 0
	// 		}
	// 	}

	// }

	// sum = append([]int{carry}, sum...)

	// return sum

	//? Approach-2 (make a sum array and a carry int. Also, because of 0-based indexing, len(num)-1-i gives the right to left traversal)

	var sum []int

	carry := 0

	for i := 0; ; i++ {
		if i >= len(num1) && i >= len(num2) {
			break
		} else if i >= len(num1) {
			sum = append([]int{num2[len(num2)-1-i] + carry}, sum...)
		} else if i >= len(num2) {
			sum = append([]int{num1[len(num1)-1-i] + carry}, sum...)
		} else {
			sum = append([]int{num1[len(num1)-1-i] + num2[len(num2)-1-i] + carry}, sum...)
		}
		carry = int(sum[0] / 10)
		sum[0] %= 10
	}

	if carry > 0 {
		sum = append([]int{carry}, sum...)
	}

	return sum
}
