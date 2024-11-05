package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type FASTA struct {
	Label    string
	Sequence string
}

func ParseFASTAs(rd io.Reader) ([]FASTA, error) {
	res := []FASTA{}
	label := ""
	builder := strings.Builder{}

	brd := bufio.NewReader(rd)
	for {
		b, err := brd.ReadByte()
		if err == io.EOF {
			res = append(res, FASTA{
				label,
				builder.String(),
			})
			return res, nil
		} else if err != nil {
			return nil, err
		}

		if b == '>' {
			res = append(res, FASTA{
				label,
				builder.String(),
			})
			builder.Reset()
			label, err = brd.ReadString('\n')
			if err != nil {
				return nil, err
			}
			label = strings.TrimSpace(label)
			continue
		} else if b == '\n' {
			continue
		}

		builder.WriteByte(b)
	}
}

func GCContent(sequence string) float64 {
	count := 0
	for _, base := range sequence {
		if base == 'G' || base == 'C' {
			count++
		}
	}

	return float64(count) / float64(len(sequence))
}

func main() {
	file, err := os.Open("rosalind_gc.txt")
	if err != nil {
		panic(err)
	}

	entries, err := ParseFASTAs(file)
	if err != nil {
		panic(err)
	}

	maxGCContent := 0.0
	maxGCName := ""
	for _, entry := range entries {
		if gCContent := GCContent(entry.Sequence); gCContent > maxGCContent {
			maxGCContent = gCContent
			maxGCName = entry.Label
		}
	}

	fmt.Printf("%v\n%v\n", maxGCName, maxGCContent*100)
}
