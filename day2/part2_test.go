package day2

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestPart2(t *testing.T) {
	input := []string{"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"}

	got := Part2(input)
	want := 4174379265

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestValidateIDRangeWithRepeatingPatterns(t *testing.T) {
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
			subTotal: 210,
		},
		{
			idRange:  "1188511880-1188511890",
			subTotal: 1188511885,
		},
		{
			idRange:  "1698522-1698528",
			subTotal: 0,
		},
		{
			idRange:  "565653-565659",
			subTotal: 565656,
		},
	}

	for _, tt := range tests {
		t.Run(tt.idRange, func(t *testing.T) {
			got := validateIDRangeWithRepeatingPatterns(tt.idRange)

			if got != tt.subTotal {
				t.Errorf("got %v, want %v", got, tt.subTotal)
			}
		})
	}
}

func TestIsValidIDWithRepeatingPatterns(t *testing.T) {
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
		{
			id:   123123123,
			want: false,
		},
		{
			id:   1212121212,
			want: false,
		},
		{
			id:   1111111,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.id), func(t *testing.T) {
			got := isValidIDWithRepeatingPatterns(tt.id)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasRepeatingPattern(t *testing.T) {
	tests := []struct {
		pattern []string
		want    bool
	}{
		{
			pattern: []string{"12", "12"},
			want:    true,
		},
		{
			pattern: []string{"123", "123", "123"},
			want:    true,
		},
		{
			pattern: []string{"12", "12", "12", "12", "12"},
			want:    true,
		},
		{
			pattern: []string{"1234", "5678"},
			want:    false,
		},
		{
			pattern: []string{"12", "34", "56", "78"},
			want:    false,
		},
		{
			pattern: []string{"123", "456", "789"},
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.pattern, ","), func(t *testing.T) {
			got := hasRepeatingPattern(tt.pattern)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitID(t *testing.T) {
	tests := []struct {
		id      string
		divisor int
		want    []string
	}{
		{
			id:      "1234",
			divisor: 4,
			want:    []string{"1", "2", "3", "4"},
		},
		{
			id:      "1234",
			divisor: 2,
			want:    []string{"12", "34"},
		},
		{
			id:      "123456",
			divisor: 3,
			want:    []string{"12", "34", "56"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			got := splitID(tt.id, tt.divisor)

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDivisors(t *testing.T) {
	tests := []struct {
		id   string
		want []int
	}{
		{
			id:   "12",
			want: []int{2},
		},
		{
			id:   "123",
			want: []int{3},
		},
		{
			id:   "1234",
			want: []int{2, 4},
		},
		{
			id:   "123456",
			want: []int{2, 3, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			got := getDivisors(tt.id)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
