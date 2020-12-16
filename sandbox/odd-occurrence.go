package main

import "fmt"

/**
A non-empty zero-indexed array A consisting of N integers is given. The array contains an odd number of elements, and each
element of the array can be paired with another element that has the same value, except for one element that is left unpaired.

For example, in array A such that:

  A[0] = 9  A[1] = 3  A[2] = 9
  A[3] = 3  A[4] = 9  A[5] = 7
  A[6] = 9
the elements at indexes 0 and 2 have value 9,
the elements at indexes 1 and 3 have value 3,
the elements at indexes 4 and 6 have value 9,
the element at index 5 has value 7 and is unpaired.
Write a function:

func Solution(A []int) int
that, given an array A consisting of N integers fulfilling the above conditions, returns the value of the unpaired element.

For example, given array A such that:

  A[0] = 9  A[1] = 3  A[2] = 9
  A[3] = 3  A[4] = 9  A[5] = 7
  A[6] = 9
the function should return 7, as explained in the example above.

Assume that:

N is an odd integer within the range [1..1,000,000];
each element of array A is an integer within the range [1..1,000,000,000];
all but one of the values in A occur an even number of times.
Complexity:

expected worst-case time complexity is O(N);
expected worst-case space complexity is O(1), beyond input storage (not counting the storage required for input arguments).
Elements of input arrays can be modified.

Test Data
[ 9, 3, 9, 3, 9, 7, 9 ],
[ 42 ],
[ 4, 42, 4 ]
**/

func Solution(a []int) int {
	if len(a) == 1 {
		return a[0]
	}

	hash := make(map[int]int)
	for i := 0; i < len(a); i++ {
		key := a[i]
		hash[key] += 1 // add one to the count

		if hash[key]%2 == 0 {
			delete(hash, key)
		}
	}

	var unpaired int
	for k, v := range hash {
		if v%2 != 0 {
			unpaired = k
			break
		}
	}

	return unpaired
}

type TestData struct {
	A []int
	R int
}

func main() {
	testData := []TestData{
		TestData{[]int{9, 3, 9, 3, 9, 7, 9}, 7},
		TestData{[]int{42}, 42},
		TestData{[]int{4, 42, 4}, 42},
	}

	tests := 0
	passed := 0
	for _, data := range testData {
		tests++
		r := Solution(data.A)
		var passfail string
		if r == data.R {
			passed++
			passfail = "Ok"
		} else {
			passfail = "FAILED!"
		}
		fmt.Println(data, r, passfail)
	}

	fmt.Printf("\nRan %d Tests, %d Passed, %d Failed...\n", tests, passed, tests-passed)
}
