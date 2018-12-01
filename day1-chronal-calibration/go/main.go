package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readFrequencyDrifts() ([]int64, error) {
	file, err := os.Open("../input.txt")
	if err != nil {
		return nil, err
	}

	result := make([]int64, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			return nil, err
		}

		result = append(result, value)
	}

	return result, nil
}

func getFinalFrequency(drifts []int64) int64 {
	var frequency int64
	frequency = 0
	for _, drift := range drifts {
		frequency += drift
	}

	return frequency
}

func getFirstDuplicateFrequency(drifts []int64) int64 {
	seen := map[int64]bool{}

	var frequency int64
	frequency = 0
OUTER:
	for true {
		for _, drift := range drifts {
			frequency += drift
			if seen[frequency] {
				// TODO Initially I just returned frequency here, but then I *have* to
				// return something at the end of the function, which should never be
				// reached. Is this the best solution?
				break OUTER
			}
			seen[frequency] = true
		}
	}

	return frequency
}

func main() {
	drifts, err := readFrequencyDrifts()
	if err != nil {
		log.Fatal(err)
	}

	println(getFinalFrequency(drifts))
	println(getFirstDuplicateFrequency(drifts))
}
