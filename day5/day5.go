package day5

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
)

func Execute() {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	partOne(file)
	partTwo(file)
}

func partOne(file *os.File) {
	polymerBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Remove unit reactions.
	polymerBytes = removeUnitReactions(polymerBytes)

	log.Printf("Answer for Day 5, Part 1: the remaining amount of units is %d.", len(polymerBytes))
}

func partTwo(file *os.File) {
	file.Seek(0, 0)
	allPolymerBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	shortestPolymer := 0
	shortestPolymerChar := 0

	// Loop through 97 (a) and 123 (z)
	for i := 97; i < 123; i++ {
		// Replace a-z.
		polymerBytes := bytes.Replace(allPolymerBytes, []byte{byte(i)}, []byte{}, -1)

		// Replace A-Z bytes.
		polymerBytes = bytes.Replace(polymerBytes, []byte{byte(i - 32)}, []byte{}, -1)

		// Remove unit reactions.
		polymerBytes = removeUnitReactions(polymerBytes)

		// Check whether this is the shortest polymer.
		if shortestPolymer == 0 || len(polymerBytes) < shortestPolymer {
			shortestPolymer = len(polymerBytes)
			shortestPolymerChar = i
		}
	}

	log.Printf("Answer for Day 5, Part 2: the shortest polymer was by the one by removing the units %s/%s, the length was %d.", string(shortestPolymerChar), string(shortestPolymerChar-32), shortestPolymer)
}

func removeUnitReactions(polymerBytes []byte) []byte {
	// Loop until no updates were done.
	didUpdate := true
	for didUpdate {

		// Set initial value for this loop.
		didUpdate = false
		for i, char := range polymerBytes {

			// Get the upper character and the lower character of this byte (distance between a and A is 32).
			upperChar := char - 32
			lowerChar := char + 32

			// Check whether the previous byte matches the upper or lower version of this char.
			if i > 0 && (polymerBytes[i-1] == lowerChar || polymerBytes[i-1] == upperChar) {

				// Construct a new bytes array without the matching characters.
				polymerBytes = append(polymerBytes[:i-1], polymerBytes[i+1:]...)
				didUpdate = true

				// Break out of the loop so we start over.
				break
			}
		}
	}

	return polymerBytes
}