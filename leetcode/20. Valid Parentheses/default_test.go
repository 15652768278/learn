package main

import "testing"

func Test_isValid(t *testing.T) {
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
				s: "()",
			},
			want: true,
		},
		{
			name: "示例2",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "示例3",
			args: args{
				s: "(]",
			},
			want: false,
		},
		{
			name: "示例4",
			args: args{
				s: "([)]",
			},
			want: false,
		},
		{
			name: "示例5",
			args: args{
				s: "{[]}",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
