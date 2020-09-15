package main

import "testing"

func Test_reverseVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "示例1",
			args: args{
				s: "hello",
			},
			want: "holle",
		},
		{
			name: "示例2",
			args: args{
				s: "leetcode",
			},
			want: "leotcede",
		},
		{
			name: "示例3",
			args: args{
				s: "aA",
			},
			want: "Aa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.args.s); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
