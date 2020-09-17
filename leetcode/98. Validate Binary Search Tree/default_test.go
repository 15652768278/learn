package main

import "testing"

func Test_isValidBST(t *testing.T) {
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
				root: BuildTreeNode([]int{2, 1, 3}),
			},
			want: true,
		},
		{
			name: "示例2",
			args: args{
				root: BuildTreeNode([]int{5, 1, 4, null, null, 3, 6}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
