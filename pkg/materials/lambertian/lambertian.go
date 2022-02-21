package lambertian


import (
    vec "github.com/ikerlb/ray-tracer/pkg/vec"
    ray "github.com/ikerlb/ray-tracer/pkg/ray"
    model "github.com/ikerlb/ray-tracer/pkg/model"

    "math/rand"
)

type Lambertian struct {
    Albedo vec.Vector;
}

/*type Material interface {
    Scatter(rIn *ray.Ray, attenuation *vec.Vector) (bool, vec.Vector);
}*/

func (l Lambertian) Scatter(rIn *ray.Ray, hR *model.HitRecord, rng *rand.Rand) (bool, *vec.Vector, *ray.Ray) {
    // scatter_direction = rec.normal + random_unit_vector();
    scatterDir := vec.Add(hR.Normal, vec.RandomUnitVector(rng))

    if scatterDir.NearZero() {
        scatterDir = hR.Normal
    }

    scatteredRay := ray.Ray{hR.Point, scatterDir}

    return true, &l.Albedo, &scatteredRay
}
