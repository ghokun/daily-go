package dailygo

import "testing"

func Test_20220318(t *testing.T) {
	type args struct {
		binaryTree *_20220318_node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				binaryTree: &_20220318_node{
					Value: "0",
					Left: &_20220318_node{
						Value: "1",
					},
				},
			},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{
				binaryTree: &_20220318_node{
					Value: "0",
					Left: &_20220318_node{
						Value: "1",
					},
					Right: &_20220318_node{
						Value: "0",
						Left: &_20220318_node{
							Value: "1",
							Left: &_20220318_node{
								Value: "1",
							},
							Right: &_20220318_node{
								Value: "1",
							},
						},
						Right: &_20220318_node{
							Value: "0",
						},
					},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _20220318(tt.args.binaryTree); got != tt.want {
				t.Errorf("_20220318() = %v, want %v", got, tt.want)
			}
		})
	}
}
