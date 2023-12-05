package day01

import (
	"bufio"
	"io"
	"strings"
)

func calibrationValue(s string) int {
	var nums []int
	for _, rn := range s {
		if '0' <= rn && rn <= '9' {
			nums = append(nums, int(rn-'0'))
		}
	}
	if len(nums) == 0 {
		panic(nums)
	}
	first, last := nums[0], nums[len(nums)-1]
	return first*10 + last
}

// input has no zeroes
var scoreMap = map[string]int{
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func calibrationValue2(s string) int {
	left, right := len(s)-1, 0
	leftVal, rightVal := 0, 0

	for ss, score := range scoreMap {
		i := strings.Index(s, ss)
		if i != -1 && i <= left {
			left, leftVal = i, score
		}
		j := strings.LastIndex(s, ss)
		if j != -1 && j >= right {
			right, rightVal = j, score
		}
	}

	return leftVal*10 + rightVal
}

func part1(r io.Reader) int {
	var acc int

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}
		acc += calibrationValue(line)
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	return acc
}

func part2(r io.Reader) int {
	var acc int

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}
		acc += calibrationValue2(line)
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	return acc
}
