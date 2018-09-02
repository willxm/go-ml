package kit

import (
	"math"
	"math/rand"
	"time"
)

//Point is ...
type Point struct {
	X float64
	Y float64
}

//EuclideanDistance ...
func EuclideanDistance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.X-p2.X, 2) + math.Pow(p1.Y-p2.Y, 2))
}

//ToSlice ...
func ToSlice(points []Point) ([]float64, []float64) {
	var xs, ys []float64
	for _, v := range points {
		xs = append(xs, v.X)
		ys = append(ys, v.Y)
	}
	return xs, ys
}

//DataRange ...
func DataRange(xs, ys []float64) (float64, float64, float64, float64) {
	return MinFloat64(xs...), MaxFloat64(xs...), MinFloat64(ys...), MaxFloat64(ys...)
}

//MaxFloat64 ...
func MaxFloat64(f ...float64) float64 {
	max := -math.MaxFloat64
	for _, v := range f {
		if v > max {
			max = v
		}
	}
	return max
}

//MinFloat64 ...
func MinFloat64(f ...float64) float64 {
	min := math.MaxFloat64
	for _, v := range f {
		if v < min {
			min = v
		}
	}
	return min
}

//Random ...
func Random(min, max float64) float64 {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return r.Float64()*(max-min) + min
}

func Centroid(points []Point) Point {
	l := len(points)
	sx, sy := 0.0, 0.0
	for i := 0; i < l; i++ {
		sx += points[i].X
		sy += points[i].Y
	}
	return Point{
		X: sx / float64(l),
		Y: sy / float64(l),
	}
}
