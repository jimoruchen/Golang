package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	maps := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		maps[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		newNode := maps[cur]
		newNode.Next = maps[cur.Next]
		newNode.Random = maps[cur.Random]
		cur = cur.Next
	}
	return maps[head]
}

var maps map[*Node]*Node

func deepCopy(head *Node) *Node {
	if head == nil {
		return nil
	}
	if n, ok := maps[head]; ok {
		return n
	}
	newNode := &Node{Val: head.Val}
	maps[head] = newNode
	newNode.Next = deepCopy(head.Next)
	newNode.Random = deepCopy(head.Random)
	return newNode
}

func copyRandomList1(head *Node) *Node {
	maps = make(map[*Node]*Node)
	return deepCopy(head)
}
