package sphere

import (
    vec "github.com/ikerlb/ray-tracer/pkg/vec"
    ray "github.com/ikerlb/ray-tracer/pkg/ray"
    h "github.com/ikerlb/ray-tracer/pkg/hittable"

    "math"
)

type Sphere struct {
    Center vec.Vector;
    Radius float64;
}

/*type HitRecord struct {
    point vec.Vector;
    normal vec.Vector;
    time float64;
}

type Hittable interface {
    hit(r *ray.Ray, tMin, tMax float64) (bool, HitRecord);
}*/

func (s Sphere) Hit(r *ray.Ray, tMin, tMax float64) (bool, *h.HitRecord){
    oc := vec.Minus(r.Origin, s.Center)
    a := r.Dir.LengthSquared()
    halfB := vec.Dot(oc, r.Dir)
    c := oc.LengthSquared() - (s.Radius * s.Radius)
    d := halfB * halfB - (a * c)
    if d < 0.0 {
        return false, nil
    }
    sqrtD := math.Sqrt(d)
    root := (-halfB - sqrtD) / a

    if root < tMin || root > tMax {
        root = (-halfB + sqrtD) / a
        if root < tMin || root > tMax {
            return false, nil
        }
    }

    at := r.At(root)
    normal := vec.Div(vec.Minus(at, s.Center), s.Radius)

    hr := h.HitRecord{Point: at, Normal: normal, Time: root}
    return true, &hr
}
