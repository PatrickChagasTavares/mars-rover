package main

import (
	"fmt"

	"github.com/patrickchagastavares/rover/internal/domain/usecase"
)

func main() {

	rovers, err := usecase.InstructionRover()
	if err != nil {
		panic(err)
	}

	// Print final positions of rovers
	for _, rover := range rovers {
		fmt.Println(rover)
	}
}
