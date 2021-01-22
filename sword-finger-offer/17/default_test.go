package main

import (
	"fmt"
	"testing"
)

func Test_printNumbers(t *testing.T) {
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
				n: 1,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printNumbers(tt.args.n); len(got) != tt.want {
				fmt.Printf("got %v", got)
				t.Errorf("printNumbers() len = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func Test_printNumbers2(t *testing.T) {
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
				n: 1,
			},
			want: 9,
		},
		{
			name: "示例2",
			args: args{
				n: 4,
			},
			want: 9999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printNumbers(tt.args.n); len(got) != tt.want {
				t.Errorf("printNumbers2() len = %v, want %v", len(got), tt.want)
			}
		})
	}
}
