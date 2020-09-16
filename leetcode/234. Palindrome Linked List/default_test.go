package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "示例1",
			args: args{
				head: BuildListNode([]int{1, 2, 3, 4, 5}),
			},
			want: false,
		},
		{
			name: "示例2",
			args: args{
				head: BuildListNode([]int{1, 2, 3, 2, 1}),
			},
			want: true,
		},
		{
			name: "示例3",
			args: args{
				head: BuildListNode([]int{1, 1, 2, 1}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
