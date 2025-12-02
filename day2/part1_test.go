package day2

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"}

	got := Part1(input)
	want := 1227775554

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestValidateIDRange(t *testing.T) {
	tests := []struct {
		idRange  string
		subTotal int
	}{
		{
			idRange:  "11-22",
			subTotal: 33,
		},
		{
			idRange:  "95-115",
			subTotal: 99,
		},
		{
			idRange:  "1188511880-1188511890",
			subTotal: 1188511885,
		},
	}

	for _, tt := range tests {
		t.Run(tt.idRange, func(t *testing.T) {
			got := validateIDRange(tt.idRange)

			if got != tt.subTotal {
				t.Errorf("got %v, want %v", got, tt.subTotal)
			}
		})
	}
}

func TestIsValidID(t *testing.T) {
	tests := []struct {
		id   int
		want bool
	}{
		{
			id:   11,
			want: false,
		},
		{
			id:   446446,
			want: false,
		},
		{
			id:   38593862,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.id), func(t *testing.T) {
			got := isValidID(tt.id)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
