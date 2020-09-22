package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "示例1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			minStack := Constructor()
			minStack.Push(-2)
			minStack.Push(0)
			minStack.Push(-3)
			if got := minStack.GetMin(); !reflect.DeepEqual(got, -3) {
				t.Errorf("GetMin() = %v, want %v", got, -3)
			}
			minStack.Pop()
			if got := minStack.Top(); !reflect.DeepEqual(got, 0) {
				t.Errorf("Top() = %v, want %v", got, 0)
			}
			if got := minStack.GetMin(); !reflect.DeepEqual(got, -2) {
				t.Errorf("GetMin() = %v, want %v", got, -2)
			}
		})
	}
}
