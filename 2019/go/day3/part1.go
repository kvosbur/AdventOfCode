package day3

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type horizontalLine struct {
	startX    int
	endX      int
	y         int
	stepCount int
	direction string
}

type verticalLine struct {
	startY    int
	endY      int
	x         int
	stepCount int
	direction string
}

func segmentsIntersect(hl horizontalLine, vl verticalLine) bool {
	return hl.startX <= vl.x && hl.endX >= vl.x &&
		vl.startY <= hl.y && vl.endY >= hl.y
}

func getManDistance(hl horizontalLine, vl verticalLine) int {
	return int(math.Abs(float64(hl.y)) + math.Abs(float64(vl.x)))
}

func getSegments(instructions []string) ([]horizontalLine, []verticalLine) {
	horizontals := []horizontalLine{}
	verticals := []verticalLine{}
	currentX := 0
	currentY := 0
	for _, instruction := range instructions {
		direction := string(instruction[0])
		amount, _ := strconv.Atoi(instruction[1:])
		switch direction {
		case "U":
			line := verticalLine{
				x:      currentX,
				startY: currentY,
				endY:   currentY + amount,
			}
			currentY += amount
			verticals = append(verticals, line)
		case "D":
			line := verticalLine{
				x:      currentX,
				startY: currentY - amount,
				endY:   currentY,
			}
			currentY -= amount
			verticals = append(verticals, line)
		case "R":
			line := horizontalLine{
				y:      currentY,
				startX: currentX,
				endX:   currentX + amount,
			}
			currentX += amount
			horizontals = append(horizontals, line)

		case "L":
			line := horizontalLine{
				y:      currentY,
				startX: currentX - amount,
				endX:   currentX,
			}
			currentX -= amount
			horizontals = append(horizontals, line)

		default:
			fmt.Println("Unexpected direction", direction)
			os.Exit(1)
		}

	}
	return horizontals, verticals
}

func Part1Solution(input []string) string {
	firstLine := strings.Split(input[0], ",")
	secondLine := strings.Split(input[1], ",")

	firstHorizontals, firstVerticals := getSegments(firstLine)
	secondHorizontals, secondVerticals := getSegments(secondLine)

	// fmt.Printf("1 horizontals %+v\n", firstHorizontals)
	// fmt.Printf("1 verticals %+v\n", firstVerticals)

	// fmt.Printf("2 horizontals %+v\n", secondHorizontals)
	// fmt.Printf("2 verticals %+v\n", secondVerticals)

	best := int(^uint(0) >> 1)
	for _, hl := range secondHorizontals {
		for _, vl := range firstVerticals {
			if segmentsIntersect(hl, vl) {
				d := getManDistance(hl, vl)
				if d != 0 && d < best {
					best = d
				}
				// fmt.Println("hl:", hl, "vl:", vl, "Intersection:", d)
			}
		}
	}

	for _, vl := range secondVerticals {
		for _, hl := range firstHorizontals {
			if segmentsIntersect(hl, vl) {
				d := getManDistance(hl, vl)
				if d != 0 && d < best {
					best = d
				}
				// fmt.Println("hl:", hl, "vl:", vl, "Intersection:", d)
			}
		}
	}

	return strconv.Itoa(best)
}

// 1064
