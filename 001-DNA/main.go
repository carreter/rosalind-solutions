package main

import (
	"fmt"
	"os"
)

func CountNucleotides(sequence string) map[rune]int {
	// Initialize count map with zero values.
	res := map[rune]int{
		'A': 0,
		'T': 0,
		'G': 0,
		'C': 0,
	}

	// Iterate over the sequence and count the bases.
	for _, base := range sequence {
		res[base] = res[base] + 1
	}

	return res
}

func main() {
	// Read in the file as a []byte.
	b, err := os.ReadFile("rosalind_dna.txt")
	if err != nil { // Don't forget to handle your errors!
		panic(err)
	}

	// Convert []byte to a string and pass it to our function.
	seq := string(b)
	freqs := CountNucleotides(seq)

	// Now we just print out the result!
	fmt.Printf("%d %d %d %d\n", freqs['A'], freqs['C'], freqs['G'], freqs['T'])
}
