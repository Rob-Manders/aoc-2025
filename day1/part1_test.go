package day1

import "testing"

func Test_Part1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "Password = 1",
			input: []string{"L50"},
			want:  1,
		},
		{
			name:  "Password = 2",
			input: []string{"L50", "R20", "L30", "R10"},
			want:  2,
		},
		{
			name: "Password = 3",
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
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part1(tt.input)

			if got != tt.want {
				t.Fatalf("Part1() = %v, wantDial %v", got, tt.want)
			}
		})
	}
}

func Test_RotateLeft(t *testing.T) {
	tests := []struct {
		name   string
		amount int
		dial   int
		want   int
	}{
		{
			name:   "50 to 82",
			amount: 68,
			dial:   50,
			want:   82,
		},
		{
			name:   "82 to 52",
			amount: 30,
			dial:   82,
			want:   52,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rotateLeft(tt.amount, tt.dial)

			if got != tt.want {
				t.Fatalf("RotateLeft() = %v, wantDial %v", got, tt.want)
			}
		})
	}
}

func Test_RotateRight(t *testing.T) {
	tests := []struct {
		name   string
		amount int
		dial   int
		want   int
	}{
		{
			name:   "95 to 55",
			amount: 30,
			dial:   50,
			want:   80,
		},
		{
			name:   "52 to 0",
			amount: 48,
			dial:   52,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rotateRight(tt.amount, tt.dial)

			if got != tt.want {
				t.Fatalf("RotateRight() = %v, wantDial %v", got, tt.want)
			}
		})
	}
}
