package day3

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func Execute() {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	partOne(file)
}

type claim struct {
	ID     string
	StartX int64
	StartY int64
	SizeW  int64
	SizeH  int64
}

func partOne(file *os.File) {
	reader := bufio.NewReader(file)
	var err error

	squareMap := map[int64]map[int64]int64{}
	claims := []claim{}
	for {
		// Read the file by line.
		line, err := reader.ReadString('\n')

		// Take off the \n.
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			claimParts := strings.Split(line, " ")
			claimID := claimParts[0]
			claimStart := claimParts[2]
			claimStart = strings.Trim(claimStart, ":")
			claimSize := claimParts[3]

			//log.Printf("ID: %s, Start: %s, Size: %s", claimID, claimStart, claimSize)

			claimStartParts := strings.Split(claimStart, ",")
			claimSizeParts := strings.Split(claimSize, "x")

			claimStartX, err := strconv.ParseInt(claimStartParts[0], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			claimStartY, err := strconv.ParseInt(claimStartParts[1], 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			claimSizeWidth, err := strconv.ParseInt(claimSizeParts[0], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			claimSizeHeight, err := strconv.ParseInt(claimSizeParts[1], 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			claims = append(claims, claim{
				ID:     claimID,
				StartX: claimStartX,
				StartY: claimStartY,
				SizeW:  claimSizeWidth,
				SizeH:  claimSizeHeight,
			})

			for i := claimStartX; i < claimStartX+claimSizeWidth; i++ {
				for i2 := claimStartY; i2 < claimStartY+claimSizeHeight; i2++ {
					if _, ok := squareMap[i]; !ok {
						squareMap[i] = map[int64]int64{}
					}

					if _, ok := squareMap[i][i2]; !ok {
						squareMap[i][i2] = 0
					}

					squareMap[i][i2]++

					//log.Printf("%d,%d", i, i2)
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

	overlap := 0
	for _, yMap := range squareMap {
		for _, count := range yMap {
			//log.Printf("%d,%d: %d", x, y, count)
			if count > 1 {
				overlap++
			}
		}
	}

	log.Printf("Answer for Day 3, Part 1: %d inches are in two or more claims.", overlap)

	for _, claim := range claims {
		overlapped := false
		for i := claim.StartX; i < claim.StartX+claim.SizeW; i++ {
			breakme := false

			for i2 := claim.StartY; i2 < claim.StartY+claim.SizeH; i2++ {
				if squareMap[i][i2] > 1 {
					breakme = true
					break
				}
			}

			if breakme {
				overlapped = true
				break
			}
		}

		if !overlapped {
			log.Printf("Answer for Day 3, Part 2: ID %s does not overlap.", claim.ID)
		}
	}
}
