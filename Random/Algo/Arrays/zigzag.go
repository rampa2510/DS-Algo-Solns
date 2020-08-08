package main

func isOutOfBounds(row, col, height, width int) bool {
	return row < 0 || row > height || col < 0 || col > width
}

// if going down take care when row==height,col==0
// if row==height that means at the bottom so col++ else row++
// if going up take care of row==0,col==width
// if width==col row++ else col++
func zig(array [][]int) []int {
	if len(array) == 0 {
		return []int{}
	}
	row, col := 0, 0
	height, width := len(array)-1, len(array[0])-1
	res := make([]int, 0)
	goingDown := true
	for !isOutOfBounds(row, col, height, width) {
		res = append(res, array[row][col])
		if goingDown {
			if row == height || col == 0 {
				goingDown = false
				if row == height {
					col++
				} else {
					row++
				}
			} else {
				row++
				col--
			}
		} else {
			if row == 0 || col == width {
				goingDown = true
				if col == width {
					row++
				} else {
					col++
				}
			} else {
				row--
				col++
			}
		}
	}
	return res
}
