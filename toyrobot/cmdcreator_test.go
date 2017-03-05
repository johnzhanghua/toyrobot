package toyrobot

import (
    "testing"
)

func TestCreateCmd(t *testing.T) {
	cmd, err := CreateCmd("PLACE 0,0,SOUTH")
	if err != nil || cmd == nil {
		t.Errorf("cmd :%v, err:%v", cmd, err)
	} 

	cmd, err = CreateCmd("MOVE")
	if err != nil || cmd == nil {
		t.Errorf("cmd :%v, err:%v", cmd, err)
	} 
	
	cmd, err = CreateCmd("LEFT")
	if err != nil || cmd == nil {
		t.Errorf("cmd :%v, err:%v", cmd, err)
	} 
	
	cmd, err = CreateCmd("RIGHT")
	if err != nil || cmd == nil {
		t.Errorf("cmd :%v, err:%v", cmd, err)
	} 
	
	cmd, err = CreateCmd("REPORT")
	if err != nil || cmd == nil {
		t.Errorf("cmd :%v, err:%v", cmd, err)
	} 
	
	// Failure case
	cmd, err = CreateCmd("\nMove\ns")
	if err == nil || cmd != nil {
		t.Errorf("Should fail as Invalid command")
	}
	
	cmd, err = CreateCmd("")
	if err == nil || cmd != nil {
		t.Errorf("Should fail as Invalid command")
	}
	
	cmd, err = CreateCmd("PLACE 5,5,SOUTH")
	if err == nil || cmd != nil {
		t.Errorf("cmd :%v, err:%v", cmd, err)
	} 
	
	cmd, err = CreateCmd("PLACE -1,-1,SOUTH")
	if err == nil || cmd != nil {
		t.Errorf("cmd :%v, err:%v", cmd, err)
	} 
	
	cmd, err = CreateCmd("PLACE 0,0,TEST")
	if err == nil || cmd != nil {
		t.Errorf("cmd :%v, err:%v", cmd, err)
	} 
}

