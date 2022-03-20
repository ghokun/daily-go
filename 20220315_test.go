package dailygo

import "testing"

func Test_20220315(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				arr: []int{3, 4, 2, 1},
			},
			want: 5,
		},
		{
			name: "Test 2",
			args: args{
				arr: []int{3, 4, -1, 1},
			},
			want: 2,
		},
		{
			name: "Test 3",
			args: args{
				arr: []int{1, 2, 0, 1},
			},
			want: 3,
		},
		{
			name: "Test 4",
			args: args{
				arr: []int{-1, -2, 0, -1},
			},
			want: 1,
		},
		{
			name: "Test 5",
			args: args{
				arr: []int{1, -1, -5, -3, 3, 4, 2, 8},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _20220315(tt.args.arr); got != tt.want {
				t.Errorf("_20220314() = %v, want %v", got, tt.want)
			}
		})
	}
}
