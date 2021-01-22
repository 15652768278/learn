package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	/**
	 * Your Codec object will be instantiated and called as such:
	 * ser := Constructor();
	 * deser := Constructor();
	 * data := ser.serialize(root);
	 * ans := deser.deserialize(data);
	 */
	c := Constructor()
	origin := []int{1, 2, 3, NULL, NULL, 4, 5}
	after := c.deserialize(c.serialize(BuildTreeNode(origin))).ToArray()
	if !reflect.DeepEqual(origin, after) {
		t.Errorf("Constructor() = %v, want %v", after, origin)
	}
}
