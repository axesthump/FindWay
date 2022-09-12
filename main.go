package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	r int
	c int
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	in.Scan()
	countOperation, _ := strconv.Atoi(in.Text())

	for i := 0; i < countOperation; i++ {
		in.Scan()
		rows, _ := strconv.Atoi(in.Text())
		in.Scan()
		columns, _ := strconv.Atoi(in.Text())
		columns++

		field := make([][]string, rows)

		for r := 0; r < rows; r++ {
			in.Scan()
			field[r] = strings.Split(in.Text(), "")
		}

		pos, _ := findStar(field)

		_, path := findStarStartPosition(make([]position, 0), pos, field, "", true)

		fmt.Println(path)
	}
}

func findStar(field [][]string) (position, string) {

	for r := 0; r < len(field); r++ {
		for c := 0; c < len(field[r]); c++ {
			if field[r][c] == "*" {
				return findStarStartPosition(make([]position, 0), position{r, c}, field, "", false)
			}
		}
	}

	return position{0, 0}, ""

}

func findStarStartPosition(oldPositions []position, newPos position, field [][]string, resultPath string, needBuildFinalPath bool) (position, string) {
	var neighborsCount int

	leftNeighbor := getLeftStart(newPos, field)
	rightNeighbor := getRightStart(newPos, field)
	upNeighbor := getUpStart(newPos, field)
	downNeighbor := getDownStart(newPos, field)

	if leftNeighbor != nil {
		neighborsCount += 1
	}
	if rightNeighbor != nil {
		neighborsCount += 1
	}
	if upNeighbor != nil {
		neighborsCount += 1
	}
	if downNeighbor != nil {
		neighborsCount += 1
	}

	if neighborsCount > 1 || needBuildFinalPath {
		if leftNeighbor != nil && !isBeenInThisPosition(leftNeighbor, oldPositions) {
			oldPositions = append(oldPositions, newPos)
			return findStarStartPosition(oldPositions, *leftNeighbor, field, resultPath+"L", needBuildFinalPath)
		}

		if upNeighbor != nil && !isBeenInThisPosition(upNeighbor, oldPositions) {
			oldPositions = append(oldPositions, newPos)
			return findStarStartPosition(oldPositions, *upNeighbor, field, resultPath+"U", needBuildFinalPath)
		}

		if rightNeighbor != nil && !isBeenInThisPosition(rightNeighbor, oldPositions) {
			oldPositions = append(oldPositions, newPos)
			return findStarStartPosition(oldPositions, *rightNeighbor, field, resultPath+"R", needBuildFinalPath)
		}

		if downNeighbor != nil && !isBeenInThisPosition(downNeighbor, oldPositions) {
			oldPositions = append(oldPositions, newPos)
			return findStarStartPosition(oldPositions, *downNeighbor, field, resultPath+"D", needBuildFinalPath)
		}

		return newPos, resultPath // В случае цикличной мапы
	} else {
		return newPos, resultPath
	}
}

func isBeenInThisPosition(neighbor *position, positions []position) bool {
	for _, pos := range positions {
		if pos.c == neighbor.c && pos.r == neighbor.r {
			return true
		}
	}
	return false
}

func getRightStart(pos position, field [][]string) *position {
	if pos.c+1 >= len(field[pos.r]) {
		return nil
	} else {
		if field[pos.r][pos.c+1] == "*" {
			return &position{pos.r, pos.c + 1}
		} else {
			return nil
		}
	}
}

func getLeftStart(pos position, field [][]string) *position {
	if pos.c == 0 {
		return nil
	} else {
		if field[pos.r][pos.c-1] == "*" {
			return &position{pos.r, pos.c - 1}
		} else {
			return nil
		}
	}
}

func getUpStart(pos position, field [][]string) *position {
	if pos.r == 0 {
		return nil
	} else {
		if field[pos.r-1][pos.c] == "*" {
			return &position{pos.r - 1, pos.c}
		} else {
			return nil
		}
	}
}

func getDownStart(pos position, field [][]string) *position {
	if pos.r+1 >= len(field) {
		return nil
	} else {
		if field[pos.r+1][pos.c] == "*" {
			return &position{pos.r + 1, pos.c}
		} else {
			return nil
		}
	}
}
