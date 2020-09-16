package main

import "testing"

func Test_deleteNode(t *testing.T) {
	type args struct {
		node *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "示例1",
			args: args{
				node: BuildListNode([]int{4, 5, 1, 9}).Next, // 5
			},
			want: BuildListNode([]int{4, 1, 9}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteNode(tt.args.node)
			if tt.args.node.Equals(tt.want) {
				t.Errorf("deleteNode() = %v, want %v", tt.args.node.ToArray(), tt.want.ToArray())
			}
		})
	}
}
