package coder

import (
	"fmt"
	"slices"
)

func Encode(seq, lcol []byte, slen int) int {
	var n int

	var fcol []byte
	var cycled = make([][]byte, slen)
	fcol = make([]byte, slen)
	copy(fcol, seq)
	slices.Sort(fcol)

	var i, temp, pos, groupCount int

	for i = 0; i < slen; i++ {
		if i != 0 && fcol[i] != fcol[i-1] {
			temp = getLcol(cycled, seq, lcol, groupCount, i, slen-1)
			if temp != -1 {
				n = temp
			}
			pos = 0
			groupCount = 0
		}

		for j := pos; j <= slen; j++ {
			if j == slen {
				j = 0
			}
			if fcol[i] == 'F' {
				fmt.Print()
			}
			if seq[j] == fcol[i] {
				cycled[groupCount] = slices.Concat(seq[j:slen], seq[:j])
				groupCount++
				pos = j + 1
				break
			}
		}
	}

	temp = getLcol(cycled, seq, lcol, groupCount, i, slen-1)
	if temp != -1 {
		n = temp
	}
	return n
}

func getLcol(cycled [][]byte, seq, lcol []byte, groupCount, i, last int) int {
	var n = -1
	slices.SortFunc(cycled, func(a, b []byte) int {
		if a == nil && b != nil {
			return 1
		} else if b == nil && a != nil {
			return -1
		}
		for k := range a {
			if a[k] > b[k] {
				return 1
			} else if a[k] < b[k] {
				return -1
			}
		}
		return 0
	})
	for m := 0; m < groupCount; m++ {
		if cycled[m] == nil {
			break
		}
		lcol[i-groupCount+m] = cycled[m][last]

		// проверка на исходную строку
		if cycled[m][last] == seq[last] &&
			cycled[m][0] == seq[0] &&
			isEqual(cycled[m], seq) {
			n = i - groupCount + m
		}
	}
	clear(cycled)
	return n
}

func isEqual(a, b []byte) bool {
	if a == nil && b != nil || b == nil && a != nil {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
