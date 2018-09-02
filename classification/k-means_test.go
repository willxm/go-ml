package classification

import (
	"testing"

	"github.com/willxm/go-ml/kit"
)

func TestKMeans(t *testing.T) {
	type args struct {
		points []kit.Point
		k      int
		iter   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				points: []kit.Point{
					{0.0, 0.0},
					{1.0, 2.0},
					{3.0, 1.0},
					{8.0, 8.0},
					{9.0, 10.0},
					{10.0, 7.0},
				},
				k:    2,
				iter: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			KMeans(tt.args.points, tt.args.k, tt.args.iter)
		})
	}
}
