package day1

import "testing"

func Test_Part2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name: "Password = 6",
			input: []string{
				"L68",
				"L30",
				"R48",
				"L5",
				"R60",
				"L55",
				"L1",
				"L99",
				"R14",
				"L82",
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part2(tt.input)

			if got != tt.want {
				t.Fatalf("Part1() = %v, wantDial %v", got, tt.want)
			}
		})
	}
}

func Test_RotateRightAndCountZeroes(t *testing.T) {
	tests := []struct {
		name       string
		amount     int
		dial       int
		wantDial   int
		wantZeroes int
	}{
		{
			name:       "95 to 55",
			amount:     30,
			dial:       50,
			wantDial:   80,
			wantZeroes: 0,
		},
		{
			name:       "52 to 10",
			amount:     58,
			dial:       52,
			wantDial:   10,
			wantZeroes: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDial, gotZeroes := rotateRightAndCountZeroes(tt.amount, tt.dial)

			if gotDial != tt.wantDial {
				t.Fatalf("RotateRight() = %v, wantDial %v", gotDial, tt.wantDial)
			}

			if gotZeroes != tt.wantZeroes {
				t.Fatalf("RotateRight() = %v, wantZeroes %v", gotZeroes, tt.wantZeroes)
			}
		})
	}
}

func Test_RotateLeftAndCountZeroes(t *testing.T) {
	tests := []struct {
		name       string
		amount     int
		dial       int
		wantDial   int
		wantZeroes int
	}{
		{
			name:       "50 to 82",
			amount:     68,
			dial:       50,
			wantDial:   82,
			wantZeroes: 1,
		},
		{
			name:       "82 to 52",
			amount:     30,
			dial:       82,
			wantDial:   52,
			wantZeroes: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDial, gotZeroes := rotateLeftAndCountZeroes(tt.amount, tt.dial)

			if gotDial != tt.wantDial {
				t.Fatalf("RotateLeft() = %v, wantDial %v", gotDial, tt.wantDial)
			}

			if gotZeroes != tt.wantZeroes {
				t.Fatalf("RotateLeft() = %v, wantZeroes %v", gotZeroes, tt.wantZeroes)
			}
		})
	}
}
