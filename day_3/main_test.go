package main

import "testing"

// TestHouse tests the House function.
func TestHouse(t *testing.T) {
    type args struct {
        s string
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "Single Move",
            args: args{s: "^"},
            want: 2, 
        },
        {
            name: "Back and Forth",
            args: args{s: "^v^v^v^v^v"},
            want: 2, 
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := House(tt.args.s); got != tt.want {
                t.Errorf("House() = %v, want %v", got, tt.want)
            }
        })
    }
}

// TestSantaRobo tests the SantaRobo function.
func TestSantaRobo(t *testing.T) {
    type args struct {
        s string
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "Back and Forth",
            args: args{s: "^>"},
            want: 3,
        },
        {
            name: "All Directions",
            args: args{s: "^>v<"},
            want: 3,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := SantaRobo(tt.args.s); got != tt.want {
                t.Errorf("SantaRobo() = %v, want %v", got, tt.want)
            }
        })
    }
}
