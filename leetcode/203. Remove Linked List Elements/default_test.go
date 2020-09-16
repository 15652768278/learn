package main

import (
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "示例1",
			args: args{
				head: BuildListNode([]int{1, 2, 6, 3, 4, 5}),
				val:  6,
			},
			want: BuildListNode([]int{1, 2, 3, 4, 5}),
		},
		{
			name: "示例2",
			args: args{
				head: BuildListNode([]int{1, 2, 6, 3, 4, 5, 6}),
				val:  6,
			},
			want: BuildListNode([]int{1, 2, 3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !got.Equals(tt.want) {
				t.Errorf("removeElements() = %v, want %v", got.ToArray(), tt.want.ToArray())
			}
		})
	}
}
