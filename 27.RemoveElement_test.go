package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	res := removeElement(nums, val)
	assert.Equal(t, res, 2, "pass")
}
