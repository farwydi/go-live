package main

import "math/rand"

const (
	GenomeSize     = 64
	GenomeHalfSize = 64 / 2
)

type (
	GenomeType int
	Genome     [GenomeSize]GenomeType
	GenomeHalf [GenomeHalfSize]GenomeType
)

//func GenomeDown(gemome) {
//
//}

func RandomGenPosition() int {
	return rand.Intn(GenomeSize)
}

func RandomHalfGenPosition() int {
	return rand.Intn(GenomeHalfSize)
}

func RandomGen() GenomeType {
	return GenomeType(rand.Intn(74))
}
