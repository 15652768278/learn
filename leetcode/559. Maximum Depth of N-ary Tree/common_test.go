package main

import (
	"reflect"
	"testing"
)

func TestBuildNTreeNode(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "示例1",
			args: args{
				nums: []int{1, null, 3, 2, 4, null, 5, 6},
			},
			want: []int{1, null, 3, 2, 4, null, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildNTreeNode(tt.args.nums)
			println(1)
			if !reflect.DeepEqual(got.ToArray(), tt.want) {
				t.Errorf("BuildNTreeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
