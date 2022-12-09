package main

import "fmt"

func main() {
	list1 := ListNode{Val: 11}

	result := mergeTwoLists(&list1, nil)

	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 != nil {
		return list2
	}
	if list1 != nil && list2 == nil {
		return list1
	}
	if list1 == nil && list2 == nil {
		return nil
	}
	newList := ListNode{}

	if list1.Val <= list2.Val {
		newList.Val = list1.Val
		list1 = list1.Next
		newList.Next = mergeTwoLists(list1, list2)
	} else {
		newList.Val = list2.Val
		list2 = list2.Next
		newList.Next = mergeTwoLists(list1, list2)
	}

	return &newList
}
