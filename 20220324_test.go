package dailygo

import "testing"

func Test_20220324(t *testing.T) {
	type args struct {
		str string
		k   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{
				str: "abcba",
				k:   2,
			},
			want: "bcb",
		},
		{
			name: "Test 2",
			args: args{
				str: "abcbadddddaaassccbbbddceed",
				k:   3,
			},
			want: "adddddaaass",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _20220324(tt.args.str, tt.args.k); got != tt.want {
				t.Errorf("_20220324() = %v, want %v", got, tt.want)
			}
		})
	}
}
