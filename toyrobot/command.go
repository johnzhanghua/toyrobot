package toyrobot

import (
	"errors"
)

type Command interface {
	Run(robot *Robot) error
	Name() string
}

type BasicCommand struct {
	Action string
}

type MoveCommand struct {
	BasicCommand
}

type LeftCommand struct {
	BasicCommand
}

type RightCommand struct {
	BasicCommand
}

type ReportCommand struct {
	BasicCommand
}

type PlaceCommand struct {
	BasicCommand
	Pos *Position
}

func (command *BasicCommand) Name() string {
	if command == nil {
		return "Nil"
	}

	return command.Action
}

func (command *MoveCommand) Run(robot *Robot) error {
	if robot == nil {
		return errors.New("Invalid robot object")
	}

	if robot.IsPlaced {
		return robot.Move()
	} else {
		return errors.New("Please run PLACE command first")
	}
}

func (command *LeftCommand) Run(robot *Robot) error {
	if robot == nil {
		return errors.New("Invalid robot object")
	}

	if robot.IsPlaced {
		return robot.Left()
	} else {
		return errors.New("Please run PLACE command first")
	}
}

func (command *RightCommand) Run(robot *Robot) error {
	if robot == nil {
		return errors.New("Invalid robot object")
	}

	if robot.IsPlaced {
		return robot.Right()
	} else {
		return errors.New("Please run PLACE command first")
	}
}

func (command *ReportCommand) Run(robot *Robot) error {
	if robot == nil {
		return errors.New("Invalid robot object")
	}

	if robot.IsPlaced {
		robot.Report()
		return nil
	} else {
		return errors.New("Please run PLACE command first")
	}
}

func (command *PlaceCommand) Run(robot *Robot) error {

	if robot == nil || command == nil || robot.Pos == nil || command.Pos == nil {
		return errors.New("Invalid null object")
	}

	robot.Pos.X = command.Pos.X
	robot.Pos.Y = command.Pos.Y
	robot.Pos.Direction = command.Pos.Direction

	robot.IsPlaced = true

	return nil
}
