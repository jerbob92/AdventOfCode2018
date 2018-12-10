package day7

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func Execute() {
	file, err := os.Open("day7/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	partOne(file)
	partTwo(file)
}

type Step struct {
	ID           string
	Requirements []*Step
	Done         bool

	// Only used in step 2.
	Busy bool
}

func partOne(file *os.File) {
	var err error
	reader := bufio.NewReader(file)

	// Save all steps.
	// We use references for everything so we have 1 copy of every step.
	steps := map[string]*Step{}
	for {
		// Read the file by line.
		line, err := reader.ReadString('\n')

		// Take off the \n.
		line = strings.TrimSpace(line)
		if len(line) > 0 {

			// Extract step name and requirement for step.
			line_parts := strings.Split(line, " ")
			step := line_parts[7]
			requires := line_parts[1]

			// Create step if it doesn't exist yet.
			if _, ok := steps[step]; !ok {
				steps[step] = &Step{
					ID:           step,
					Requirements: []*Step{},
				}
			}

			// Create requirement step if it doesn't exist yet.
			if _, ok := steps[requires]; !ok {
				steps[requires] = &Step{
					ID:           requires,
					Requirements: []*Step{},
				}
			}

			// Add reference to the requirement step to the step.
			steps[step].Requirements = append(steps[step].Requirements, steps[requires])
		}

		if err != nil {
			break
		}
	}

	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	instructions := ""

	// Loop through steps until all steps are completed.
	for {
		availableSteps := []string{}

		// Loop through all steps.
		for _, step := range steps {

			// Step has already been done.
			if step.Done {
				continue
			}

			// Step has no requirements so it's available.
			if len(step.Requirements) == 0 {
				availableSteps = append(availableSteps, step.ID)
				continue
			}

			// Check whether requirements of steps are fulfilled.
			allRequirementsFulfilled := true
			for _, requirement := range step.Requirements {
				if !requirement.Done {
					allRequirementsFulfilled = false
					break
				}
			}

			// When all requirements are fulfilled, the step is avialable.
			if allRequirementsFulfilled {
				availableSteps = append(availableSteps, step.ID)
			}
		}

		// When no available steps, we are done.
		if len(availableSteps) == 0 {
			break
		}

		// Sort the available steps, lowest one is done first.
		sort.Strings(availableSteps)

		// Take first step.
		currentStep := availableSteps[0]

		// Add step to instructions.
		instructions += currentStep

		// Mark step as done.
		steps[currentStep].Done = true
	}

	log.Printf("Answer for Day 7, Part 1: The order of instructions is %s.", instructions)
}

type Worker struct {
	ID          string
	CurrentStep string
	SecondsLeft int
}

func partTwo(file *os.File) {
	var err error
	reader := bufio.NewReader(file)

	// Save all steps.
	// We use references for everything so we have 1 copy of every step.
	steps := map[string]*Step{}
	for {
		// Read the file by line.
		line, err := reader.ReadString('\n')

		// Take off the \n.
		line = strings.TrimSpace(line)
		if len(line) > 0 {

			// Extract step name and requirement for step.
			line_parts := strings.Split(line, " ")
			step := line_parts[7]
			requires := line_parts[1]

			// Create step if it doesn't exist yet.
			if _, ok := steps[step]; !ok {
				steps[step] = &Step{
					ID:           step,
					Requirements: []*Step{},
				}
			}

			// Create requirement step if it doesn't exist yet.
			if _, ok := steps[requires]; !ok {
				steps[requires] = &Step{
					ID:           requires,
					Requirements: []*Step{},
				}
			}

			// Add reference to the requirement step to the step.
			steps[step].Requirements = append(steps[step].Requirements, steps[requires])
		}

		if err != nil {
			break
		}
	}

	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	workers := []*Worker{
		{
			ID: "1",
		},
		{
			ID: "2",
		},
		{
			ID: "3",
		},
		{
			ID: "4",
		},
		{
			ID: "5",
		},
	}

	totalSeconds := 0
	for {
		availableWorkers := []*Worker{}

		// Recalculate workers.
		for _, worker := range workers {
			if worker.CurrentStep != "" {
				// Worker is busy, remove one second.
				worker.SecondsLeft--

				// No seconds left, mark worker as available again.
				if worker.SecondsLeft == 0 {

					// MArk step as done.
					steps[worker.CurrentStep].Done = true
					steps[worker.CurrentStep].Busy = false
					worker.CurrentStep = ""
					availableWorkers = append(availableWorkers, worker)
				}
			} else {
				availableWorkers = append(availableWorkers, worker)
			}
		}

		// No workers available. Go to next second.
		if len(availableWorkers) == 0 {
			logCurrentSecond(totalSeconds, workers)
			totalSeconds++
			continue
		}

		availableSteps := []string{}

		// Loop through all steps.
		for _, step := range steps {

			// Step has already been done or is being worked on.
			if step.Done || step.Busy {
				continue
			}

			// Step has no requirements so it's available.
			if len(step.Requirements) == 0 {
				availableSteps = append(availableSteps, step.ID)
				continue
			}

			// Check whether requirements of steps are fulfilled.
			allRequirementsFulfilled := true
			for _, requirement := range step.Requirements {
				if !requirement.Done {
					allRequirementsFulfilled = false
					break
				}
			}

			// When all requirements are fulfilled, the step is avialable.
			if allRequirementsFulfilled {
				availableSteps = append(availableSteps, step.ID)
			}
		}

		// When no available steps and all workers are available, we are done.
		if len(availableSteps) == 0 && len(availableWorkers) == len(workers) {
			logCurrentSecond(totalSeconds, workers)
			totalSeconds++
			break
		}

		// Sort the available steps, lowest one is done first.
		sort.Strings(availableSteps)

		// Loop through available workers.
		for i, availableWorker := range availableWorkers {
			// Check when worker count is more than available steps.
			if i > len(availableSteps) - 1 {
				break
			}

			// Mark step as busy.
			steps[availableSteps[i]].Busy = true

			// Set worker to work.
			availableWorker.CurrentStep = availableSteps[i]
			availableWorker.SecondsLeft = 60 + int(availableSteps[i][0] - 64)
		}

		logCurrentSecond(totalSeconds, workers)
		totalSeconds++
	}

	// Don't count the last second.
	totalSeconds--
	log.Printf("Answer for Day 7, Part 2: The total amount of seconds is %d.", totalSeconds)
}

func logCurrentSecond(totalSeconds int, workers []*Worker) {
	if totalSeconds == 0 {
		//log.Println("Second   Worker 1   Worker 2   Worker 3   Worker 4   Worker 5")
	}
	//log.Printf("   %d        %s          %s        %s          %s        %s", totalSeconds, workers[0].CurrentStep, workers[1].CurrentStep, workers[2].CurrentStep, workers[3].CurrentStep, workers[4].CurrentStep)
}