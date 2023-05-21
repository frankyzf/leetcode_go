package main

import "testing"

func Test_isPalindrome2(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			name: "test 2",
			args: args{
				x: -121,
			},
			want: false,
		},
		{
			name: "test 3",
			args: args{
				x: 10,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome2(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
