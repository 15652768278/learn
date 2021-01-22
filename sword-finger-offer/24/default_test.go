package main

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
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
				head: BuildListNode([]int{1, 2, 3, 4, 5}),
			},
			want: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got.ToArray(), tt.want) {
				t.Errorf("reverseList() = %v, want %v", got.ToArray(), tt.want)
			}
		})
	}
}
