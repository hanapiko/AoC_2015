package main

import (
	"testing"
)

func TestInstruction(t *testing.T) {
	type args struct {
		grid        [][]int
		instruction string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantGrid [][]int
	}{
		{
			name: "Turn on lights",
			args: args{
				grid:        initializeGrid(size),
				instruction: "turn on 0,0 through 2,2",
			},
			wantErr: false,
			wantGrid: func() [][]int {
				g := initializeGrid(size)
				for x := 0; x <= 2; x++ {
					for y := 0; y <= 2; y++ {
						g[x][y] = 1
					}
				}
				return g
			}(),
		},
		{
			name: "Turn off lights",
			args: args{
				grid:        initializeGrid(size),
				instruction: "turn on 0,0 through 2,2",
			},
			wantErr: false,
			wantGrid: func() [][]int {
				g := initializeGrid(size)
				for x := 0; x <= 2; x++ {
					for y := 0; y <= 2; y++ {
						g[x][y] = 1
					}
				}
				return g
			}(),
		},
		{
			name: "Toggle lights",
			args: args{
				grid:        initializeGrid(size),
				instruction: "toggle 0,0 through 1,1",
			},
			wantErr: false,
			wantGrid: func() [][]int {
				g := initializeGrid(size)
				for x := 0; x <= 1; x++ {
					for y := 0; y <= 1; y++ {
						g[x][y] = 2
					}
				}
				return g
			}(),
		},
		{
			name: "Invalid instruction format",
			args: args{
				grid:        initializeGrid(size),
				instruction: "turn on 0,0",
			},
			wantErr: true,
			wantGrid: initializeGrid(size),
		},
		{
			name: "Invalid action",
			args: args{
				grid:        initializeGrid(size),
				instruction: "unknown 0,0 through 1,1",
			},
			wantErr: true,
			wantGrid: initializeGrid(size),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset grid before each test
			grid := initializeGrid(size)
			copyGrid(tt.args.grid, grid)

			if err := Instruction(grid, tt.args.instruction); (err != nil) != tt.wantErr {
				t.Errorf("Instruction() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !equalGrid(grid, tt.wantGrid) {
				t.Errorf("Instruction() = %v, want %v", grid, tt.wantGrid)
			}
		})
	}
}

// Helper function to initialize a grid
func initializeGrid(size int) [][]int {
	grid := make([][]int, size)
	for i := range grid {
		grid[i] = make([]int, size)
	}
	return grid
}

// Helper function to copy a grid
func copyGrid(src, dst [][]int) {
	for i := range src {
		copy(dst[i], src[i])
	}
}

// Helper function to check if two grids are equal
func equalGrid(a, b [][]int) bool {
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}


func Test_calculateTotalBrightness(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "All lights off",
			grid: initializeGrid(size),
			want: 0,
		},
		{
			name: "Turn on some lights",
			grid: func() [][]int {
				g := initializeGrid(size)
				for x := 0; x <= 2; x++ {
					for y := 0; y <= 2; y++ {
						g[x][y] = 1
					}
				}
				return g
			}(),
			want: 9, // 3x3 lights each with brightness 1
		},
		{
			name: "Turn off and toggle lights",
			grid: func() [][]int {
				g := initializeGrid(size)
				for x := 0; x <= 2; x++ {
					for y := 0; y <= 2; y++ {
						g[x][y] = 1
					}
				}
				// Turn off all lights
				for x := 0; x <= 2; x++ {
					for y := 0; y <= 2; y++ {
						g[x][y]--
					}
				}
				// Toggle the middle light
				g[1][1] += 2
				return g
			}(),
			want: 2, // Only the middle light with brightness 2
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateTotalBrightness(tt.grid); got != tt.want {
				t.Errorf("calculateTotalBrightness() = %v, want %v", got, tt.want)
			}
		})
	}
}