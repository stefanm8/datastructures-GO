package datastructures

import (
	"fmt"
	"testing"
)

var List = NewLinkedList()

func TestLinkedListAppend(t *testing.T) {
	List.Append(1)
	List.Append(2.23)
	List.Append("s")
	if List.Len() != 3 {
		t.Error("Missing elements")
	}
	if List.Head().Data() != "s" {
		t.Error("Invalid head")
	}

	if List.Tail().Data() != 1 {
		t.Error("Invalid tail")
	}
}

func TestLinkedListTruncate(t *testing.T) {
	List.Append(1)
	c := List.Len()
	List.Truncate()
	if List.Head() != nil && c == List.Len() && List.Tail() != nil {
		t.Error("Truncate does not remove items")

	}

}

func TestLinkedListPop(t *testing.T) {
	List.Append(1)
	List.Append(2)
	List.Append(3)
	List.Append(4)
	List.Append(5)

	c := List.Len()
	n := List.PopHead()
	if c == List.Len() {
		t.Error("Element was not removed")
	}
	if n != 5 {
		t.Error("Invalid head pop")
	}

	n = List.PopTail()
	if n != 1 {
		t.Error("Invalid tail pop")

	}

	n = List.Pop(3)
	for e := List.Head(); e != nil; e = e.Next() {
		if e.Data() == 3 {
			t.Error("Element was not Removed")
		}

	}

}

func TestLinkedListTraversal(t *testing.T) {
	List.Truncate()
	a := []int{}
	for i := 0; i < 100; i++ {
		List.Append(i)
		a = append(a, i)
	}
	i := 0
	for n := List.Tail(); n != nil; n = n.Next() {
		i++
		e := n.Prev()
		fmt.Println(e.Data())
		fmt.Println(a[i])
		if a[i] != e.Data() {
			t.Error("Invalid order")
		}

	}

}

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	e := q.Dequeue()

	if e != 1 {
		fmt.Errorf("Queue not fifo")
	}
	if q.Len() != 2 {
		fmt.Errorf("Iteme not removed from queue")
	}

}

func TestStack(t *testing.T) {
	s := NewStack()
	s.Append(1)
	s.Append(2)
	s.Append(3)
	e := s.Pop()
	if e != 3 {
		fmt.Errorf("Stack not LIFO")
	}

	if s.Len() != 2 {
		fmt.Errorf("Iteme not removed from stack")
	}

}
