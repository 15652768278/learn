package main

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "示例1",
			args: args{
				head: BuildListNode([]int{1, 2, 3, 4, 5}),
				n:    2,
			},
			want: BuildListNode([]int{1, 2, 3, 5}),
		},
		{
			name: "示例2",
			args: args{
				head: BuildListNode([]int{1, 2, 3, 4, 5}),
				n:    1,
			},
			want: BuildListNode([]int{1, 2, 3, 4}),
		},
		{
			name: "示例3",
			args: args{
				head: BuildListNode([]int{1, 2, 3, 4, 5}),
				n:    5,
			},
			want: BuildListNode([]int{2, 3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
