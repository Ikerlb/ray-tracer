package ray

import (
    vec "github.com/ikerlb/ray-tracer/pkg/vec"
)

type Ray struct {
    Origin vec.Vector
    Dir    vec.Vector
}

func (r Ray) At(t float64) vec.Vector {
    return vec.Add(r.Origin, vec.Scale(r.Dir, t))
}
