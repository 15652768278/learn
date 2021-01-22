package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
func copyRandomList(head *Node) *Node {
	sentry := &Node{Next: head}
	if head == nil {
		return sentry.Next
	}
	cacheMap := make(map[*Node]*Node)
	for c1, c2 := head, sentry; c1 != nil; c1, c2 = c1.Next, c2.Next {
		copyNode := &Node{Val: c1.Val}
		cacheMap[c1] = copyNode
		c2.Next = copyNode
	}
	for c1, c2 := head, sentry.Next; c1 != nil; c1, c2 = c1.Next, c2.Next {
		if c1.Random != nil {
			c2.Random = cacheMap[c1.Random]
		}
	}
	return sentry.Next
}
