package pointservice

import (
	"fmt"
	"testing"

	"github.com/NetworkPy/grpcTask/internal/point"
)

// Так себе тесты, просто, чтобы проверить работу функции с N точками

func TestNoCross(t *testing.T) {
	testArray1 := [][]int64{{100, 100}, {200, 200}, {200, 100}, {100, 200}}
	testArray2 := []*point.Point{{X: 100, Y: 100}, {X: 200, Y: 200}, {X: 200, Y: 100}, {X: 100, Y: 200}}
	noCrossSection(testArray2)
	noCross(testArray1)
	fmt.Println(testArray2)
	fmt.Println(testArray1)
}
