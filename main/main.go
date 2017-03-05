package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"log"
	"time"
	"toyrobot"
)

const Version string = "1.0"

func usage() {
	fmt.Print(
` 
This application accepts the following commands:
		
PLACE X,Y,F
MOVE
LEFT
RIGHT
REPORT

. The application is a simulation of a toy robot moving on a square tabletop, of dimensions 5 units x 5 units.
. PLACE will put the toy robot on the table in position X,Y and facing NORTH, SOUTH, EAST or WEST.
. The origin (0,0) can be considered to be the SOUTH WEST most corner.
. The first valid command to the robot is a PLACE command, after that, any sequence of commands may be issued, in any order, including another PLACE command. The application should discard all commands in the sequence until a valid PLACE command has been executed.
. MOVE will move the toy robot one unit forward in the direction it is currently facing.
. LEFT and RIGHT will rotate the robot 90 degrees in the specified direction without changing the position of the robot.
. REPORT will announce the X,Y and orientation of the robot.
`)
}

func main() {
	
	// set log
	t := time.Now()

	f, err := os.OpenFile("toyrobot_log_" + t.Format("2006-01-02"), os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
	    fmt.Printf("error opening logging file: %v\n", err)
	}
	defer f.Close()
	
	log.SetOutput(f)

	// init the robot pos
	p := &toyrobot.Position{0, 0, toyrobot.SOUTH}
	robot := &toyrobot.Robot{Pos: p, IsPlaced: false}

	// print usage
	fmt.Printf("Toy Robot Simulator %s\n", Version)

	usage()

	// read inputs
	reader := bufio.NewReader(os.Stdin)
	for {
		args, _ := reader.ReadString('\n')
		args = strings.TrimSpace(args)
		
		cmd, err := toyrobot.CreateCmd(args)
		if err != nil || cmd == nil {
			fmt.Printf("Error :%v \n", err)
		} else {
			err = cmd.Run(robot)
			if err != nil {
				fmt.Printf("Failed to run cmd %v,  error : %v \n", cmd.Name(), err)
			}
		}
	}
}
