package main

import (
	"reflect"
	"testing"
)

func TestPackage(t *testing.T) {
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
				nums: []int{1, 2, 3, 4, 5, 6},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Constructor(tt.args.nums)
			shuffle1 := got.Shuffle()
			shuffle2 := got.Shuffle()
			if reflect.DeepEqual(shuffle1, shuffle2) {
				t.Errorf("Shuffle1 = %v, Shuffle2 %v", shuffle1, shuffle2)
			}
			if reset := got.Reset(); !reflect.DeepEqual(reset, tt.want) {
				t.Errorf("Reset() = %v, want %v", reset, tt.want)
			}
		})
	}
}
