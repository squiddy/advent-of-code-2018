package main

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
)

type Fabric [][]int64

type Claim struct {
	id     int64
	x      int64
	y      int64
	width  int64
	height int64
}

func claimArea(fabric *Fabric, claim Claim) {
	for dy := claim.y; dy < claim.y+claim.height; dy++ {
		for dx := claim.x; dx < claim.x+claim.width; dx++ {
			idx := dy*1000 + dx
			(*fabric)[idx] = append((*fabric)[idx], claim.id)
		}
	}
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	claims := make([]Claim, 0)

	pattern := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		matches := pattern.FindStringSubmatch(line)
		id, _ := strconv.ParseInt(matches[1], 10, 64)
		x, _ := strconv.ParseInt(matches[2], 10, 64)
		y, _ := strconv.ParseInt(matches[3], 10, 64)
		width, _ := strconv.ParseInt(matches[4], 10, 64)
		height, _ := strconv.ParseInt(matches[5], 10, 64)
		claims = append(claims, Claim{id: id, x: x, y: y, width: width, height: height})
	}

	// part 1
	fabric := make(Fabric, 1000*1000)
	for _, claim := range claims {
		claimArea(&fabric, claim)
	}

	claimed := 0
	for _, claims := range fabric {
		if len(claims) > 1 {
			claimed++
		}
	}
	log.Print(claimed)

	// part 2
	ids := make(map[int64]bool)
	for _, claim := range claims {
		ids[claim.id] = true
	}
	for _, claims := range fabric {
		if len(claims) > 1 {
			for _, claimID := range claims {
				delete(ids, claimID)
			}
		}
	}
	// TODO This feels wrong, how do I get the first key of a map?
	log.Print(reflect.ValueOf(ids).MapKeys()[0])
}
