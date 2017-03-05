package toyrobot

import (
	"errors"
)

type DIRECTION int

const MAX_DIRECTION int = 4

const (
	NORTH DIRECTION = 0 + iota
	EAST
	SOUTH
	WEST
)

var directionMap = map[DIRECTION]string{
	NORTH: "NORTH",
	EAST:  "EAST",
	SOUTH: "SOUTH",
	WEST:  "WEST",
}

func CheckDirection(arg string) (int, error) {
	for key, value := range directionMap {
		if value == arg {
			return int(key), nil
		}
	}

	return -1, errors.New("Invalid Direction parameter")
}

func RotateDirection(dir DIRECTION, step int) int {
	direction := int(dir) + step
	if direction < 0 {
		direction = MAX_DIRECTION - 1
	} else {
		direction = direction % MAX_DIRECTION
	}

	return direction
}
