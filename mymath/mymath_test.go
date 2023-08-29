package mymath

import "testing"

func TestSqrt(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Sqrt",
			args: args{
				x: 9,
			},
			want: 3,
		},{
			name: "Sqrt",
			args: args{
				x: 81,
			},
			want: 9,
		},{
			name: "Sqrt",
			args: args{
				x: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sqrt(tt.args.x); got != tt.want {
				t.Errorf("Sqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
