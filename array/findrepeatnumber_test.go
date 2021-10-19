package array

import "testing"

func Test_findRepeatNumberMap(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{
				nums: []int{2, 3, 1, 0, 2, 5},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatNumberMap(tt.args.nums); got != tt.want {
				t.Errorf("findRepeatNumberMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findRepeatNumberIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{
				nums: []int{2, 3, 1, 0, 2, 5},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatNumberIndex(tt.args.nums); got != tt.want {
				t.Errorf("findRepeatNumberIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findRepeatNumberSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ok",
			args: args{
				nums: []int{2, 3, 1, 0, 2, 5},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatNumberSort(tt.args.nums); got != tt.want {
				t.Errorf("findRepeatNumberSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
