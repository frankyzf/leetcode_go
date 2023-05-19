package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RandomizedSet(t *testing.T) {
	p := Constructor()
	assert.Equal(t, p.Insert(1), true, "1")
	assert.Equal(t, p.Remove(2), false, "2")
	assert.Equal(t, p.Insert(2), true, "3")
	assert.Contains(t, []int{1, 2}, p.GetRandom(), "4")
	assert.Equal(t, p.Remove(1), true, "4")
	assert.Equal(t, p.Insert(2), false, "5")
	assert.Contains(t, []int{2}, p.GetRandom(), "6")
}
