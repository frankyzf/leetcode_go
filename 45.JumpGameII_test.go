package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_jump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	got := jump(nums)
	want := 2
	assert.Equal(t, got, want, "1")

	nums = []int{2, 3, 0, 1, 4}
	got = jump(nums)
	want = 2
	assert.Equal(t, got, want, "2")
}
