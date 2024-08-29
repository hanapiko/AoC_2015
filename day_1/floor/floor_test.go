package floor

import "testing"

func TestPosition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "First basement entry",
			args: args{s: "()())"},
			want: 5, 
		},
		{
			name: "No basement entry",
			args: args{s: "((()))"},
			want: 0, 
		},	
		{
			name: "first floor",
			args: args{s: "()()"},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Position(tt.args.s); got != tt.want {
				t.Errorf("Position() = %v, want %v", got, tt.want)
			}
		})
	}
}
