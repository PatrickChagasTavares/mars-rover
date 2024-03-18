package usecase

import (
	"bufio"
	"fmt"
	"os"

	"github.com/patrickchagastavares/rover/internal/domain/model"
	"github.com/patrickchagastavares/rover/pkg/scanner"
)

func InstructionRover() ([]model.Rover, error) {
	maxX, maxY, err := readPlateauSize()
	if err != nil {
		return nil, fmt.Errorf("failed to read plateau size: %w", err)
	}

	rovers, err := readRoverData(maxX, maxY)
	if err != nil {
		return nil, fmt.Errorf("failed to read rover data: %w", err)
	}

	return rovers, nil
}

func readPlateauSize() (int, int, error) {
	var maxX, maxY int
	fmt.Println("Input the upper-right coordinates: e.g: 5 5")
	if err := scanner.ScanTerminal("%d %d\n", &maxX, &maxY); err != nil {
		return 0, 0, fmt.Errorf("failed to scan upper-right coordinates: %w", err)
	}
	return maxX, maxY, nil
}

func readRoverData(maxX, maxY int) ([]model.Rover, error) {
	var rovers []model.Rover
	bScanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Input start coordinates rover: e.g: 1 2 N")
	for bScanner.Scan() {
		line := bScanner.Text()
		if line == "" {
			break
		}

		rover, err := parseRoverData(line, maxX, maxY)
		if err != nil {
			return nil, fmt.Errorf("failed to parse rover data: %w", err)
		}

		if err := moveRover(&rover); err != nil {
			return nil, fmt.Errorf("failed to move rover: %w", err)
		}

		rovers = append(rovers, rover)
		fmt.Println("input coordinates to next rover or press enter to exit")
	}

	if err := bScanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read scanner: %w", err)
	}

	return rovers, nil
}

func parseRoverData(line string, maxX, maxY int) (model.Rover, error) {
	var x, y int
	var heading string

	if err := scanner.ScanLine(line, "%d %d %s\n", &x, &y, &heading); err != nil {
		return model.Rover{}, fmt.Errorf("failed to scan start position at rover: %w", err)
	}

	return model.NewRover(x, y, maxX, maxY, heading), nil
}

func moveRover(rover *model.Rover) error {
	var instructions string

	fmt.Println("Input instructions for rover: e.g: LMLMLMLMM")
	if err := scanner.ScanTerminal("%s\n", &instructions); err != nil {
		return fmt.Errorf("failed scan instructions: %w", err)
	}

	// Execute instructions for current rover
	for _, instruction := range instructions {
		switch instruction {
		case 'L', 'l':
			rover.TurnLeft()
		case 'R', 'r':
			rover.TurnRight()
		case 'M', 'm':
			rover.Move()
		}
	}

	return nil
}
