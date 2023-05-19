package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates2(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	n := removeDuplicates2(nums)
	assert.Equal(t, n, 5, "equal")
	res := reflect.DeepEqual(nums[0:n], []int{1, 1, 2, 2, 3})
	assert.Equal(t, res, true, "equal")
}
