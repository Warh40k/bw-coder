package coder

import (
	"slices"
)

func Decode(lcol []byte, slen, n int) []byte {
	var fcol = make([]byte, slen)
	copy(fcol, lcol)
	slices.Sort(fcol)

	var i = n
	var f byte
	var seq = make([]byte, 0, slen)

	for k := 0; k < slen; k++ {
		f = fcol[i]
		seq = append(seq, f)
		var p int
		// поиск порядка символа
		for i-p-1 >= 0 && fcol[i] == fcol[i-p-1] {
			p++
		}

		// поиск символа в последней колонке
		for j := 0; j < slen; j++ {
			if lcol[j] == f {
				if p != 0 {
					p--
				} else {
					i = j
					break
				}
			}
		}

	}

	return seq
}
