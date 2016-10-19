package poly2tri

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOne(t *testing.T) {
	polyline := []*Point{
		&Point{X: 155702000, Y: 156210000},
		&Point{X: 156718000, Y: 156210000},
		&Point{X: 156718000, Y: 160782000},
		&Point{X: 156464000, Y: 160782000},
		&Point{X: 156464000, Y: 161544000},
		&Point{X: 155956000, Y: 161544000},
		&Point{X: 155956000, Y: 160782000},
		&Point{X: 155702000, Y: 160782000},
	}
	sc := NewSweepContext(polyline)
	tris := sc.Triangulate()
	assert.Equal(t, 6, len(tris))
}
