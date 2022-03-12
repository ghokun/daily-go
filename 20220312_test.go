package dailygo

import "testing"

func Test_20220312(t *testing.T) {
	type args struct {
		number  int
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				number:  17,
				numbers: []int{10, 15, 3, 7},
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				number:  10,
				numbers: []int{5, 15, 3, 7},
			},
			want: true,
		},
		{
			name: "Test 3",
			args: args{
				number:  2,
				numbers: []int{1, 2, 3, 4, 0},
			},
			want: true,
		},
		{
			name: "Test 4",
			args: args{
				number:  10,
				numbers: []int{5, 4, 3, 2, 1},
			},
			want: false,
		},
		{
			name: "Test 5",
			args: args{
				number:  10,
				numbers: []int{5, 5, 3, 2, 1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _20220312(tt.args.number, tt.args.numbers); got != tt.want {
				t.Errorf("_20220312() = %v, want %v", got, tt.want)
			}
		})
	}
}
