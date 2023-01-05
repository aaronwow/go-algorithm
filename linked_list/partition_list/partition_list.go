package algorithm

import "gandi.icu/go-algorithm/utils"

func PartitionList(head *utils.ListNode, x int) *utils.ListNode {
	var sh, st, eh, et, bh, bt *utils.ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = nil
		if cur.Val < x {
			if sh == nil {
				sh, st = cur, cur
			} else {
				st.Next = cur
				st = cur
			}
		}

		if cur.Val == x {
			if eh == nil {
				eh, et = cur, cur
			} else {
				et.Next = cur
				et = cur
			}
		}

		if cur.Val > x {
			if bh == nil {
				bh, bt = cur, cur
			} else {
				bt.Next = cur
				bt = cur
			}
		}
		cur = next
	}

	if st != nil {
		st.Next = eh
		if et == nil {
			et = st
		}
	}

	if et != nil {
		et.Next = bh
	}
	var res *utils.ListNode
	if sh != nil {
		res = sh
	} else if eh != nil {
		res = eh
	} else {
		res = bh
	}
	return res
}
