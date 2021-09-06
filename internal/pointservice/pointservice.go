package pointservice

import (
	"context"
	"sort"

	"github.com/NetworkPy/grpcTask/internal/point"
)

type PointService struct{}

func NewPointService() *PointService {
	return &PointService{}
}

// CreateGoodPoints...
func (p *PointService) CreateGoodPoints(ctx context.Context, req *point.PointsReq, res *point.PointsRes) error {
	points := req.GetPoints()
	noCrossSection(points)
	res.Points = points
	return nil
}

func noCross(points [][]int64) {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] > points[j][0] {
			return true
		}

		if points[i][0] < points[j][0] {
			return false
		}

		if points[i][1] < points[j][1] {
			return false
		}

		return true
	})
}

// noCrossSection...
func noCrossSection(points []*point.Point) {
	// Отрезки не пересекутся, если мы объединим их
	// по убыванию X.
	// Достаточно отсортировать массив по полю X
	// Если X = X, то сортировать по Y
	sort.Slice(points, func(i, j int) bool {
		if points[i].X > points[j].X {

			return true
		}

		if points[i].X < points[j].X {
			return false
		}

		if points[i].Y < points[j].Y {
			return false
		}

		return true
	})

}
