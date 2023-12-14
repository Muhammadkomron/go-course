package main

import (
	"fmt"
	"math/rand"
	"time"
)

type integers []int

func createIntegersSlice() integers {
	numbers := integers{}
	// First solution
	//for i := 1; i <= 10; i++ {
	//	numbers = append(numbers, i)
	//}
	// Second solution
	//for i := range [5]int{} {
	//	numbers = append(numbers, i)
	//}
	// Third solution
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for range [10]int{} {
		numbers = append(numbers, r.Intn(10))
	}
	return numbers
}

func (i integers) checkNumbers() {
	// First solution
	//for _, number := range i {
	//	if number == 0 {
	//		fmt.Printf("%v is Neutral\n", number)
	//	} else if number%2 == 0 {
	//		fmt.Printf("%v is Even\n", number)
	//	} else {
	//		fmt.Printf("%v is Odd\n", number)
	//	}
	//}
	// Second solution
	for n := 0; n < len(i); n++ {
		if i[n] == 0 {
			fmt.Printf("%v is Neutral\n", i[n])
		} else if i[n]%2 == 0 {
			fmt.Printf("%v is Even\n", i[n])
		} else {
			fmt.Printf("%v is Odd\n", i[n])
		}
	}
}
