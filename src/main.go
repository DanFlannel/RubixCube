package main

import (
	"fmt"
	rubixcube "rubix_cube/src/rubix_cube"
	"sync"
	"time"

	"github.com/manifoldco/promptui"
)

const TurnLeft = "Turn Left"
const TurnRight = "Turn Right"
const TurnUp = "Turn Up"
const TurnDown = "Turn Down"
const RotateClockwise = "Rotate Face Clockwise"
const RotateCounterClockwise = "Rotate Face Counter Clockwise"
const Exit = "Exit"

func Process(stopIt *bool, wg *sync.WaitGroup, r rubixcube.RubixCube) {
	defer wg.Done()

	for {
		if *stopIt {
			fmt.Println("Thanks for playing :)")
			return
		}

		prompt := promptui.Select{
			Label: "Select Action",
			Items: []string{RotateClockwise, RotateCounterClockwise, TurnLeft, TurnRight, TurnUp, TurnDown, Exit},
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {
		case TurnLeft:
			r.RotateLeft()
			r.Print()
		case TurnRight:
			r.RotateRight()
			r.Print()
		case TurnUp:
			r.RotateUp()
			r.Print()
		case TurnDown:
			r.RotateDown()
			r.Print()
		case RotateClockwise:
			r.RotateCW()
			r.Print()
		case RotateCounterClockwise:
			r.RotateCCW()
			r.Print()
		case "Exit":
			*stopIt = true
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	print("Welcome to Rubix Cube\n")
	print("You can rotate the cube (left, right, up down) to change the front face or rotate the face of the cube clockwise or counter clockwise\n")
	print("Select your action from the list\n")
	r := rubixcube.New()
	r.Print()

	var stopIt bool
	var wg sync.WaitGroup

	wg.Add(1)
	go Process(&stopIt, &wg, *r)
	wg.Wait()
}
