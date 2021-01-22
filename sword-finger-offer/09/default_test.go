package main

import (
	"reflect"
	"testing"
)

func TestCQueue(t *testing.T) {
	cq := Constructor()
	cq.AppendTail(3)
	if ret := cq.DeleteHead(); !reflect.DeepEqual(3, ret) {
		t.Errorf("except %d, actual %d", 3, ret)
	}
	if ret := cq.DeleteHead(); !reflect.DeepEqual(-1, ret) {
		t.Errorf("except %d, actual %d", -1, ret)
	}

	cq = Constructor()
	if ret := cq.DeleteHead(); !reflect.DeepEqual(-1, ret) {
		t.Errorf("except %d, actual %d", -1, ret)
	}
	cq.AppendTail(5)
	cq.AppendTail(2)
	if ret := cq.DeleteHead(); !reflect.DeepEqual(5, ret) {
		t.Errorf("except %d, actual %d", 5, ret)
	}
	if ret := cq.DeleteHead(); !reflect.DeepEqual(2, ret) {
		t.Errorf("except %d, actual %d", 2, ret)
	}
}
