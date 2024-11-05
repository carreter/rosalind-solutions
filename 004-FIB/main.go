package main

import (
	"fmt"
)

func RabbitPairsAfter(months int, litterSize int) int {
	if months <= 2 {
		return 1
	}

	population := []int{1, 1}

	for len(population) < months {
		population = append(population, population[len(population)-1]+population[len(population)-2]*litterSize)
	}

	return population[months-1]
}

func main() {
	fmt.Println(RabbitPairsAfter(31, 3))
}
