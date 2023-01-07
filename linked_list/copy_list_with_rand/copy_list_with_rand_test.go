package algorithm

import (
	"testing"

	"gandi.icu/go-algorithm/utils"
)

func TestCopyListWithRand(t *testing.T) {
	list := utils.GenerateRandomListWithRandLinedList(100)
	newList := CopyListWithRand(list)

	// compare value of two list
	n1, n2 := list, newList
	if (n1 == nil && n2 != nil) || (n1 != nil && n2 == nil) {
		t.Errorf("Fail: length of two list is different")
	}
	for n1 != nil && n2 != nil {
		// check node val
		if n1.Val != n2.Val {
			t.Errorf("Fail: value different")
		}
		// check rand node
		if (n1.Rand == nil && n2.Rand != nil) || (n1.Rand != nil && n2.Rand == nil) {
			t.Errorf("Fail: rand node is different")
		}
		if n1.Rand != nil && n2.Rand != nil {
			if n1.Rand.Val != n2.Rand.Val {
				t.Errorf("Fail: rand node val is different")
			}
		}
		n1 = n1.Next
		n2 = n2.Next

		if (n1 == nil && n2 != nil) || (n1 != nil && n2 == nil) {
			t.Errorf("Fail: length of two list is different")
		}
	}
}
func TestCopyListWithRandWithoutHashMap(t *testing.T) {
	list := utils.GenerateRandomListWithRandLinedList(100)
	newList := CopyListWithRandWithoutHashMap(list)

	// compare value of two list
	n1, n2 := list, newList
	if (n1 == nil && n2 != nil) || (n1 != nil && n2 == nil) {
		t.Errorf("Fail: length of two list is different")
	}
	for n1 != nil && n2 != nil {
		// check node val
		if n1.Val != n2.Val {
			t.Errorf("Fail: value different")
		}
		// check rand node
		if (n1.Rand == nil && n2.Rand != nil) || (n1.Rand != nil && n2.Rand == nil) {
			t.Errorf("Fail: rand node is different")
		}
		if n1.Rand != nil && n2.Rand != nil {
			if n1.Rand.Val != n2.Rand.Val {
				t.Errorf("Fail: rand node val is different")
			}
		}
		n1 = n1.Next
		n2 = n2.Next

		if (n1 == nil && n2 != nil) || (n1 != nil && n2 == nil) {
			t.Errorf("Fail: length of two list is different")
		}
	}
}
