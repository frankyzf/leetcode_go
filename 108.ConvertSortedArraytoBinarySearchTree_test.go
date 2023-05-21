package main

import (
	"reflect"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				nums: []int{-10, -3, 0, 5, 9},
			},
			want: &TreeNode{
				Val:   0,
				Left:  &TreeNode{Val: -3, Left: &TreeNode{Val: -10}, Right: nil},
				Right: &TreeNode{Val: 9, Left: &TreeNode{Val: 5}, Right: nil},
			},
		},
		{
			name: "test 2",
			args: args{
				nums: []int{1, 3},
			},
			want: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: &TreeNode{Val: 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
