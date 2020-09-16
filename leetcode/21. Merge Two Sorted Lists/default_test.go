package main

import (
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "示例1",
			args: args{
				l1: BuildListNode([]int{1, 2, 4}),
				l2: BuildListNode([]int{1, 3, 4}),
			},
			want: BuildListNode([]int{1, 1, 2, 3, 4, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !got.Equals(tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got.ToArray(), tt.want.ToArray())
			}
		})
	}
}
