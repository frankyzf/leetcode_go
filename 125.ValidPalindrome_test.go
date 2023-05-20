package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// // TODO: Add test cases.
		// {
		// 	name: "test1",
		// 	args: args{
		// 		s: "A man, a plan, a canal: Panama",
		// 	},
		// 	want: true,
		// },
		// {
		// 	name: "test2",
		// 	args: args{
		// 		s: "race a car",
		// 	},
		// 	want: false,
		// },
		// {
		// 	name: "test3",
		// 	args: args{
		// 		s: " ",
		// 	},
		// 	want: true,
		// },
		{
			name: "test4",
			args: args{
				s: "0P",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
