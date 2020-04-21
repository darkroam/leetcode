package reverseKGroup

//25. K 个一组翻转链表
//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//k 是一个正整数，它的值小于或等于链表的长度。
//如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//示例：
//给你这个链表：1->2->3->4->5
//	当 k = 2 时，应当返回: 2->1->4->3->5
//	当 k = 3 时，应当返回: 3->2->1->4->5
//说明：
//	你的算法只能使用常数的额外空间。
//	你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/add-two-numbers
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

import "fmt"

func PrintList(list *ListNode) {
	p := list
	for p != nil {
		//fmt.Printf("%d-%v-%p ", p.Val, p, p.Next)
		fmt.Printf("%#v", p.Val)
		p = p.Next
	}
	fmt.Println()
}

func ListToArray(list *ListNode) []int {
	tmp := make([]int, 0, 10)
	if list == nil {
		return tmp
	}
	for list.Next != nil {
		tmp = append(tmp, list.Val)
		list = list.Next
	}
	tmp = append(tmp, list.Val)
	return tmp
}

func ArrayToList(nums []int) *ListNode {
	if len(nums) <= 0 {
		return nil
	}
	head := &ListNode{
		Val:  nums[0],
		Next: nil,
	}
	tail := head
	for _, v := range nums[1:] {
		node := new(ListNode)
		tail.Next = node
		tail = tail.Next
		tail.Val = v
	}
	return head
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	cur := &ListNode{}
	tmp1, tmp2, tmp3, res, ptr := cur, cur, cur, cur, cur

	for {
		ptr, tmp3 = head, head
		for i := 0; i < k; i++ {
			if head == nil {
				cur.Next = ptr
				return res.Next
			}
			head = head.Next
		}
		for i := 0; i < k; i++ {
			tmp2 = ptr.Next
			tmp1 = cur.Next
			cur.Next = ptr
			ptr.Next = tmp1
			ptr = tmp2
		}
		cur = tmp3
	}
}

//fmt.Print("2 res  ")
//PrintList(res.Next)
//fmt.Print("2 head ")
//PrintList(head)
//fmt.Print("2 cur  ")
//PrintList(cur)

//space: O(?)
//time : O(?)
