package day03

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

type Point struct {
	R, C int
}

func (p Point) Neighbors() []Point {
	return []Point{
		{p.R - 1, p.C - 1},
		{p.R - 1, p.C},
		{p.R - 1, p.C + 1},
		{p.R, p.C - 1},
		{p.R, p.C + 1},
		{p.R + 1, p.C - 1},
		{p.R + 1, p.C},
		{p.R + 1, p.C + 1},
	}
}

func IsSymbol(r rune) bool {
	switch {
	case r == '\n':
		return false
	case r == '.':
		return false
	case '0' <= r && r <= '9':
		return false
	default:
		return true
	}
}

type Schematic struct {
	Symbols map[Point]struct{} // set of symbol locations
	Parts   map[Point]int      // map of leftmost part locations
}

var regexpNumber = regexp.MustCompile(`\d+`)

func NewSchematic(r io.Reader) (*Schematic, error) {
	schematic := &Schematic{
		Symbols: make(map[Point]struct{}),
		Parts:   make(map[Point]int),
	}

	var lineNum int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		// Read in symbols.
		for i, rn := range line {
			if IsSymbol(rn) {
				p := Point{R: lineNum, C: i}
				schematic.Symbols[p] = struct{}{}
			}
		}

		// Read in parts.
		locations := regexpNumber.FindAllStringIndex(line, -1)
		for _, loc := range locations {
			match := line[loc[0]:loc[1]]
			n, err := strconv.Atoi(match)
			if err != nil {
				return nil, fmt.Errorf("invalid number on line %d, column %d", lineNum, loc[0])
			}
			p := Point{R: lineNum, C: loc[0]}
			schematic.Parts[p] = n
		}

		// Prepare for next line.
		lineNum++
	}

	return schematic, scanner.Err()
}

func (s Schematic) PartNumbers() []int {
	var res []int

	for loc, val := range s.Parts {
		var check []Point

		for i := 0; i < val; i++ {
			// Check the neighbors of every point in the number,
			// with some redundancy.
			check = append(check, Point{loc.R, loc.C + i}.Neighbors()...)
		}

		for _, c := range check {
			_, ok := s.Symbols[c]
			if ok {
				// Save part values next to symbols.
				res = append(res, val)
				break
			}
		}
	}

	return res
}

func Part1(r io.Reader) (int, error) {
	s, err := NewSchematic(r)
	if err != nil {
		return 0, err
	}

	var acc int
	for _, n := range s.PartNumbers() {
		acc += n
	}
	return acc, nil
}
