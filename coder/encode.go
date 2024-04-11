package coder

import "slices"

func Encode(seq []byte) []byte {
	slen := len(seq)
	var fcol, lcol []byte
	fcol = make([]byte, slen)
	copy(fcol, seq)
	lcol = make([]byte, slen)
	slices.Sort(fcol)

	var pos int

	// сделать счетчик для одинаковых букв
	// когда нашел одну букву, след такая же должна начинаться с той же позиции
	for i := 0; i < slen; i++ {
		if i == 0 || fcol[i] != fcol[i-1] {
			pos = slen
		}
		for j := pos - 1; j >= 0; j-- {
			if seq[j] == fcol[i] {
				prev := j - 1
				if j-1 < 0 {
					prev = slen - 1
				}
				lcol[i] = seq[prev]
				pos = prev + 1
				break
			}
		}
	}

	return lcol
}
