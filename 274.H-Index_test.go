package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hIndex(t *testing.T) {
	citations := []int{3, 0, 6, 1, 5}
	got := hIndex(citations)
	want := 3
	assert.Equal(t, got, want, "1")

	citations = []int{1, 3, 1}
	got = hIndex(citations)
	want = 1
	assert.Equal(t, got, want, "2")
}
