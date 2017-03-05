package toyrobot

import (
	_ "fmt"
	"strings"
	"testing"
)

func TestCreatePosition(t *testing.T) {

	args := strings.Split("0,0,SOUTH", ",")
	pos, err := CreatePosition(args)
	if err != nil {
		t.Errorf("pos :%v, err: %v", pos, err)
	}

	args = strings.Split("0,4,EAST", ",")
	pos, err = CreatePosition(args)
	if err != nil {
		t.Errorf("pos :%v, err: %v", pos, err)
	}

	args = strings.Split("4,0,WEST", ",")
	pos, err = CreatePosition(args)
	if err != nil {
		t.Errorf("pos :%v, err: %v", pos, err)
	}

	args = strings.Split("4,4,NORTH", ",")
	pos, err = CreatePosition(args)
	if err != nil {
		t.Errorf("pos :%v, err: %v", pos, err)
	}

	// failure case
	args = strings.Split("5,5,SOUTH", ",")
	pos, err = CreatePosition(args)
	if err == nil {
		t.Errorf("pos :%v, err: %v", pos, err)
	}

	args = strings.Split("-1,-1,SOUTH", ",")
	pos, err = CreatePosition(args)
	if err == nil {
		t.Errorf("pos :%v, err: %v", pos, err)
	}

	args = strings.Split("0,5,SOUTH", ",")
	pos, err = CreatePosition(args)

	if err == nil {
		t.Errorf("pos :%v, err: %v", pos, err)
	}

	args = strings.Split("0, 4, DIR", ",")
	pos, err = CreatePosition(args)
	if err == nil {
		t.Errorf("pos :%v, err: %v", pos, err)
	}

	args = strings.Split("0,4,DIR", ",")
	pos, err = CreatePosition(args)
	if err == nil {
		t.Errorf("pos :%v, err: %v", pos, err)
	}
}
