package coder

import (
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
	var pos int
	var groupCount int
	var last = slen - 1
	// сделать счетчик для одинаковых букв
	// когда нашел одну букву, след такая же должна начинаться с той же позиции
	for i := 0; i < slen; i++ {
		if i == 0 || i == slen-1 || fcol[i] != fcol[i-1] {
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
			if seq[j] == fcol[i] {
				cycled[groupCount] = slices.Concat(seq[j:], seq[:j])
				groupCount++
				pos = j + 1
				break
			}
		}
	}

	/*// получение таблицы сдвигов
	for i := 1; i < slen; i++ {
		cycled[i] = make([]byte, slen)
		j := i
		m := 0
		for m < slen {
			if j == slen {
				j = 0
			}
			cycled[i][m] = seq[j]
			m++
			j++
		}
	}

	// лексикографическая сортировка
	slices.SortFunc(cycled, func(a, b []byte) int {
		for i := range a {
			if a[i] > b[i] {
				return 1
			} else if a[i] < b[i] {
				return -1
			}
		}
		return 0
	})

	// получить первый и последний столбец
	// и найти номер исходной строки
	last := slen - 1
	var n int
	for i := 0; i < slen; i++ {
		lcol[i] = cycled[i][last]
		if cycled[i][0] == seq[0] && cycled[i][last] == seq[last] {
			if bytes.Compare(cycled[i], seq) == 0 {
				n = i
			}
		}
	}*/
	return lcol, n
}
