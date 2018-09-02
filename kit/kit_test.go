package kit

import (
	"reflect"
	"testing"
)

func TestEuclideanDistance(t *testing.T) {
	type args struct {
		p1 Point
		p2 Point
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{Point{7.0, 10.0}, Point{4.0, 6.0}},
			want: 5.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EuclideanDistance(tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("EuclideanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxFloat64(t *testing.T) {
	type args struct {
		f []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{f: []float64{1.0, 2.0, 5.0}},
			want: 5.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxFloat64(tt.args.f...); got != tt.want {
				t.Errorf("MaxFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinFloat64(t *testing.T) {
	type args struct {
		f []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{f: []float64{1.0, 2.0, 5.0}},
			want: 1.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinFloat64(tt.args.f...); got != tt.want {
				t.Errorf("MinFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSlice(t *testing.T) {
	type args struct {
		points []Point
	}
	tests := []struct {
		name  string
		args  args
		want  []float64
		want1 []float64
	}{
		// TODO: Add test cases.
		{
			name:  "case1",
			args:  args{[]Point{Point{7.0, 10.0}, Point{4.0, 6.0}, Point{5.0, 11.0}}},
			want:  []float64{7.0, 4.0, 5.0},
			want1: []float64{10.0, 6.0, 11.0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToSlice(tt.args.points)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ToSlice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDataRange(t *testing.T) {
	type args struct {
		xs []float64
		ys []float64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
		want2 float64
		want3 float64
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				xs: []float64{7.0, 4.0, 5.0},
				ys: []float64{10.0, 6.0, 11.0},
			},
			want:  4.0,
			want1: 7.0,
			want2: 6.0,
			want3: 11.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := DataRange(tt.args.xs, tt.args.ys)
			if got != tt.want {
				t.Errorf("DataRange() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("DataRange() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("DataRange() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("DataRange() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}

func TestRandom(t *testing.T) {
	type args struct {
		min float64
		max float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{10.0, 15.0},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Random(tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("Random() = %v, want %v", got, tt.want)
			}
		})
	}
}
