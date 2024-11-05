package main

import (
	"fmt"
	"os"
	"strings"
)

func DNAToRNA(sequence string) string {
	// Initialize a builder to store the result in.
	builder := strings.Builder{}

	// Iterate over the sequence, replacing Ts with Us.
	for _, base := range sequence {
		if base == 'T' {
			builder.WriteRune('U')
		} else {
			builder.WriteRune(base)
		}
	}

	return builder.String()
}

func main() {
	// Read in the file as a []byte.
	b, err := os.ReadFile("rosalind_rna.txt")
	if err != nil { // Don't forget to handle your errors!
		panic(err)
	}

	// Convert []byte to a string and pass it to our function.
	seq := string(b)
	rna := DNAToRNA(seq)

	// Now we just print out the result!
	fmt.Println(rna)
}
