package day2

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func Execute() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	partOne(file)
	partTwo(file)
}

func partOne(file *os.File) {
	reader := bufio.NewReader(file)
	var err error

	seenTwice := 0
	seenThrice := 0

	for {

		// Read the file by line.
		line, err := reader.ReadString('\n')

		// Take off the \n.
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			seenChars := map[string]int64{}

			for _, char := range line {
				if _, ok := seenChars[string(char)]; !ok {
					seenChars[string(char)] = 0
				}

				seenChars[string(char)] += 1
			}

			hadTwice := false
			hadTrice := false
			for char, seenCount := range seenChars {
				if seenCount < 2 {
					continue
				} else if seenCount == 2 {
					if !hadTwice {
						seenTwice += 1
						hadTwice = true
					}
					//log.Printf("Saw char %s 2 times", char)
				} else if seenCount == 3 {
					if !hadTrice {
						seenThrice += 1
						hadTrice = true
					}
					//log.Printf("Saw char %s 3 times", char)
				} else {
					log.Printf("Saw char %s more than 3 times", char)
				}
			}
		}

		if err != nil {
			break
		}
	}

	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	log.Printf("Answer for Day 2, Part 1: Saw %d characters twice, saw %d characters trice, so checksum is %d * %d = %d", seenTwice, seenThrice, seenTwice, seenThrice, seenTwice * seenThrice)
}

func partTwo(file *os.File) {
	file.Seek(0, 0)
	reader := bufio.NewReader(file)
	var err error

	allLines := []string{}
	for {

		// Read the file by line.
		line, err := reader.ReadString('\n')

		// Take off the \n.
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			allLines = append(allLines, line)
		}

		if err != nil {
			break
		}
	}

	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	match1 := ""
	match2 := ""
	difference := ""
	differencePosition := 0

	// Bruteforce matches.
	for i, line := range allLines {
		for i2, line2 := range allLines {
			// Don't match own line.
			if i == i2 {
				continue
			}

			misMatches := 0
			misMatchChar := ""
			misMatchPosition := 0
			for chari, char := range line {
				if uint8(char) == line2[chari] {
					//log.Println("Matched characters")
				} else {
					misMatches++
					misMatchChar = string(char)
					misMatchPosition = chari
					if misMatches > 1 {
						//log.Println("More than 1 mismatch")
						break
					}
				}
			}

			if misMatches < 2 {
				match1 = line
				match2 = line2
				difference = misMatchChar
				differencePosition = misMatchPosition
				break
			}
		}
	}

	answer := []rune(match1)
	answer = append(answer[:differencePosition], answer[differencePosition+1:]...)

	log.Printf("Answer for Day 2, Part 2: Left match: %s, right match: %s, difference character: %s, answer: %s", match1, match2, difference, string(answer))
}
