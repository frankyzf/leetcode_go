package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotate(t *testing.T) {
	nums := []int{-1, -100, 3, 99}
	k := 2
	rotate(nums, k)
	res := reflect.DeepEqual(nums, []int{3, 99, -1, -100})
	assert.Equal(t, res, true, "eequal")
}
