/**
수학문제, 머리를 잘 굴리자
아마 다음번에 다시봐도 이해 못할듯...
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	var testCases int
	fmt.Scanf("%d", &testCases)

	for testCase := 0; testCase < testCases; testCase++ {
		var x, y int
		fmt.Scanf("%d %d", &x, &y)
		answer := getAnswer(x, y)
		fmt.Printf("%d\n", answer)
	}
}

func getAnswer(x int, y int) int {
	distance := y - x
	squareRoot := int(math.Sqrt(float64(distance)))

	if distance == squareRoot*squareRoot {
		return 2*squareRoot - 1
	} else if squareRoot*squareRoot < distance && distance <= squareRoot*squareRoot+squareRoot {
		return 2 * squareRoot
	}

	return 2*squareRoot + 1
}
