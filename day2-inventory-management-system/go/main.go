package main

import (
	"bufio"
	"log"
	"os"
)

func readBoxIds() ([]string, error) {
	file, err := os.Open("../input.txt")
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}

func countLetterOccurances(id string) map[rune]int {
	counts := map[rune]int{}

	for _, char := range id {
		counts[char]++
	}

	return counts
}

func calculateChecksum(ids []string) int {
	boxesLetterTwice := 0
	boxesLetterThrice := 0

	for _, id := range ids {
		counts := countLetterOccurances(id)
		twice := false
		thrice := false
		for _, count := range counts {
			switch count {
			case 2:
				twice = true
			case 3:
				thrice = true
			}
		}
		if twice {
			boxesLetterTwice++
		}
		if thrice {
			boxesLetterThrice++
		}
	}

	return boxesLetterTwice * boxesLetterThrice
}

func getDistance(idA string, idB string) int {
	// TODO this assumes the input is ASCII
	distance := 0
	for idx := 0; idx < len(idA); idx++ {
		if idA[idx] != idB[idx] {
			distance++
		}
	}
	return distance
}

func findPrototypeBoxes(ids []string) (string, string) {
	for idx, idA := range ids {
		for _, idB := range ids[idx:] {
			if getDistance(idA, idB) == 1 {
				return idA, idB
			}
		}
	}
	// TODO Is there something like Rust's Result type?
	return "", ""
}

func getCommonLetters(idA string, idB string) string {
	// TODO this assumes the input is ASCII
	result := []byte{}
	for idx := 0; idx < len(idA); idx++ {
		if idA[idx] == idB[idx] {
			result = append(result, idA[idx])
		}
	}
	return string(result)
}

func main() {
	boxIds, err := readBoxIds()
	if err != nil {
		log.Fatal(err)
	}

	println(calculateChecksum(boxIds))
	boxA, boxB := findPrototypeBoxes(boxIds)
	println(getCommonLetters(boxA, boxB))
}
