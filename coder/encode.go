package coder

import (
	"fmt"
	"slices"
)

// составить таблицу сдвигов исходной строки
// отсортировать в лексикографическом порядке
// извлечь последние буквы каждого сдвига - получим последний столбец
// как-то определить номер исходной строки - видимо, путем полного сравнения :(
// но можно сравнивать не все сдвиги, а только те, что начинаются и заканчиваются на ту же букву
func Encode(seq []byte) ([]byte, int) {
	slen := len(seq)
	var fcol, lcol []byte
	var cycled = make([][]byte, slen)
	lcol = make([]byte, slen)
	fcol = make([]byte, slen)
	copy(fcol, seq)
	slices.Sort(fcol)

	var n int
	var i int
	var pos int
	var groupCount int
	var last = slen - 1
	// сделать счетчик для одинаковых букв
	// когда нашел одну букву, след такая же должна начинаться с той же позиции
	for i = 0; i < slen; i++ {
		if i != 0 && fcol[i] != fcol[i-1] {
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
			}
			pos = 0
			groupCount = 0
			clear(cycled)
		}

		for j := pos; j <= slen; j++ {
			if j == slen {
				j = 0
			}
			if fcol[i] == '_' {
				fmt.Println()
			}
			if seq[j] == fcol[i] {
				cycled[groupCount] = slices.Concat(seq[j:], seq[:j])
				groupCount++
				pos = j + 1
				break
			}
		}
	}

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
	}

	return lcol, n
}
