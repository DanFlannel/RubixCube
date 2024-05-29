package rubixcube

import "strings"

type RubixFaceImpl interface {
	RotateFaceCW()
	RotateFaceCCW()
	Print()
}

type RubixFace struct {
	RubixFaceImpl
	colors [][]string
	top    *RubixFace
	left   *RubixFace
	right  *RubixFace
	back   *RubixFace
	bottom *RubixFace
}

func CreateFace(prefix string) *RubixFace {
	res := &RubixFace{
		colors: [][]string{
			{prefix + "1", prefix + "2", prefix + "3"},
			{prefix + "4", prefix + "5", prefix + "6"},
			{prefix + "7", prefix + "8", prefix + "9"},
		},
	}
	return res
}

// associate individual faces with each other
func (r *RubixFace) AddEdges(top *RubixFace, left *RubixFace, right *RubixFace, back *RubixFace, bottom *RubixFace) {
	r.top = top
	r.left = left
	r.right = right
	r.back = back
	r.bottom = bottom
}

func (r *RubixFace) RotateFaceCW() {
	// rotate the face clockwise
	r.rotateCW()

	// get the edge values of the faces that will be rotated
	var values []string

	topValues := r.top.GetBottomValues()
	leftValues := r.left.GetRightValues()
	rightValues := r.right.GetLeftValues()
	bottomValues := r.bottom.GetTopValues()

	values = append(values, topValues...)
	values = append(values, leftValues...)
	values = append(values, rightValues...)
	values = append(values, bottomValues...)

	// set the edge values of the faces that will be rotated
	right := values[0:3]
	top := values[3:6]
	bottom := values[6:9]
	left := values[9:12]

	// apply the new edge values
	r.top.SetBottomValues(top)
	r.left.SetRightValues(left)
	r.right.SetLeftValues(right)
	r.bottom.SetTopValues(bottom)
}

func (r *RubixFace) RotateFaceCCW() {
	// rotate the face counter clockwise
	r.rotateCCW()

	// get the edge values of the faces that will be rotated
	var values []string

	topValues := r.top.GetBottomValues()
	leftValues := r.left.GetRightValues()
	rightValues := r.right.GetLeftValues()
	bottomValues := r.bottom.GetTopValues()

	values = append(values, topValues...)
	values = append(values, leftValues...)
	values = append(values, rightValues...)
	values = append(values, bottomValues...)

	// set the edge values of the faces that will be rotated
	right := values[9:12]
	top := values[6:9]
	bottom := values[3:6]
	left := values[0:3]

	// apply the new edge values
	r.top.SetBottomValues(top)
	r.left.SetRightValues(left)
	r.right.SetLeftValues(right)
	r.bottom.SetTopValues(bottom)
	// rotate the edges
}

func (r *RubixFace) GetLeftValues() []string {
	return []string{r.colors[0][0], r.colors[1][0], r.colors[2][0]}
}

func (r *RubixFace) GetRightValues() []string {
	return []string{r.colors[0][2], r.colors[1][2], r.colors[2][2]}
}

func (r *RubixFace) GetTopValues() []string {
	return []string{r.colors[0][0], r.colors[0][1], r.colors[0][2]}
}

func (r *RubixFace) GetBottomValues() []string {
	return []string{r.colors[2][0], r.colors[2][1], r.colors[2][2]}
}

func (r *RubixFace) SetLeftValues(values []string) {
	for i, val := range values {
		r.colors[i][0] = val
	}
}

func (r *RubixFace) SetRightValues(values []string) {
	for i, val := range values {
		r.colors[i][2] = val
	}
}

func (r *RubixFace) SetTopValues(values []string) {
	for i, val := range values {
		r.colors[0][i] = val
	}
}

func (r *RubixFace) SetBottomValues(values []string) {
	for i, val := range values {
		r.colors[2][i] = val
	}
}

// rotates our matrix by 90 degrees clockwise
func (r *RubixFace) rotateCW() {
	for i := 0; i < len(r.colors)/2; i++ {
		// helpers to make sense of the below a bit easier
		top := i
		bottom := len(r.colors) - 1 - i
		for j := i; j < len(r.colors)-i-1; j++ {
			temp := r.colors[top][j]
			r.colors[top][j] = r.colors[len(r.colors)-1-j][top]
			r.colors[len(r.colors)-1-j][top] = r.colors[bottom][len(r.colors)-1-j]
			r.colors[bottom][len(r.colors)-1-j] = r.colors[j][bottom]
			r.colors[j][bottom] = temp
		}
	}
}

// rotates our matrix by 90 degrees counter clockwise
func (r *RubixFace) rotateCCW() {
	for i := 0; i < len(r.colors)/2; i++ {
		// helpers to make sense of the below a bit easier
		top := i
		bottom := len(r.colors) - 1 - i
		for j := i; j < len(r.colors)-i-1; j++ {
			temp := r.colors[top][j]
			r.colors[top][j] = r.colors[j][bottom]
			r.colors[j][bottom] = r.colors[bottom][bottom-(j-top)]
			r.colors[bottom][bottom-(j-top)] = r.colors[bottom-(j-top)][top]
			r.colors[bottom-(j-top)][top] = temp
		}
	}
}

// Print prints the face
func (r *RubixFace) Print(offset string) {
	// print the face
	strs := r.PrintString()
	for _, str := range strs {
		print(offset + str)
		println()
	}
}

// PrintString returns a string representation of the face
func (r *RubixFace) PrintString() []string {
	var strs []string
	for _, row := range r.colors {
		str := leftEdge
		for _, color := range row {
			str += color + "  "
		}
		str = strings.Trim(str, " ")
		str += rightEdge
		strs = append(strs, str)
	}
	return strs
}
