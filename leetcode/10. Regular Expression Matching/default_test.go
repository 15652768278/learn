package main

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "示例1",
			args: args{
				s: "aa",
				p: "a",
			},
			want: false,
		},
		{
			name: "示例2",
			args: args{
				s: "aa",
				p: "a*",
			},
			want: true,
		},
		{
			name: "示例3",
			args: args{
				s: "ab",
				p: ".*",
			},
			want: true,
		},
		{
			name: "示例4",
			args: args{
				s: "aab",
				p: "c*a*b",
			},
			want: true,
		},
		{
			name: "示例5",
			args: args{
				s: "mississippi",
				p: "mis*is*p*.",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
