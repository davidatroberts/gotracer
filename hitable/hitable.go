package hitable

import (
	"GoTracer/gtmath"
	"fmt"
)

// HitRecord stores record of a hit
type HitRecord struct {
	t         float64
	P, Normal gtmath.Vector
}

// Hitable interface for things that can be hit
type Hitable interface {
	Hit(ray gtmath.Ray, tMin, tMax float64, rec *HitRecord) bool
}

// List list of things that can be hit
type List struct {
	List []Hitable
}

// Hit iterates through list and passes through
func (l *List) Hit(ray gtmath.Ray, tMin, tMax float64, rec *HitRecord) bool {
	// var tmpRecord HitRecord
	tmpRecord := &HitRecord{}
	hitAnything := false
	closestSoFar := tMax
	for _, h := range l.List {
		if h.Hit(ray, tMin, closestSoFar, tmpRecord) {
			hitAnything = true
			closestSoFar = tmpRecord.t
			rec = tmpRecord
			fmt.Printf("%v\n", rec.Normal)

		}
	}

	return hitAnything
}
