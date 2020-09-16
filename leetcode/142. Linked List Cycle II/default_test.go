package main

import (
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "示例1",
			args: args{
				head: BuildCycleListNode([]int{1, 2, 3}, -1),
			},
			want: nil,
		},
		{
			name: "示例2",
			args: args{
				head: BuildCycleListNode([]int{3, 2, 0, -4}, 1),
			},
			want: &ListNode{Val: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := detectCycle(tt.args.head)
			if tt.want != nil && tt.want.Val != got.Val {
				t.Errorf("detectCycle() = %v, want %v", tt.want.Val, got.Val)
			}
		})
	}
}
