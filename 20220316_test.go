package dailygo

import "testing"

func Test_20220316(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "case 1",
			args: args{
				a: 3,
				b: 4,
			},
			want:  3,
			want1: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := _20220316(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("_20220316() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("_20220316() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
