package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_majorityElement(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	assert.Equal(t, majorityElement(nums), 2, "equal")
}
