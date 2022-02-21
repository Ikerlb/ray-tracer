package dielectric


import (
    vec "github.com/ikerlb/ray-tracer/pkg/vec"
    ray "github.com/ikerlb/ray-tracer/pkg/ray"
    model "github.com/ikerlb/ray-tracer/pkg/model"
    util "github.com/ikerlb/ray-tracer/pkg/util"

    "math/rand"
    "math"
)

type Dielectric struct {
    IndexOfRefraction float64;
}

/*type Material interface {
    Scatter(rIn *ray.Ray, attenuation *vec.Vector) (bool, vec.Vector);
}*/

func (d Dielectric) Scatter(rIn *ray.Ray, hR *model.HitRecord, rng *rand.Rand) (bool, *vec.Vector, *ray.Ray) {
    attenuation := vec.Vector{1.0, 1.0, 1.0}
    refractionRatio := d.IndexOfRefraction
    if hR.FrontFace  {
        refractionRatio = 1.0 / d.IndexOfRefraction
    }

    unitDirection := vec.Unit(rIn.Dir)
    cosTheta := math.Min(vec.Dot(vec.Neg(unitDirection), hR.Normal), 1.0)
    sinTheta := math.Sqrt(1.0 - (cosTheta * cosTheta))

    var direction vec.Vector

    if refractionRatio * sinTheta > 1.0 || util.Reflectance(cosTheta, refractionRatio) > rng.Float64() {
        direction = vec.Reflect(unitDirection, hR.Normal)
    } else {
        direction = vec.Refract(unitDirection, hR.Normal, refractionRatio)
    }

    scattered := ray.Ray{hR.Point, direction}
    return true, &attenuation, &scattered
}
