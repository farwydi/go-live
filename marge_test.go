package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMerge(t *testing.T) {

	rand.Seed(13)

	genome1 := RandGenomeGenerator()
	genome2 := RandGenomeGenerator()

	fmt.Println(genome1)
	fmt.Println(genome2)

	genome := Merge(genome1, genome2)
	fmt.Println(genome)

	for i, gen := range genome {
		if gen != genome1[i] {
			if gen != genome2[i] {
				t.Errorf("Genome eq")
			}
		}
	}

	if genome1 == genome {
		t.Errorf("Genome eq")
	}

	if genome2 == genome {
		t.Errorf("Genome eq")
	}

	if genome1 == genome2 {
		t.Errorf("Genome eq")
	}
}
