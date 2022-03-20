package dailygo

import "testing"

func Test_20220317(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				message: "111",
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				message: "1111",
			},
			want: 5,
		},
		{
			name: "Test 3",
			args: args{
				message: "11111",
			},
			want: 8,
		},
		{
			name: "Test 4",
			args: args{
				message: "99999",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _20220317(tt.args.message); got != tt.want {
				t.Errorf("_20220317() = %v, want %v", got, tt.want)
			}
		})
	}
}
