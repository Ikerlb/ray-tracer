package camera

import (
    vec "github.com/ikerlb/ray-tracer/pkg/vec"
    ray "github.com/ikerlb/ray-tracer/pkg/ray"
)

type Camera struct {
    origin vec.Vector;
    lowerLeftCorner vec.Vector;
    horizontal vec.Vector;
    vertical vec.Vector;
}

func Init(aspectRatio, viewportHeight, focalLength float64) Camera{

    viewportWidth := aspectRatio * viewportHeight

    origin := vec.Vector{0, 0, 0}
    horizontal := vec.Vector{viewportWidth, 0, 0}
    vertical := vec.Vector{0, viewportHeight, 0}

    horizontalS := vec.Scale(horizontal, 0.5)
    verticalS := vec.Scale(vertical, 0.5)
    depthS := vec.Vector{0, 0, focalLength}
    lowerLeftCorner := vec.Minus(origin, horizontalS, verticalS, depthS)

    return Camera{origin, lowerLeftCorner, horizontal, vertical}
}


func (c Camera) GetRay(u, v float64) ray.Ray{
    uH := vec.Scale(c.horizontal, u)
    uV := vec.Scale(c.vertical, v)
    s := vec.Add(c.lowerLeftCorner, uH, uV)
    return ray.Ray{c.origin, vec.Minus(s, c.origin)}
}
