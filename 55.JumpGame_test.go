package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	got := canJump(nums)
	want := true
	assert.Equal(t, got, want, "1")
}
