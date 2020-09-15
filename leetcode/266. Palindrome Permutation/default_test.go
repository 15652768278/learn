package main

import "testing"

func Test_canPermutePalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "示例1",
			args: args{
				s: "code",
			},
			want: false,
		},
		{
			name: "示例2",
			args: args{
				s: "aab",
			},
			want: true,
		},
		{
			name: "示例3",
			args: args{
				s: "carerac",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPermutePalindrome(tt.args.s); got != tt.want {
				t.Errorf("canPermutePalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
