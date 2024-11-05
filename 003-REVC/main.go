package main

import (
	"fmt"
	"os"
	"strings"
)

func ReverseComplement(sequence string) string {
	// Initialize a builder to store the result in.
	builder := strings.Builder{}

	// Iterate over the sequence in reverse, taking the complement
	// of each base.
	for i := len(sequence) - 1; i >= 0; i-- {
		switch sequence[i] {
		case 'A':
			builder.WriteRune('T')
		case 'T':
			builder.WriteRune('A')
		case 'G':
			builder.WriteRune('C')
		case 'C':
			builder.WriteRune('G')
		}
	}

	return builder.String()
}

func main() {
	// Read in the file as a []byte.
	b, err := os.ReadFile("rosalind_revc.txt")
	if err != nil { // Don't forget to handle your errors!
		panic(err)
	}

	// Convert []byte to a string and pass it to our function.
	seq := string(b)
	revc := ReverseComplement(seq)

	// Now we just print out the result!
	fmt.Println(revc)
}
