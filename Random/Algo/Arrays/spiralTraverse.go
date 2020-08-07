package main

func spiralTraverse(array [][]int) []int {
	arr := make([]int, 0)
	startRow, endRow := 0, len(array)-1
	startCol, endCol := 0, len(array[0])-1
	// fmt.Println(startCol, startCol, endCol, endRow)

	for startRow <= endRow && startCol <= endCol {
		// fmt.Println(startRow, endRow, startCol, endCol)

		for col := startCol; col <= endCol; col++ {
			arr = append(arr, array[startRow][col])
		}

		for row := startRow + 1; row <= endRow; row++ {
			arr = append(arr, array[row][endCol])
		}

		for col := endCol - 1; col >= startCol; col-- {
			if startRow == endRow {
				break
			}
			// fmt.Println(startRow, endRow, startCol, endCol)

			arr = append(arr, array[endRow][col])
		}

		for row := endRow - 1; row > startRow; row-- {
			if endCol == startCol {
				break
			}
			arr = append(arr, array[row][startCol])
		}
		startRow++
		endRow--
		startCol++
		endCol--
		// fmt.Println(startRow, endRow, startCol, endCol)
	}
	return arr
}
