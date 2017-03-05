package toyrobot

import (
	"testing"
)

func TestCmdRun(t *testing.T) {

	p := &Position{4, 4, NORTH}
	robot := &Robot{Pos: p, IsPlaced: false}

	cmd, err := CreateCmd("PLACE 0,0,SOUTH")
	if err != nil || cmd == nil {
		t.Errorf("cmd :%v, err:%v", cmd, err)
	} else {
		err = cmd.Run(robot)
		if err != nil {
			t.Errorf("Failed to run cmd :%v", err)
		}
	}

	cmd, err = CreateCmd("MOVE")
	if err != nil || cmd == nil {
		t.Errorf("cmd :%v, err:%v", cmd, err)
	} else {
		err = cmd.Run(robot)

		if err == nil {
			t.Errorf("Should fail by moving to invalid position")
		}

		if !robot.IsPlaced || robot.Pos.X != 0 ||
			robot.Pos.Y != 0 || robot.Pos.Direction != SOUTH {
			t.Errorf("isPlaced: %v, pos :%v", robot.IsPlaced, robot.Pos)
		}
	}
	
	cmd, err = CreateCmd("LEFT")
	if err != nil || cmd == nil {
		t.Errorf("cmd :%v, err:%v", cmd, err)
	} else {
		err = cmd.Run(robot)

		if err != nil {
			t.Errorf("Failed to run cmd :%v", err)
		}

		if !robot.IsPlaced || robot.Pos.X != 0 ||
			robot.Pos.Y != 0 || robot.Pos.Direction != EAST {
			t.Errorf("isPlaced: %v, pos :%v", robot.IsPlaced, robot.Pos)
		}
	}
	
	cmd, err = CreateCmd("RIGHT")
	if err != nil || cmd == nil {
		t.Errorf("cmd :%v, err:%v", cmd, err)
	} else {
		err = cmd.Run(robot)

		if err != nil {
			t.Errorf("Failed to run cmd :%v", err)
		}

		if !robot.IsPlaced || robot.Pos.X != 0 ||
			robot.Pos.Y != 0 || robot.Pos.Direction != SOUTH {
			t.Errorf("isPlaced: %v, pos :%v", robot.IsPlaced, robot.Pos)
		}
	}
	
	cmd, err = CreateCmd("REPORT")
	if err != nil || cmd == nil {
		t.Errorf("cmd :%v, err:%v", cmd, err)
	} else {
		err = cmd.Run(robot)
		if err != nil {
			t.Errorf("Failed to run cmd :%v", err)
		}
	}
}
