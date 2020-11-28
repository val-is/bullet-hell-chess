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
