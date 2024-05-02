package bwcoder

import (
	"fmt"
	"math"
	"strings"
)

func GetBin(num, bitCount int) string {
	var numBit = 1
	if num != 0 {
		numBit = int(math.Log2(float64(num))) + 1
	}
	zeroCount := bitCount - numBit
	return strings.Repeat("0", zeroCount) + fmt.Sprintf("%b", num)
}

func GetDec(bnum []byte, bitSize int) int {
	var result int
	for i := bitSize - 1; i >= 0; i-- {
		result += 1 << (bitSize - i - 1) * int(bnum[i]-48) // 0 или 1
	}

	return result
}
