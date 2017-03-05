package toyrobot

import (
    "testing"
)

func TestCheckDirection(t *testing.T) {
	
	dir, err := CheckDirection("SOUTH")
	if dir != int(SOUTH) || err != nil {
		t.Errorf("dir :%v, err: %v", dir, err)
	}
	
	dir, err = CheckDirection("NORTH")
	if dir != int(NORTH) || err != nil {
		t.Errorf("dir :%v, err: %v", dir, err)
	}
	
	dir, err = CheckDirection("EAST")
	if dir != int(EAST) || err != nil {
		t.Errorf("dir :%v, err: %v", dir, err)
	}
	
	dir, err = CheckDirection("WEST")
	if dir != int(WEST) || err != nil {
		t.Errorf("dir :%v, err: %v", dir, err)
	}
	
	// failure cases
	dir, err = CheckDirection("SOUTH2")
	if dir != -1 || err == nil {
		t.Errorf("dir :%v, err: %v", dir, err)
	}
	
	dir, err = CheckDirection("")
	if dir != -1 || err == nil {
		t.Errorf("dir :%v, err: %v", dir, err)
	}
	
	dir, err = CheckDirection("\ntest\ntest")
	if dir != -1 || err == nil {
		t.Errorf("dir :%v, err: %v", dir, err)
	}
}

