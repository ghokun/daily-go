package dailygo

import "testing"

func Test_20220323(t *testing.T) {
	type args struct {
		steps         int
		possibleSteps []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				steps:         4,
				possibleSteps: []int{1, 2},
			},
			want: 5,
		},
		{
			name: "Test 2",
			args: args{
				steps:         5,
				possibleSteps: []int{1, 2, 3},
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _20220323(tt.args.steps, tt.args.possibleSteps); got != tt.want {
				t.Errorf("_20220323() = %v, want %v", got, tt.want)
			}
		})
	}
}
