package main

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例1",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "示例2",
			args: args{
				n: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbStairs2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例1",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "示例2",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "示例3",
			args: args{
				n: 44,
			},
			want: 1134903170,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs2(tt.args.n); got != tt.want {
				t.Errorf("climbStairs2() = %v, want %v", got, tt.want)
			}
		})
	}
}
