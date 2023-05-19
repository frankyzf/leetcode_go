package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MergeSortedArray(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	expect := []int{1, 2, 2, 3, 5, 6}
	res := reflect.DeepEqual(nums1, expect)
	assert.Equal(t, res, true, "equal")
}
