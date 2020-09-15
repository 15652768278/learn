package main

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "无重复",
			args: args{
				nums: []int{0, 1, 2, 3},
			},
			want: 4,
		},
		{
			name: "有1重复",
			args: args{
				nums: []int{0, 0, 2, 3},
			},
			want: 3,
		},
		{
			name: "有2重复",
			args: args{
				nums: []int{0, 0, 0, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
