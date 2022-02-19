package model

import (
    vec "github.com/ikerlb/ray-tracer/pkg/vec"
    ray "github.com/ikerlb/ray-tracer/pkg/ray"

    "math/rand"
)

type HitRecord struct {
    Point vec.Vector;
    Normal vec.Vector;
    Time float64;
    Material Material;
}

type Hittable interface {
    Hit(r *ray.Ray, tMin, tMax float64) (bool, *HitRecord);
}

type Material interface {
    Scatter(rIn *ray.Ray, hR *HitRecord, rng *rand.Rand) (bool, *vec.Vector, *ray.Ray);
}
