package world

import (
    ray "github.com/ikerlb/ray-tracer/pkg/ray"
    h "github.com/ikerlb/ray-tracer/pkg/hittable"
)

type World struct {
   List []h.Hittable
}

func (w World) Hit(r *ray.Ray, tMin, tMax float64) (bool, *h.HitRecord) {
	var res *h.HitRecord
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
