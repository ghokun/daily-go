package dailygo

import "testing"

func Test_20220320(t *testing.T) {
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
				arr: []int{2, 4, 6, 2, 5},
			},
			want: 13,
		},
		{
			name: "Test 1",
			args: args{
				arr: []int{5, 1, 1, 5},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _20220320(tt.args.arr); got != tt.want {
				t.Errorf("_20220320() = %v, want %v", got, tt.want)
			}
		})
	}
}
