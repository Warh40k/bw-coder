package coder

import (
	"bytes"
	"slices"
)

// составить таблицу сдвигов исходной строки
// отсортировать в лексикографическом порядке
// извлечь последние буквы каждого сдвига - получим последний столбец
// как-то определить номер исходной строки - видимо, путем полного сравнения :(
// но можно сравнивать не все сдвиги, а только те, что начинаются и заканчиваются на ту же букву
func Encode(seq []byte) ([]byte, int) {
	slen := len(seq)
	var lcol []byte
	var cycled [][]byte
	lcol = make([]byte, slen)
	cycled = make([][]byte, slen)
	cycled[0] = seq

	// получение таблицы сдвигов
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
	}

	return lcol, n
}
