package leetcode

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
//
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//
//你可以假设除了数字 0 之外，这两个数都不会以 0开头。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/add-two-numbers
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//  2 4 3
//+ 5 6 4
//= 7 0 8

//  2 4 6
//+ 5 6 4
//= 7 0 0 1

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//s := 0
	//l3 := &ListNode{Val: 0}
	//var result *ListNode
	//for l1.Next != nil || l2.Next != nil {
	//	l3.Next = &ListNode{Val: 0}
	//	if l1.Next == nil && l2.Next != nil {
	//		l1.Next = &ListNode{Val: 0}
	//	}
	//	if l1.Next != nil && l2.Next == nil {
	//		l2.Next = &ListNode{Val: 0}
	//	}
	//	v := l1.Val + l2.Val + s
	//	s = 0
	//	if v > 10 {
	//		l3.Val, s = v % 10,v/10
	//	} else {
	//		l3.Val = v
	//	}
	//	l1 = l1.Next
	//	l2 = l2.Next
	//	l3 = l3.Next
	//	if l1.Next == nil && l2.Next == nil && s == 10 {
	//		l3.Next = &ListNode{Val: 1}
	//	}
	//}
	return l1
}
