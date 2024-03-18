package model

import (
	"fmt"
	"strings"
)

type Rover struct {
	x, y       int
	minX, maxX int
	minY, maxY int
	heading    string
}

func NewRover(x, y, mX, mY int, heading string) Rover {
	return Rover{
		minX: 0, minY: 0,
		maxX: mX, maxY: mY,
		x: x, y: y,
		heading: strings.ToUpper(heading),
	}
}

func (r *Rover) TurnLeft() {
	switch r.heading {
	case "N":
		r.heading = "W"
	case "E":
		r.heading = "N"
	case "S":
		r.heading = "E"
	case "W":
		r.heading = "S"
	}
}

func (r *Rover) TurnRight() {
	switch r.heading {
	case "N":
		r.heading = "E"
	case "E":
		r.heading = "S"
	case "S":
		r.heading = "W"
	case "W":
		r.heading = "N"
	}
}

func (r *Rover) Move() {
	switch r.heading {
	case "N":
		if r.maxY == r.y {
			return
		}
		r.y++
	case "E":
		if r.maxX == r.x {
			return
		}
		r.x++
	case "S":
		if r.minY == r.y {
			return
		}
		r.y--
	case "W":
		if r.minX == r.x {
			return
		}
		r.x--
	}
}

func (r Rover) String() string {
	return fmt.Sprintf("%d %d %s", r.x, r.y, r.heading)
}
