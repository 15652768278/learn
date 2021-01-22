package main

import "testing"

func Test_isSubStructure(t *testing.T) {
	type args struct {
		A *TreeNode
		B *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "示例1",
			args: args{
				A: BuildTreeNode([]int{1, 2, 3}),
				B: BuildTreeNode([]int{3, 1}),
			},
			want: false,
		},
		{
			name: "示例2",
			args: args{
				A: BuildTreeNode([]int{3, 4, 5, 1, 2}),
				B: BuildTreeNode([]int{4, 1}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubStructure(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("isSubStructure() = %v, want %v", got, tt.want)
			}
		})
	}
}
