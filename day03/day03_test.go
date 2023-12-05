package day03

import (
	"fmt"
	"maps"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewSchematic(t *testing.T) {
	cases := []struct {
		in   string
		want *Schematic
	}{
		{in: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
			want: &Schematic{
				Symbols: map[Point]struct{}{
					{1, 3}: {},
					{3, 6}: {},
					{4, 3}: {},
					{5, 5}: {},
					{8, 3}: {},
					{8, 5}: {},
				},
				Parts: map[Point]int{
					{0, 0}: 467,
					{0, 5}: 114,
					{2, 2}: 35,
					{2, 6}: 633,
					{4, 0}: 617,
					{5, 7}: 58,
					{6, 2}: 592,
					{7, 6}: 755,
					{9, 1}: 664,
					{9, 5}: 598,
				},
			}},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got, err := NewSchematic(strings.NewReader(tc.in))
			if err != nil {
				t.Error(err)
			}
			if !maps.Equal(tc.want.Symbols, got.Symbols) {
				t.Errorf("symbols: want %v, got %v", tc.want.Symbols, got.Symbols)
			}
			if !maps.Equal(tc.want.Parts, got.Parts) {
				t.Errorf("parts: want %v, got %v", tc.want.Parts, got.Parts)
			}
		})
	}
}

func TestSchematic_PartNumbers(t *testing.T) {
	cases := []struct {
		in   string
		want []int
	}{
		{in: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
			want: []int{
				35,
				467,
				592,
				598,
				617,
				633,
				664,
				755,
			},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			s, err := NewSchematic(strings.NewReader(tc.in))
			if err != nil {
				t.Fatal(err)
			}

			got := s.PartNumbers()

			if diff := cmp.Diff(tc.want, got, cmpopts.SortSlices(
				func(a, b int) bool { return a < b },
			)); diff != "" {
				t.Errorf("(-want, +got):\n%s", diff)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	f, err := os.Open("testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	want := 0
	got, err := Part1(f)
	if err != nil {
		t.Error(err)
	}
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
