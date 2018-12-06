package day1

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func Execute() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	partOne(file)
	partTwo(file)
}

func partOne(file *os.File) {
	frequency := int64(0)
	reader := bufio.NewReader(file)
	var err error
	for {

		// Read the file by line.
		line, err := reader.ReadString('\n')

		// Take off the \n.
		line = strings.TrimSpace(line)
		if len(line) > 0 {

			// Parse the frequency change into an integer.
			change, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			// Apply the frequency change.
			frequency += change
		}

		if err != nil {
			break
		}
	}

	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	log.Printf("Answer for Day 1, Part 1: Resulting frequency is %d", frequency)
}

func partTwo(file *os.File) {
	file.Seek(0, 0)
	seenFrequencies := map[int64]bool{}
	var reachedTwice *int64
	var err error
	frequency := int64(0)
	reader := bufio.NewReader(file)
	for {
		// Read the file by line.
		line, err := reader.ReadString('\n')

		// Take off the \n.
		line = strings.TrimSpace(line)
		if len(line) > 0 {

			// Parse the frequency change into an integer.
			change, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			// Apply the frequency change.
			frequency += change

			// Check whether we reached the frequency twice already.
			if _, ok := seenFrequencies[frequency]; !ok {
				seenFrequencies[frequency] = true
			} else {
				// We have already seen this frequency, save and break out.
				currentFrequency := frequency
				reachedTwice = &currentFrequency
				break
			}
		}

		if err != nil {
			// Reset the file pointer when we didn't reach a frequency twice yet.
			if reachedTwice == nil && err == io.EOF {
				file.Seek(0, 0)
				continue
			}
			break
		}
	}

	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	log.Printf("Answer for Day 1, Part 2: First frequency seen twice is %d", *reachedTwice)
}
