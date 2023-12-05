package day01

import (
	"os"
	"testing"
)

func Test_calibrationValue(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}

	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			got := calibrationValue(tc.in)
			if tc.want != got {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}

// 29, 83, 13, 24, 42, 14, and 76.

func Test_calibrationValue2(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	}

	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			got := calibrationValue2(tc.in)
			if tc.want != got {
				t.Errorf("want %d, got %d", tc.want, got)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	f, err := os.Open("testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	want := 55447
	got := part1(f)

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func Test_part2(t *testing.T) {
	f, err := os.Open("testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	want := 54706
	got := part2(f)

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
