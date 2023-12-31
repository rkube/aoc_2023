package main

// Write a recursive function that calculates the difference between arrays.
// Return if all differences are zero

func calc_differences(array []int) int {
	// Calculate the differences between the array elements
	differences := make([]int, len(array)-1)
	for ix, _ := range differences {
		differences[ix] = array[ix+1] - array[ix]
	}

	// Check if all elements are zero
	all_zero := true
	for _, v := range array {
		if v != 0 {
			all_zero = false
		}
	}
	if all_zero {
		return 0
	} else {
		return array[len(array)-1] + calc_differences(differences)
	}
}

func calc_differences_2(array []int) int {
	// Calculate the differences between the array elements
	differences := make([]int, len(array)-1)
	for ix, _ := range differences {
		differences[ix] = array[ix+1] - array[ix]
	}

	// Check if all elements are zero
	all_zero := true
	for _, v := range array {
		if v != 0 {
			all_zero = false
		}
	}
	if all_zero {
		return 0
	} else {
		return array[0] - calc_differences_2(differences)
	}
}
