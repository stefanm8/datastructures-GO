package datastructures

type Node struct {
	data       interface{}
	next, prev *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	len  int
}

func (l *LinkedList) Head() *Node {
	return l.head
}

func (l *LinkedList) Tail() *Node {
	return l.tail
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (n *Node) Data() interface{} {
	return n.data
}

func (l *LinkedList) Len() int {
	return l.len
}

func (l *LinkedList) Pop(d interface{}) interface{} {
	var temp *Node
	for n := l.head; n != nil; n = n.next {
		if n.data == d {
			temp = n
			l.remove(n)
			break
		}
	}
	return temp.data
}

func (l *LinkedList) PopHead() interface{} {
	t := l.head
	l.remove(l.head)
	return t.data
}

func (l *LinkedList) PopTail() interface{} {
	t := l.tail
	l.remove(l.tail)
	return t.data
}

func (l *LinkedList) Append(d interface{}) {
	n := &Node{data: d}
	l.insert(n)
}

func (l *LinkedList) Truncate() {
	for n := l.Head(); n != nil; n = n.Next() {
		n.data = nil
		n.prev = nil
		n.next = nil
	}
	l.len = 0
	l.head = nil
	l.tail = nil

}

func (l *LinkedList) insert(n *Node) {
	n.next = l.head
	if l.head != nil {
		l.head.prev = n
	} else {
		l.tail = n
	}
	l.head = n
	l.len++
}
func (l *LinkedList) remove(n *Node) {
	temp := n
	p := temp.prev
	nex := temp.next
	if p == nil && nex == nil {
		l.head = nil
	}

	if nex != nil {
		nex.prev = p
		if p == nil {
			l.head = nex
		}

	}
	if p != nil {
		p.next = nex
	}
	l.len--
	n.next = nil
	n.prev = nil
}

func NewLinkedList() *LinkedList {
	return &LinkedList{len: 0}
}
