package toyrobot

import (
	"errors"
	"fmt"
)

type Robot struct {
	Pos      *Position
	IsPlaced bool
}

func (robot *Robot) Left() error {
	if robot == nil {
		return errors.New("Invalid robot object")
	}

	return robot.Pos.Rotate(-1)
}

func (robot *Robot) Right() error {
	if robot == nil {
		return errors.New("Invalid robot object")
	}

	return robot.Pos.Rotate(1)
}

func (robot *Robot) Move() error {
	if robot == nil {
		return errors.New("Invalid robot object")
	}

	switch robot.Pos.Direction {
	case NORTH:
		return robot.Pos.MoveNorth()
	case EAST:
		return robot.Pos.MoveEast()
	case SOUTH:
		return robot.Pos.MoveSouth()
	case WEST:
		return robot.Pos.MoveWest()
	default:
		// do nothing
	}

	return nil
}

func (robot *Robot) Report() {
	fmt.Printf("Output: %v,%v,%v\n", robot.Pos.X, robot.Pos.Y, directionMap[robot.Pos.Direction])
}
