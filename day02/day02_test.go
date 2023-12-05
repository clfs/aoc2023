package day02

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGame_UnmarshalText(t *testing.T) {
	cases := []struct {
		in   string
		want Game
	}{
		{
			in: "Game 13: 2 blue, 5 green; 2 blue, 2 green; 2 blue, 2 red, 4 green",
			want: Game{
				ID: 13,
				Rounds: []Round{
					{Blue: 2, Green: 5},
					{Blue: 2, Green: 2},
					{Blue: 2, Red: 2, Green: 4},
				},
			},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			var got Game
			if err := got.UnmarshalText([]byte(tc.in)); err != nil {
				t.Error(err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("mismatch (-want, +got):\n%s", diff)
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

	want := 2256
	got := Part1(f)

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	f, err := os.Open("testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	want := 74229
	got := Part2(f)

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
