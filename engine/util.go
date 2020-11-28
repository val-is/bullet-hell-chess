package engine

import "fmt"

var (
	actorCount = 0
)

func NewId(basename string) string {
	id := fmt.Sprintf("%s%d", basename, actorCount)
	actorCount++
	return id
}

// x1, y1 w1, h1 describe box, px/py describe point
func CheckBoundingBox(x1, y1, w1, h1, px, py float64) bool {
	return px <= x1+w1 && px >= x1 && py <= y1+h1 && py >= y1
}
