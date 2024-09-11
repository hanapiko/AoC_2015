// main_test.go
package main

import "testing"

func TestCalculateWrappingPaper(t *testing.T) {
    type args struct {
        l, w, h int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "Test Case 1",
            args: args{l: 1, w: 1, h: 10},
            want: 43, // 2*(1*1 + 1*10 + 10*1) + min(1*1, 1*10, 10*1) = 43
        },
        {
            name: "Test Case 2",
            args: args{l: 2, w: 3, h: 4},
            want: 58, // 2*(2*3 + 3*4 + 4*2) + min(2*3, 3*4, 4*2) = 58
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := calculateWrappingPaper(tt.args.l, tt.args.w, tt.args.h)
            if got != tt.want {
                t.Errorf("calculateWrappingPaper() = %d, want %d", got, tt.want)
            }
        })
    }
}

func TestCalculateRibbon(t *testing.T) {
    type args struct {
        l, w, h int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "Test Case 1",
            args: args{l: 1, w: 1, h: 10},
            want: 14,
        },
        {
            name: "Test Case 2",
            args: args{l: 2, w: 3, h: 4},
            want: 34,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := calculateRibbon(tt.args.l, tt.args.w, tt.args.h)
            if got != tt.want {
                t.Errorf("calculateRibbon() = %d, want %d", got, tt.want)
            }
        })
    }
}