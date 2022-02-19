package world

import (
    ray "github.com/ikerlb/ray-tracer/pkg/ray"
    m "github.com/ikerlb/ray-tracer/pkg/model"
)

type World struct {
   List []m.Hittable
}

func (w World) Hit(r *ray.Ray, tMin, tMax float64) (bool, *m.HitRecord) {
	var res *m.HitRecord
	anyHit := false

	closest := tMax

	for _, h := range w.List {
		if hit, hr := h.Hit(r, tMin, closest); hit {
			anyHit = true
			res = hr
			closest = hr.Time
		}
	}

	return anyHit, res
}
