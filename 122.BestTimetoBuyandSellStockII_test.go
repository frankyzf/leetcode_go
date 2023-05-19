package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProfit2(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	got := maxProfit2(prices)
	want := 7
	assert.Equal(t, got, want, "equal")
}
