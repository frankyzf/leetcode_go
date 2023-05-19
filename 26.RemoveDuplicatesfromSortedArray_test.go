package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}
	n := removeDuplicates(nums)
	assert.Equal(t, n, 2, "num")
	res := reflect.DeepEqual(nums[0:n], []int{1, 2})
	assert.Equal(t, res, true, "nums")
}
