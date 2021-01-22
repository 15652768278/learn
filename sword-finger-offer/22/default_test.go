package main

import (
	"reflect"
	"testing"
)

func Test_getKthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
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
				k:    2,
			},
			want: []int{4, 5},
		},
		{
			name: "示例2",
			args: args{
				head: BuildListNode([]int{1, 2, 3, 4, 5}),
				k:    5,
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKthFromEnd(tt.args.head, tt.args.k); !reflect.DeepEqual(got.ToArray(), tt.want) {
				t.Errorf("getKthFromEnd() = %v, want %v", got.ToArray(), tt.want)
			}
		})
	}
}
