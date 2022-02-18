package hittable

import (
    vec "github.com/ikerlb/ray-tracer/pkg/vec"
    ray "github.com/ikerlb/ray-tracer/pkg/ray"
)

type HitRecord struct {
    Point vec.Vector;
    Normal vec.Vector;
    Time float64;
}

type Hittable interface {
    Hit(r *ray.Ray, tMin, tMax float64) (bool, *HitRecord);
}
