package main

import "testing"

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "示例1",
			args: args{
				root: BuildTreeNode([]int{3, 9, 20, null, null, 15, 7}),
			},
			want: true,
		},
		{
			name: "示例1",
			args: args{
				root: BuildTreeNode([]int{1, 2, 2, 3, 3, null, null, 4, 4}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
