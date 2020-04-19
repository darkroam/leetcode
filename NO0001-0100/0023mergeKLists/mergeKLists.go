package mergeKLists

//23. 合并K个排序链表
//合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
//示例:
//	输入:
//	[
//	  1->4->5,
//	  1->3->4,
//	  2->6
//	]
//输出: 1->1->2->3->4->4->5->6
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

func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	for len(lists) > 1 {
		var i = 0
		var j = len(lists) - 1

		for i < j {
			lists[i] = mergeLists(lists[i], lists[j])
			i++
			j--
		}

		if i == j {
			lists = lists[:i+1]
		} else {
			lists = lists[:i]
		}
	}

	return lists[0]
}

func mergeLists(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			cur, l1 = cur.Next, l1.Next
		} else {
			cur.Next = l2
			cur, l2 = cur.Next, l2.Next
		}
	}
	if l1 == nil {
		cur.Next = l2
	} else {
		cur.Next = l1
	}
	return head.Next
}

func mergeKLists2(lists []*ListNode) *ListNode {
	head := &ListNode{}
	tmp := &ListNode{}
	var tmpIndex int
	cur := head

	for {
		count := 0
		for _, v := range lists {
			if v != nil {
				break
			}
			count++
		}
		if count == len(lists) {
			break
		}
		for i, v := range lists {
			if tmp.Next == nil {
				tmp.Next = v
				tmpIndex = i
				continue
			}
			if v == nil {
				count++
				continue
			}
			if tmp.Next.Val > v.Val {
				tmp.Next = v
				tmpIndex = i
			}
		}
		if tmp.Next != nil {
			cur.Next = tmp.Next
			cur = cur.Next
			tmp.Next = nil
			lists[tmpIndex] = lists[tmpIndex].Next
		}
	}

	return head.Next
}

//space: O(?)
//time : O(?)
