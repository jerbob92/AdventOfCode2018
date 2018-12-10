package day4

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Execute() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	partOne(file)
}

type guard struct {
	ID            int
	MinuteStart   int
	MinutesAsleep int
	SleepHistory  map[int]int
}

func partOne(file *os.File) {
	var err error
	reader := bufio.NewReader(file)
	rows := []string{}

	for {
		// Read the file by line.
		line, err := reader.ReadString('\n')

		// Take off the \n.
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			rows = append(rows, line)
		}

		if err != nil {
			break
		}
	}

	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	// Sort in chronological order.
	sort.Strings(rows)

	// Map to save info about the guard.
	Guards := map[string]*guard{}

	// Save the guard we are currently at.
	var currentGuard *guard

	// Loop through the events.
	for _, row := range rows {

		// Split by space to get some info on the event.
		row_parts := strings.Split(row, " ")

		// Select event.
		if strings.Contains(row, "begins shift") {

			// Guard begins shift, save or create new map entry.
			guard_id := row_parts[3]
			if val, ok := Guards[guard_id]; ok {
				currentGuard = val
			} else {
				idInt, err := strconv.Atoi(guard_id[1:])
				if err != nil {
					log.Println(err)
				}
				Guards[guard_id] = &guard{
					ID:           idInt,
					SleepHistory: map[int]int{},
				}
				currentGuard = Guards[guard_id]
			}
		} else if strings.Contains(row, "falls asleep") {
			// Guard falls asleep, save in the current guard the current minute.
			minute := row_parts[1][3 : len(row_parts[1])-1]
			minuteInt, err := strconv.Atoi(minute)
			if err != nil {
				log.Println(err)
			}
			currentGuard.MinuteStart = minuteInt
		} else if strings.Contains(row, "wakes up") {
			// Guard wakes up.
			minute := row_parts[1][3 : len(row_parts[1])-1]
			minuteInt, err := strconv.Atoi(minute)
			if err != nil {
				log.Println(err)
			}

			// Compare the saved minute with the current minute to calculate time alseep.
			// Add time aslepe to the user.
			currentGuard.MinutesAsleep += minuteInt - currentGuard.MinuteStart

			// Loop through every minute that the user was asleep and save it to the history.
			for i := currentGuard.MinuteStart; i < minuteInt; i++ {
				if _, ok := currentGuard.SleepHistory[i]; !ok {
					currentGuard.SleepHistory[i] = 0
				}
				currentGuard.SleepHistory[i]++
			}
		}
	}

	// Save which guard was asleep the most, also save which guard was asleep the most at a single minute.
	var mostAsleep *guard
	var mostAsleepOnSameMinute *guard
	mostAsleepOnSameMinuteCount := 0
	mostAsleepOnSameMinuteCountMinute := 0
	for _, Guard := range Guards {
		// Detect guard that was asleep the most.
		if mostAsleep == nil || Guard.MinutesAsleep > mostAsleep.MinutesAsleep {
			mostAsleep = Guard
		}

		// Get how much this guard was asleep on the same minute and save it.
		sameMinuteCount := 0
		sameMinuteCountMinute := 0
		for minute, count := range Guard.SleepHistory {
			if count > sameMinuteCount {
				sameMinuteCount = count
				sameMinuteCountMinute = minute
			}
		}

		// Check whether this guard was more asleep at a given minute than other guards.
		if mostAsleepOnSameMinute == nil || sameMinuteCount > mostAsleepOnSameMinuteCount {
			mostAsleepOnSameMinute = Guard
			mostAsleepOnSameMinuteCount = sameMinuteCount
			mostAsleepOnSameMinuteCountMinute = sameMinuteCountMinute
		}
	}

	// Get the minute the guard was most asleep for the guard that was asleep the most time in total.
	MinuteMostAsleep := 0
	CountMostAsleep := 0
	for minute, count := range mostAsleep.SleepHistory {
		if count > CountMostAsleep {
			CountMostAsleep = count
			MinuteMostAsleep = minute
		}
	}

	log.Printf("Answer for Day 4, Part 1: ID #%d slept the most on minute %d, so the answer is %d * %d = %d.", mostAsleep.ID, MinuteMostAsleep, mostAsleep.ID, MinuteMostAsleep, mostAsleep.ID * MinuteMostAsleep)
	log.Printf("Answer for Day 4, Part 2: ID #%d slept the most on minute %d (%d times), so the answer is %d * %d = %d.", mostAsleepOnSameMinute.ID, mostAsleepOnSameMinuteCountMinute, mostAsleepOnSameMinuteCount, mostAsleepOnSameMinute.ID, mostAsleepOnSameMinuteCountMinute, mostAsleepOnSameMinute.ID * mostAsleepOnSameMinuteCountMinute)
}
