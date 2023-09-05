package day3

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func absoluteInt(i int) int {
	return int(math.Abs(float64(i)))
}

func (hl horizontalLine) getIntersectLength(x int) int {
	if hl.direction == "R" {
		return absoluteInt(x - hl.startX)
	}
	return absoluteInt(hl.endX - x)
}

func (vl verticalLine) getIntersectLength(y int) int {
	if vl.direction == "U" {
		return absoluteInt(y - vl.startY)
	}
	return absoluteInt(vl.endY - y)
}

func getSegmentsWithStep(instructions []string) ([]horizontalLine, []verticalLine) {
	horizontals := []horizontalLine{}
	verticals := []verticalLine{}
	currentX := 0
	currentY := 0
	steps := 0
	for _, instruction := range instructions {
		direction := string(instruction[0])
		amount, _ := strconv.Atoi(instruction[1:])
		switch direction {
		case "U":
			line := verticalLine{
				x:         currentX,
				startY:    currentY,
				endY:      currentY + amount,
				stepCount: steps,
				direction: direction,
			}
			currentY += amount
			verticals = append(verticals, line)
		case "D":
			line := verticalLine{
				x:         currentX,
				startY:    currentY - amount,
				endY:      currentY,
				stepCount: steps,
				direction: direction,
			}
			currentY -= amount
			verticals = append(verticals, line)
		case "R":
			line := horizontalLine{
				y:         currentY,
				startX:    currentX,
				endX:      currentX + amount,
				stepCount: steps,
				direction: direction,
			}
			currentX += amount
			horizontals = append(horizontals, line)

		case "L":
			line := horizontalLine{
				y:         currentY,
				startX:    currentX - amount,
				endX:      currentX,
				stepCount: steps,
				direction: direction,
			}
			currentX -= amount
			horizontals = append(horizontals, line)

		default:
			fmt.Println("Unexpected direction", direction)
			os.Exit(1)
		}

		steps += amount
	}
	return horizontals, verticals
}

func getStepDistance(hl horizontalLine, vl verticalLine) int {
	return hl.stepCount + vl.stepCount + hl.getIntersectLength(vl.x) + vl.getIntersectLength(hl.y)
}

func Part2Solution(input []string) string {
	firstLine := strings.Split(input[0], ",")
	secondLine := strings.Split(input[1], ",")

	firstHorizontals, firstVerticals := getSegmentsWithStep(firstLine)
	secondHorizontals, secondVerticals := getSegmentsWithStep(secondLine)

	fmt.Printf("1 horizontals %+v\n", firstHorizontals)
	fmt.Printf("1 verticals %+v\n", firstVerticals)

	fmt.Printf("2 horizontals %+v\n", secondHorizontals)
	fmt.Printf("2 verticals %+v\n", secondVerticals)

	best := int(^uint(0) >> 1)
	for _, hl := range secondHorizontals {
		for _, vl := range firstVerticals {
			if segmentsIntersect(hl, vl) {
				d := getStepDistance(hl, vl)
				if d != 0 && d < best {
					best = d
				}
				// fmt.Println("hl:", hl, "vl:", vl, "Intersection:", d, hl.getIntersectLength(vl.x), vl.getIntersectLength(hl.y))
				fmt.Printf("hl: %+v vl: %+v, intersect: %v  hlDist:%v, vlDist:%v\n", hl, vl, d, hl.getIntersectLength(vl.x), vl.getIntersectLength(hl.y))
			}
		}
	}

	for _, vl := range secondVerticals {
		for _, hl := range firstHorizontals {
			if segmentsIntersect(hl, vl) {
				d := getStepDistance(hl, vl)
				if d != 0 && d < best {
					best = d
				}
				// fmt.Println("hl:", hl, "vl:", vl, "Intersection:", d, hl.getIntersectLength(vl.x), vl.getIntersectLength(hl.y))
				fmt.Printf("hl: %+v vl: %+v, intersect: %v  hlDist:%v, vlDist:%v\n", hl, vl, d, hl.getIntersectLength(vl.x), vl.getIntersectLength(hl.y))
			}
		}
	}

	return strconv.Itoa(best)
}

// 25676 correct
