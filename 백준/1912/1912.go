/**
부분합 중 최대를 묻는 문제.
다이나믹 프로그래밍으로 쉽게 해결 가능하다.

여기에서 fmt.Scanf를 사용하면 시간 초과가 뜬다.
bufio 패키지를 이용해 입력을 받으면 해결됨.
 */
package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MAX_NUMBERS_ARRAY_SIZE = 100000
)

var numbers, maxSubsetSums [MAX_NUMBERS_ARRAY_SIZE]int
var inputSize int

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &inputSize)
	for i := 0; i < inputSize; i++ {
		fmt.Fscan(reader, &numbers[i])
	}

	answer := getAnswer()
	fmt.Printf("%d", answer)
}

func getAnswer() int {
	maxSubsetSums[0] = numbers[0]
	answer := maxSubsetSums[0]
	for i := 1; i < inputSize; i++ {
		maxSubsetSums[i] = maxInt(maxSubsetSums[i-1]+numbers[i], numbers[i])
		answer = maxInt(maxSubsetSums[i], answer)
	}

	return answer
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
