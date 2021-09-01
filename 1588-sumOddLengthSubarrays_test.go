package leetcode

import "testing"

func Test_sumOddLengthSubarrays(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{arr: []int{1, 4, 2, 5, 3}}, 58},
		{"", args{arr: []int{1, 2}}, 3},
		{"", args{arr: []int{10, 11, 12}}, 66},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOddLengthSubarrays(tt.args.arr); got != tt.want {
				t.Errorf("sumOddLengthSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
