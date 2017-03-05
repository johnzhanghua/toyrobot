package toyrobot

import (
    "testing"
)

func TestLeft(t *testing.T) {
	p := &Position{0,0,SOUTH}
	r := &Robot{Pos: p, IsPlaced: true}
	
	r.Left()
	if p.Direction != EAST {
		t.Errorf("Wrong value after left :%v", p.Direction)
	}
	r.Left() 
	if p.Direction != NORTH {
		t.Errorf("Wrong value after left :%v", p.Direction)
	}
	r.Left()
	if p.Direction != WEST {
		t.Errorf("Wrong value after left : %d", p.Direction)
	}
	r.Left() 
	if p.Direction != SOUTH {
		t.Errorf("Wrong value after left : %d", p.Direction)
	}
	
	r.Left()
	if p.Direction != EAST {
		t.Errorf("Wrong value after left :%v", p.Direction)
	}
	
}

func TestRight(t *testing.T) {
	p := &Position{0,0,SOUTH}
	r := &Robot{Pos: p, IsPlaced: true}
	
	r.Right()
	if p.Direction != WEST {
		t.Errorf("Wrong value after left :%v", p.Direction)
	}
	r.Right() 
	if p.Direction != NORTH {
		t.Errorf("Wrong value after left :%v", p.Direction)
	}
	r.Right()
	if p.Direction != EAST {
		t.Errorf("Wrong value after left : %d", p.Direction)
	}
	r.Right() 
	if p.Direction != SOUTH {
		t.Errorf("Wrong value after left : %d", p.Direction)
	}
	
	r.Right()
	if p.Direction != WEST {
		t.Errorf("Wrong value after left :%v", p.Direction)
	}
}


func TestMove(t *testing.T) {
	p := &Position{0,0,SOUTH}
	r := &Robot{Pos: p, IsPlaced: true}
	
	err := r.Move()
	if err == nil {
		t.Errorf("Should not move")
	}
	
	r.Left()
	
	err = r.Move()
	if err != nil {
		t.Errorf("%v", err)
	}
	
	if p.X != 1 || p.Y !=0 {
		t.Errorf("x value is %v, y value is %v", p.X, p.Y)
	}
	
	p = &Position{4,4,NORTH}
	r = &Robot{Pos: p, IsPlaced: true}
	
	err = r.Move()
	if err == nil {
		t.Errorf("Should not move")
	}
	
	r.Right()
	err = r.Move()
	if err == nil {
		t.Errorf("Should not move")
	}
	
	r.Right()
	err = r.Move()
	if err != nil {
		t.Errorf("%v", err)
	}
	
	if (p.X != 4 || p.Y != 3 || p.Direction != SOUTH) {
		t.Errorf("x: %v, y: %v, dir :%v", p.X, p.Y, p.Direction)
	}
}