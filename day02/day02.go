package day02

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

type Round struct {
	Red, Green, Blue int
}

var (
	regexpRed   = regexp.MustCompile(`(\d+) red`)
	regexpGreen = regexp.MustCompile(`(\d+) green`)
	regexpBlue  = regexp.MustCompile(`(\d+) blue`)
)

func (r *Round) UnmarshalText(text []byte) error {
	var err error

	redMatches := regexpRed.FindSubmatch(text)
	if redMatches == nil {
		r.Red = 0
	} else {
		r.Red, err = strconv.Atoi(string(redMatches[1]))
		if err != nil {
			return err
		}
	}

	greenMatches := regexpGreen.FindSubmatch(text)
	if greenMatches == nil {
		r.Green = 0
	} else {
		r.Green, err = strconv.Atoi(string(greenMatches[1]))
		if err != nil {
			return err
		}
	}

	blueMatches := regexpBlue.FindSubmatch(text)
	if blueMatches == nil {
		r.Blue = 0
	} else {
		r.Blue, err = strconv.Atoi(string(blueMatches[1]))
		if err != nil {
			return err
		}
	}

	return nil
}

type Bag struct {
	Red, Green, Blue int
}

func (b Bag) Power() int {
	return b.Red * b.Green * b.Blue
}

func (b Bag) IsPossible(g Game) bool {
	for _, r := range g.Rounds {
		if r.Red > b.Red || r.Green > b.Green || r.Blue > b.Blue {
			return false
		}
	}
	return true
}

type Game struct {
	ID     int
	Rounds []Round
}

func (g *Game) SmallestBag() Bag {
	var bag Bag

	for _, r := range g.Rounds {
		bag.Red = max(bag.Red, r.Red)
		bag.Green = max(bag.Green, r.Green)
		bag.Blue = max(bag.Blue, r.Blue)
	}

	return bag
}

var regexpGameID = regexp.MustCompile(`^Game (\d+)`)

func (g *Game) UnmarshalText(text []byte) error {
	matches := regexpGameID.FindSubmatch(text)
	if matches == nil {
		return errors.New("no match for ID regexp")
	}

	numID, err := strconv.Atoi(string(matches[1]))
	if err != nil {
		return fmt.Errorf("matches for ID regexp invalid: %q", matches)
	}

	g.ID = numID

	_, tail, ok := bytes.Cut(text, []byte(": "))
	if !ok {
		return errors.New("bad bytes.Cut")
	}

	var rounds []Round

	for _, b := range bytes.Split(tail, []byte(";")) {
		var r Round
		if err := r.UnmarshalText(b); err != nil {
			return err
		}
		rounds = append(rounds, r)
	}

	g.Rounds = rounds

	return nil
}

func Part1(r io.Reader) int {
	bag := Bag{Red: 12, Green: 13, Blue: 14}

	var acc int

	var g Game

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Bytes()
		if len(line) == 0 {
			continue
		}
		if err := g.UnmarshalText(line); err != nil {
			panic(err)
		}
		if bag.IsPossible(g) {
			acc += g.ID
		}
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	return acc
}

func Part2(r io.Reader) int {
	var acc int

	var g Game

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Bytes()
		if len(line) == 0 {
			continue
		}
		if err := g.UnmarshalText(line); err != nil {
			panic(err)
		}
		acc += g.SmallestBag().Power()
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	return acc
}
