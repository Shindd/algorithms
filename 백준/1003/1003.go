/**
일반적인 피보나치 문제와 기법은 동일하다. f(n) = f(n-2) + f(n-1)
0과 1의 출력횟수를 알아야 하므로
	fibonacci[N][0] 에는 0의 출력 횟수를 저장한다.
	fibonacci[N][1] 에는 1의 출력 횟수를 저장한다.
*/
package main

import (
	"fmt"
)

var fibonacci [41][2]int

func main() {
	var testCases int
	fmt.Scanf("%d", &testCases)

	prepare()

	for testCase := 0; testCase < testCases; testCase++ {
		var number int
		fmt.Scanf("%d", &number)
		zeroCount, oneCount := getAnswer(number)
		fmt.Printf("%d %d\n", zeroCount, oneCount)
	}
}

func prepare() {
	fibonacci[0][0], fibonacci[0][1] = 1, 0
	fibonacci[1][0], fibonacci[1][1] = 0, 1
	for i := 2; i < 41; i++ {
		fibonacci[i][0] = fibonacci[i-2][0] + fibonacci[i-1][0]
		fibonacci[i][1] = fibonacci[i-2][1] + fibonacci[i-1][1]
	}
}

func getAnswer(number int) (int, int) {
	return fibonacci[number][0], fibonacci[number][1]
}
