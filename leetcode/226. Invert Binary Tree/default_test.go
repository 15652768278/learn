package main

import (
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "示例1",
			args: args{
				root: BuildTreeNode([]int{4, 2, 7, 1, 3, 6, 9}),
			},
			want: []int{4, 7, 2, 9, 6, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got.ToArray(), tt.want) {
				t.Errorf("invertTree() = %v, want %v", got.ToArray(), tt.want)
			}
		})
	}
}
