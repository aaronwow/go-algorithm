package algorithm

import (
	"sort"
	"testing"

	"gandi.icu/go-algorithm/utils"
)

func TestShellSort(t *testing.T) {
	arr := utils.GenerateRandomInts(10)
	ShellSort(arr)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("ShellSort failed")
	}
}
