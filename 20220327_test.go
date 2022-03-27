package dailygo

import (
	"testing"
)

func Test_20220327(t *testing.T) {
	type args struct {
		size int
		logs []int
		last int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				size: 3,
				logs: []int{1, 2, 3, 4, 5},
				last: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _20220327(tt.args.size, tt.args.logs, tt.args.last); got != tt.want {
				t.Errorf("_20220327() = %v, want %v", got, tt.want)
			}
		})
	}
}
