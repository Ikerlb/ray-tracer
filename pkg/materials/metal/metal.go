package metal


import (
    vec "github.com/ikerlb/ray-tracer/pkg/vec"
    ray "github.com/ikerlb/ray-tracer/pkg/ray"
    model "github.com/ikerlb/ray-tracer/pkg/model"
    util "github.com/ikerlb/ray-tracer/pkg/util"

    "math/rand"
)

type Metal struct {
    Albedo vec.Vector;
    Fuzz float64;
}

/*type Material interface {
    Scatter(rIn *ray.Ray, attenuation *vec.Vector) (bool, vec.Vector);
}*/

func (m Metal) Scatter(rIn *ray.Ray, hR *model.HitRecord, rng *rand.Rand) (bool, *vec.Vector, *ray.Ray) {
    fuzz := util.Clamp(m.Fuzz, m.Fuzz, 1.0)

    reflected := vec.Reflect(vec.Unit(rIn.Dir), hR.Normal)

    fuzzed := vec.Scale(vec.RandomInUnitSphere(rng), fuzz)
    scattered := ray.Ray{hR.Point, vec.Add(reflected, fuzzed)}

    if vec.Dot(scattered.Dir, hR.Normal) > 0 {
        return true, &m.Albedo, &scattered
    }
    return false, nil, nil
}
