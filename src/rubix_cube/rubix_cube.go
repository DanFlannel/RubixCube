package rubixcube

import (
	"fmt"
	"strings"
)

type RubixCubeImpl interface {
	RotateCCW()
	RotateCW()
	Print()
}

type RubixCube struct {
	faces       []RubixFace
	currentFace *RubixFace
}

const offset = "\t     "
const header = "+------------+"
const leftEdge = "| "
const rightEdge = " |"
const middleHeader = "+------------+------------+------------+------------+\n"

func New() *RubixCube {
	res := &RubixCube{
		faces: []RubixFace{
			*CreateFace("U"),
			*CreateFace("L"),
			*CreateFace("F"),
			*CreateFace("R"),
			*CreateFace("B"),
			*CreateFace("D"),
		},
	}

	top := &res.faces[0]
	left := &res.faces[1]
	front := &res.faces[2]
	right := &res.faces[3]
	back := &res.faces[4]
	down := &res.faces[5]

	top.AddEdges(back, left, right, down, front)
	left.AddEdges(top, back, front, right, down)
	front.AddEdges(top, left, right, back, down)
	right.AddEdges(top, front, back, left, down)
	back.AddEdges(top, right, left, front, down)
	down.AddEdges(front, left, right, top, back)

	res.currentFace = front
	return res
}

// these are our cli commands
func (r *RubixCube) RotateCCW() {
	r.currentFace.RotateFaceCCW()
}

func (r *RubixCube) RotateCW() {
	r.currentFace.RotateFaceCW()
}

func (r *RubixCube) RotateLeft() {
	r.currentFace.top.rotateCW()
	r.currentFace.bottom.rotateCW()
	r.currentFace = r.currentFace.right
}

func (r *RubixCube) RotateRight() {
	r.currentFace.top.rotateCCW()
	r.currentFace.bottom.rotateCCW()
	r.currentFace = r.currentFace.left
}

func (r *RubixCube) RotateUp() {
	r.currentFace = r.currentFace.bottom
}

func (r *RubixCube) RotateDown() {
	r.currentFace = r.currentFace.top
}

func (r *RubixCube) Print() {
	// print the cube
	r.printTopRow()
	r.printMiddleRow()
	r.printBottomRow()
}

func (r *RubixCube) printTopRow() {
	top := r.currentFace.top
	fmt.Println(offset + header)
	top.Print(offset)
}

func (r *RubixCube) printMiddleRow() {
	// get the correct faces
	var middleRow []RubixFace
	middleRow = append(middleRow, *r.currentFace.left, *r.currentFace, *r.currentFace.right, *r.currentFace.back)

	// print the top header
	print(middleHeader)

	var faceStrs [][]string
	// for each face in the middle row get the string representation
	for _, face := range middleRow {
		faceStrs = append(faceStrs, face.PrintString())
	}

	// for each row in the face
	for i := 0; i < 3; i++ {
		str := ""
		// for each face in the middle row concatenate the strings
		for _, faceStr := range faceStrs {
			str += faceStr[i]
		}
		// replace the double pipes with single pipes
		str = strings.Replace(str, "||", "|", -1)

		// print the row
		print(str)
		println()
	}

	// print the bottom header
	print(middleHeader)
}

func (r *RubixCube) printBottomRow() {
	bottom := r.currentFace.bottom
	bottom.Print(offset)
	fmt.Println(offset + header)
}
