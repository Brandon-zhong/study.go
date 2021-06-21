package main

import (
	"fmt"
	"testing"
)

/*
206. 反转链表
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func generateLinkList() *ListNode {
	var head = ListNode{Val: 0}
	var ptr *ListNode = &head
	for i := 1; i < 5; i++ {
		ptr.Next = &ListNode{Val: i}
		ptr = ptr.Next
	}
	return &head
}

func printLinkList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val, " ")
		node = node.Next
	}
}

//指针法
func reverseList1(head *ListNode) *ListNode {

	var newPtr *ListNode
	var oldPtr, oldPtr2 *ListNode = head, head

	for oldPtr != nil {
		oldPtr = oldPtr.Next
		oldPtr2.Next = newPtr
		newPtr = oldPtr2
		oldPtr2 = oldPtr
	}
	return newPtr
}

//递归法
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	start, end := demo(head)
	end.Next = nil
	return start
}

func demo(head *ListNode) (*ListNode, *ListNode) {
	if head.Next == nil {
		return head, head
	}
	start, end := demo(head.Next)
	end.Next = head
	end = head
	return start, end
}

func TestSolution_206(t *testing.T) {

	node := generateLinkList()
	newNode := reverseList2(node)
	printLinkList(newNode)

}
