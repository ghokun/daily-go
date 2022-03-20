package dailygo

import (
	"reflect"
	"testing"
)

func Test_20220313(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				numbers: []int{1, 2, 3, 4, 5},
			},
			want: []int{120, 60, 40, 30, 24},
		},
		{
			name: "Test 2",
			args: args{
				numbers: []int{3, 2, 1},
			},
			want: []int{2, 3, 6},
		},
		{
			name: "Test 3",
			args: args{
				numbers: []int{3, 0, 1},
			},
			want: []int{0, 3, 0},
		},
		{
			name: "Test 4",
			args: args{
				numbers: []int{2, 7, 4, 3, 1},
			},
			want: []int{84, 24, 42, 56, 168},
		},
		{
			name: "Test 5",
			args: args{
				numbers: []int{2, 7, 4, 3, 0},
			},
			want: []int{0, 0, 0, 0, 168},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _20220313(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_20220313() = %v, want %v", got, tt.want)
			}
		})
	}
}
