package pages

import "math/rand"

var Value = Print()

func Print() int64 {
	return rand.Int63()
}
