package dailygo

import "testing"

func Test_20220325(t *testing.T) {
	tests := []struct {
		name string
		want float64
	}{
		{
			name: "Test 1",
			want: 3.141,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _20220325(); got != tt.want {
				t.Errorf("_20220325() = %v, want %v", got, tt.want)
			}
		})
	}
}
