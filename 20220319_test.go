package dailygo

import "testing"

func Test_20220319(t *testing.T) {
	type args struct {
		binaryTree *_20220319_node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				binaryTree: &_20220319_node{
					Value: "0",
					Left: &_20220319_node{
						Value: "1",
					},
				},
			},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{
				binaryTree: &_20220319_node{
					Value: "0",
					Left: &_20220319_node{
						Value: "1",
					},
					Right: &_20220319_node{
						Value: "0",
						Left: &_20220319_node{
							Value: "1",
							Left: &_20220319_node{
								Value: "1",
							},
							Right: &_20220319_node{
								Value: "1",
							},
						},
						Right: &_20220319_node{
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
			if got := _20220319(tt.args.binaryTree); got != tt.want {
				t.Errorf("_20220319() = %v, want %v", got, tt.want)
			}
		})
	}
}
