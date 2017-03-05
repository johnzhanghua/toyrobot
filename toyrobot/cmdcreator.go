package toyrobot

import (
	"errors"
	"strings"
)

type CreateCmdFunc func(args []string, creator *commandCreator) (Command, error)

type commandCreator struct {
	action  string
	creator CreateCmdFunc
	argsLen int
}

// the command list , add more to extend it
var commandCreatorList = []commandCreator{
	commandCreator{
		action:  "PLACE",
		creator: createPlaceCmd,
		argsLen: 2,
	},
	commandCreator{
		action:  "MOVE",
		creator: createMoveCmd,
		argsLen: 1,
	},
	commandCreator{
		action:  "LEFT",
		creator: createLeftCmd,
		argsLen: 1,
	},
	commandCreator{
		action:  "RIGHT",
		creator: createRightCmd,
		argsLen: 1,
	},
	commandCreator{
		action:  "REPORT",
		creator: createReportCmd,
		argsLen: 1,
	},
}

func CreateCmd(arg string) (Command, error) {

	args := strings.Split(arg, " ")

	if len(args) < 1 {
		return nil, errors.New("Invalid Command")
	}

	for _, cmdCreator := range commandCreatorList {
		if args[0] == cmdCreator.action {
			return cmdCreator.creator(args, &cmdCreator)
		}
	}

	return nil, errors.New("Invalid Command")
}

func createMoveCmd(args []string, creator *commandCreator) (Command, error) {
	if creator == nil {
		return nil, errors.New("Invalid command creator object")
	}

	if len(args) == creator.argsLen {
		return &MoveCommand{BasicCommand{Action: args[0]}}, nil
	} else {
		return nil, errors.New("Invalid Command")
	}
}

func createLeftCmd(args []string, creator *commandCreator) (Command, error) {
	if creator == nil {
		return nil, errors.New("Invalid command creator object")
	}

	if len(args) == creator.argsLen {
		return &LeftCommand{BasicCommand{Action: args[0]}}, nil
	} else {
		return nil, errors.New("Invalid Command")
	}
}

func createRightCmd(args []string, creator *commandCreator) (Command, error) {
	if creator == nil {
		return nil, errors.New("Invalid command creator object")
	}

	if len(args) == creator.argsLen {
		return &RightCommand{BasicCommand{Action: args[0]}}, nil
	} else {
		return nil, errors.New("Invalid Command")
	}
}

func createReportCmd(args []string, creator *commandCreator) (Command, error) {
	if creator == nil {
		return nil, errors.New("Invalid command creator object")
	}

	if len(args) == creator.argsLen {
		return &ReportCommand{BasicCommand{Action: args[0]}}, nil
	} else {
		return nil, errors.New("Invalid Command")
	}
}

func createPlaceCmd(args []string, creator *commandCreator) (Command, error) {
	if creator == nil {
		return nil, errors.New("Invalid command creator object")
	}

	if len(args) != creator.argsLen {
		return nil, errors.New("Invalid Command")
	}

	posArgs := strings.Split(args[1], ",")

	if pos, erro := CreatePosition(posArgs); erro == nil {
		return &PlaceCommand{BasicCommand{Action: args[0]}, pos}, nil
	} else {
		return nil, erro
	}
}
