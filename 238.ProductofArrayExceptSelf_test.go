package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_productExecptSelf(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	got := productExceptSelf(nums)
	want := []int{24, 12, 8, 6}
	c := reflect.DeepEqual(got, want)
	assert.Equal(t, c, true, "1")

	nums = []int{-1, 1, 0, -3, 3}
	got = productExceptSelf(nums)
	want = []int{0, 0, 9, 0, 0}
	c = reflect.DeepEqual(got, want)
	assert.Equal(t, c, true, "2")
}
