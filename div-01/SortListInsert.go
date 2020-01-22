package main

import (
	"fmt"
)

type NodeI struct {
	Data int
	Next *NodeI
}
type List struct {
	Head *NodeI
	Tail *NodeI
}

func PrintList(l *NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}
func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}
	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}
// func main() {
// 	var link *NodeI
// 	link = listPushBack(link, 0)
// 	link = listPushBack(link, 2)
// 	link = listPushBack(link, 18)
// 	link = listPushBack(link, 33)
// 	link = listPushBack(link, 37)
// 	link = listPushBack(link, 37)
// 	link = listPushBack(link, 39)
// 	link = listPushBack(link, 52)
// 	link = listPushBack(link, 53)
// 	link = listPushBack(link, 57)

// 	// link = listPushBack(link, 9)
// 	link = SortListInsert(link, 53)
// 	// link = SortListInsert(link, 2)
// 	PrintList(link)
// }
func SortListInsert(l *NodeI, data_ref int) *NodeI {
	LinkI := &NodeI{}
	LinkI.Data = data_ref
	LinkI.Next = nil
	if l == nil || l.Data >= LinkI.Data {
		LinkI.Next = l
		return LinkI
	} else {
		oldL := l
		for oldL.Next != nil && (LinkI.Data > oldL.Next.Data) {
			oldL = oldL.Next
		}
		LinkI.Next = oldL.Next
		oldL.Next = LinkI
	}
	return l
}