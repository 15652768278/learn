package main

import (
	"reflect"
	"testing"
)

func Test_reversePrint(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "示例1",
			args: args{
				head: BuildListNode([]int{1, 3, 2}),
			},
			want: []int{2, 3, 1},
		},
		{
			name: "示例2",
			args: args{
				head: BuildListNode([]int{1}),
			},
			want: []int{1},
		},
		{
			name: "示例3",
			args: args{
				head: BuildListNode([]int{}),
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrint(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrint() = %v, want %v", got, tt.want)
			}
		})
	}
}
