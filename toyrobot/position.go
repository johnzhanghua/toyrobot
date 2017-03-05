package toyrobot

import (
	"errors"
	"log"
	"strconv"
)

// 5x5 unit table (0-4, 0-4)
const MAX_X int = 5
const MAX_Y int = 5

const POSITION_ARGS int = 3

type Position struct {
	X, Y      int
	Direction DIRECTION
}

func (p *Position) MoveNorth() error {
	if p == nil {
		return errors.New("Invalid position object")
	}

	if !moveY(p, 1) {
		return errors.New("Can not move to invalid North position")
	}

	return nil
}

func (p *Position) MoveEast() error {
	if p == nil {
		return errors.New("Invalid position object")
	}

	if !moveX(p, 1) {
		return errors.New("Can not move to invalid East position")
	}

	return nil
}

func (p *Position) MoveSouth() error {
	if p == nil {
		return errors.New("Invalid position object")
	}

	if !moveY(p, -1) {
		return errors.New("Can not move to invalid South position")
	}

	return nil
}

func (p *Position) MoveWest() error {
	if p == nil {
		return errors.New("Invalid position object")
	}

	if !moveX(p, -1) {
		return errors.New("Can not move to invalid West position")
	}

	return nil
}

func (p *Position) Rotate(step int) error {
	if p == nil {
		return errors.New("Invalid position object")
	}

	p.Direction = DIRECTION(RotateDirection(p.Direction, step))

	return nil
}

func CreatePosition(args []string) (*Position, error) {
	if len(args) != POSITION_ARGS {
		log.Printf("Position arguments wrong format")
		return nil, errors.New("Position arguments wrong format")
	}
	xString := args[0]
	yString := args[1]
	fString := args[2]

	x, err := strconv.Atoi(xString)
	if err != nil || !checkXPosition(x) {
		log.Printf("Invalid X parameter")
		return nil, errors.New("Invalid X parameter")
	}

	y, err := strconv.Atoi(yString)
	if err != nil || !checkYPosition(y) {
		log.Printf("Invalid Y parameter")
		return nil, errors.New("Invalid Y parameter")

	}

	dir, err := CheckDirection(fString)
	if err != nil {
		log.Printf("Invalid F paramter")
		return nil, errors.New("Invalid F parameter")
	}

	return &Position{X: x, Y: y, Direction: DIRECTION(dir)}, nil
}

func moveX(p *Position, step int) bool {
	x := p.X + step

	if !checkXPosition(x) {
		return false
	}

	p.X = x

	return true
}

func moveY(p *Position, step int) bool {
	y := p.Y + step

	if !checkYPosition(y) {
		return false
	}

	p.Y = y

	return true
}

func checkXPosition(x int) bool {
	return (x >= 0 && x < MAX_X)
}

func checkYPosition(y int) bool {
	return (y >= 0 && y < MAX_Y)
}
