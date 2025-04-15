package main

// 链表
// 力扣：https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci/

// 方法一：暴力
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	for headA != nil {
// 		demo := headB // 要使用中间变量，不能直接用headB
// 		for demo != nil {
// 			if demo == headA {
// 				return demo
// 			}
// 			demo = demo.Next
// 		}
//
// 		headA = headA.Next
// 	}
// 	return nil
// }

// 方法二：计算长度
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	demoA, demoB := headA, headB
	for demoA != nil {
		demoA = demoA.Next
		lenA++
	}
	for demoB != nil {
		demoB = demoB.Next
		lenB++
	}

	demoA, demoB = headA, headB
	// 让两条链表到达长度持平的位置
	if lenA > lenB {
		for i := 1; i <= lenA-lenB; i++ {
			demoA = demoA.Next
		}
	} else {
		for i := 1; i <= lenB-lenA; i++ {
			demoB = demoB.Next
		}
	}

	for demoA != demoB {
		demoA = demoA.Next
		demoB = demoB.Next
	}
	return demoA
}

// 方法三：哈希集合
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	vis := map[*ListNode]bool{}
// 	for tmp := headA; tmp != nil; tmp = tmp.Next {
// 		vis[tmp] = true
// 	}
//
// 	for tmp := headB; tmp != nil; tmp = tmp.Next {
// 		if vis[tmp] {
// 			return tmp
// 		}
// 	}
// 	return nil
// }
