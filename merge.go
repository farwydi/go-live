package main

// Механизм селекции

func inArray(b int, arr [GenomeHalfSize]int) bool {
	for _, a := range arr {
		if a == b {
			return true
		}
	}

	return false
}

func Merge(genome1, genome2 Genome) Genome {

	// Создаём список длиной половиной генома
	// И заполняем его случайными номерами адресов генома
	// Не повторяются
	var AGen [GenomeHalfSize]int
	for i := 0; i < GenomeHalfSize; i++ {

		NGen := RandomGenPosition()
		for inArray(NGen, AGen) {
			NGen = RandomHalfGenPosition()
		}

		AGen[i] = NGen
	}

	genome := genome1
	for _, i := range AGen {
		genome[i] = genome2[i]
	}

	return genome
}
