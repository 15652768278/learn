package main

import "testing"

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//{
		//	name: "示例1",
		//	args: args{
		//		s: "abcdefg",
		//		k: 2,
		//	},
		//	want: "bacdfeg",
		//},
		{
			name: "示例2",
			args: args{
				s: "a",
				k: 2,
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
