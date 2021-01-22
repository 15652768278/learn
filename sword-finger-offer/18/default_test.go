package main

import (
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "示例1",
			args: args{
				head: BuildListNode([]int{4, 5, 1, 9}),
				val:  5,
			},
			want: []int{4, 1, 9},
		},
		{
			name: "示例2",
			args: args{
				head: BuildListNode([]int{4, 5, 1, 9}),
				val:  1,
			},
			want: []int{4, 5, 9},
		},
		{
			name: "示例3",
			args: args{
				head: BuildListNode([]int{4, 5, 1, 9}),
				val:  10,
			},
			want: []int{4, 5, 1, 9},
		},
		{
			name: "示例4",
			args: args{
				head: BuildListNode([]int{4, 5, 1, 9}),
				val:  4,
			},
			want: []int{5, 1, 9},
		},
		{
			name: "示例5",
			args: args{
				head: BuildListNode([]int{4, 5, 1, 9}),
				val:  9,
			},
			want: []int{4, 5, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(tt.args.head, tt.args.val); !reflect.DeepEqual(got.ToArray(), tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got.ToArray(), tt.want)
			}
		})
	}
}
