package main

import (
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head *ListNode
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "示例1",
			args: args{
				head: BuildListNode([]int{1, 2, 3, 4, 5}),
				m:    2,
				n:    4,
			},
			want: BuildListNode([]int{1, 4, 3, 2, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.m, tt.args.n); !got.Equals(tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got.ToArray(), tt.want.ToArray())
			}
		})
	}
}
