package iteration

import "fmt"

const repeatCount = 5

func Repeat(character string) string {
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}

func AdditionalForLoops() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	// 1, 2, 3

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	// 7, 8, 9

	for {
		fmt.Println("loop")
		break
	}

	//loop
	//for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	//1, 3, 5
}
