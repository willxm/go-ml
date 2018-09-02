package classification

import (
	"log"
	"math"

	"github.com/willxm/go-ml/kit"
)

func randomCenter(minX, maxX, minY, maxY float64) (float64, float64) {
	return kit.Random(minX, maxX), kit.Random(minY, maxY)
}

// KMeans ...
func KMeans(points []kit.Point, k, iter int) map[kit.Point][]kit.Point {
	minX, maxX, minY, maxY := kit.DataRange(kit.ToSlice(points))
	//Init k random points
	var randomPoints []kit.Point
	for i := 0; i < k; i++ {
		x := kit.Random(minX, maxX)
		y := kit.Random(minY, maxY)
		randomPoints = append(randomPoints, kit.Point{X: x, Y: y})
	}
	// iteration
	kps := randomPoints
	for i := 0; i < iter; i++ {
		kps = kmeansStep(points, kps)
	}
	res := kmeansClass(points, kps)
	log.Printf("%+v", res)
	return res
}

func kmeansStep(points, kps []kit.Point) []kit.Point {
	kmap := make(map[kit.Point][]kit.Point, len(kps))
	for _, p := range points {
		minDist := float64(math.MaxFloat64)
		point := kit.Point{}
		for _, kp := range kps {
			dist := kit.EuclideanDistance(p, kp)
			if minDist > dist {
				minDist = dist
				point.X = kp.X
				point.Y = kp.Y
			}
		}
		kmap[point] = append(kmap[point], p)
	}
	res := make([]kit.Point, len(kps))
	i := 0
	for _, v := range kmap {
		res[i] = kit.Centroid(v)
		i++
	}
	return res
}

func kmeansClass(points, kps []kit.Point) map[kit.Point][]kit.Point {
	kmap := make(map[kit.Point][]kit.Point, len(kps))
	for _, p := range points {
		minDist := float64(math.MaxFloat64)
		point := kit.Point{}
		for _, kp := range kps {
			dist := kit.EuclideanDistance(p, kp)
			if minDist > dist {
				minDist = dist
				point.X = kp.X
				point.Y = kp.Y
			}
		}
		kmap[point] = append(kmap[point], p)
	}
	return kmap
}
