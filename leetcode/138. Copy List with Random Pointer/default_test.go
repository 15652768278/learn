package main

import (
	"reflect"
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "示例1",
			args: args{
				head: BuildNode([][]int{
					{7, NULL},
					{13, 0},
					{11, 4},
					{10, 2},
					{1, 0},
				}),
			},
			want: []int{7, 13, 11, 10, 1, 0, 4, 2, 0},
		},
		{
			name: "示例2",
			args: args{
				head: BuildNode([][]int{
					{1, 1},
					{2, 1},
				}),
			},
			want: []int{1, 2, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyRandomList(tt.args.head); !reflect.DeepEqual(got.ToArray(), tt.want) {
				t.Errorf("copyRandomList() = %v, want %v", got.ToArray(), tt.want)
			}
		})
	}
}
