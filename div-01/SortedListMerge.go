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

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	n1 = ListSort(n1)
	n2 = ListSort(n2)

	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	if n1.Data <= n2.Data {
		n1.Next = SortedListMerge(n1.Next, n2)
		return n1
	} else {
		n2.Next = SortedListMerge(n1, n2.Next)
		return n2
	}
}

func ListSort(l *NodeI) *NodeI {

	if l != nil {
		NewH := l
		for NewH.Next != nil {
			NewH2 := l
			for NewH2.Next != nil {
				if NewH2.Next.Data < NewH2.Data {
					NewH2.Data, NewH2.Next.Data = NewH2.Next.Data, NewH2.Data
				}
				NewH2 = NewH2.Next
			}
			NewH = NewH.Next
		}
	}
	return l
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
func main() {
	var link *NodeI
	var link2 *NodeI

	link = listPushBack(link, 0)
	link = listPushBack(link, 7)
	link = listPushBack(link, 39)
	link = listPushBack(link, 92)
	link = listPushBack(link, 97)
	link = listPushBack(link, 93)
	link = listPushBack(link, 91)
	link = listPushBack(link, 28)
	link = listPushBack(link, 64)
	

	link2 = listPushBack(link2, 80)
	link2 = listPushBack(link2, 23)
	link2 = listPushBack(link2, 27)
	link2 = listPushBack(link2, 30)
	link2 = listPushBack(link2, 85)
	link2 = listPushBack(link2, 81)
	link2 = listPushBack(link2, 75)
	link2 = listPushBack(link2, 70)
	PrintList(link2)
	PrintList(link)
	PrintList(SortedListMerge(link2, link))
}