package main

import "testing"

func Test_cuttingRope(t *testing.T) {
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
			want: 1,
		},
		{
			name: "示例2",
			args: args{
				n: 10,
			},
			want: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cuttingRope(tt.args.n); got != tt.want {
				t.Errorf("cuttingRope() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cuttingRope2(t *testing.T) {
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
			want: 1,
		},
		{
			name: "示例2",
			args: args{
				n: 10,
			},
			want: 36,
		},
		{
			name: "示例3",
			args: args{
				n: 120,
			},
			want: 953271190,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cuttingRope2(tt.args.n); got != tt.want {
				t.Errorf("cuttingRope2() = %v, want %v", got, tt.want)
			}
		})
	}
}