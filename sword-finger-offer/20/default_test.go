package main

import "testing"

func Test_isNumber(t *testing.T) {
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
				s: "+100",
			},
			want: true,
		},
		{
			name: "示例2",
			args: args{
				s: "5e2",
			},
			want: true,
		},
		{
			name: "示例3",
			args: args{
				s: "-123",
			},
			want: true,
		},
		{
			name: "示例4",
			args: args{
				s: "3.1416",
			},
			want: true,
		},
		{
			name: "示例5",
			args: args{
				s: "-1E-16",
			},
			want: true,
		},
		{
			name: "示例6",
			args: args{
				s: "0123",
			},
			want: true,
		},
		{
			name: "示例7",
			args: args{
				s: "12e",
			},
			want: false,
		},
		{
			name: "示例8",
			args: args{
				s: "1a3.14",
			},
			want: false,
		},
		{
			name: "示例9",
			args: args{
				s: "1.2.3",
			},
			want: false,
		},
		{
			name: "示例10",
			args: args{
				s: "+-5",
			},
			want: false,
		},
		{
			name: "示例11",
			args: args{
				s: "12e+5.4",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.s); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}