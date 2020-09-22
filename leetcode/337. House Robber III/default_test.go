package main

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例1",
			args: args{
				root: BuildTreeNode([]int{3, 2, 3, null, 3, null, 1}),
			},
			want: 7,
		},
		{
			name: "示例2",
			args: args{
				root: BuildTreeNode([]int{3, 4, 5, 1, 3, null, 1}),
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.root); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
