package main

import "testing"

func Test_forbidden(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"No forbidden substrings", args{"hello"}, true},
		{"contain ab", args{"helablo"}, false},
		{"contain cd", args{"hellocd"}, false},
		{"contain pq", args{"pqhello"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := forbidden(tt.args.s); got != tt.want {
				t.Errorf("forbidden() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_double(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"conains aa", args{"helloaa"}, true},
		{"contain qq", args{"heqqlablo"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := double(tt.args.s); got != tt.want {
				t.Errorf("double() = %v, want %v", got, tt.want)
			}
		})
	}
}
