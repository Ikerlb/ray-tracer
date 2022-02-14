package main

import (
    "fmt"
    "math"

    ray "github.com/ikerlb/ray-tracer/pkg/ray"
    vec "github.com/ikerlb/ray-tracer/pkg/vec"
)

func hitSphere(center vec.Vector, radius float64, r ray.Ray) float64 {
    oc := vec.Minus(r.Origin, center)
    a  := r.Dir.LengthSquared()
    halfB  := vec.Dot(oc, r.Dir)
    c  :=  oc.LengthSquared() - (radius * radius)
    d  := halfB * halfB - (a * c)
    if d < 0.0 {
        return -1.0
    } else {
        return (- halfB - math.Sqrt(d)) / a
    }
}

func main() {
    // Image data
    w := 400
    aspectRatio := 16.0 / 9.0
    h := int64(float64(w) / aspectRatio)

    // Camera data
    viewportHeight := 2.0
    viewportWidth := aspectRatio * viewportHeight
    focalLength := 1.0

    origin := vec.Vector{0, 0, 0}
    horizontal := vec.Vector{viewportWidth, 0, 0}
    vertical := vec.Vector{0, viewportHeight, 0}

    horizontalS := vec.Scale(horizontal, 0.5)
    verticalS := vec.Scale(vertical, 0.5)
    depthS := vec.Vector{0, 0, focalLength}
    lowerLeftCorner := vec.Minus(origin, horizontalS, verticalS, depthS)

    sphereCenter := vec.Vector{0, 0, -1}
    //colorRed := vec.Vector{1, 0, 0}

    fmt.Printf("P3\n%d %d\n255\n", w, h)

    for y := (h - 1); y >= 0; y = y - 1 {
        for x := 0; x < w; x++ {
            hU := vec.Scale(horizontal, float64(x) / float64(w - 1))
            vV := vec.Scale(vertical, float64(y) / float64(h - 1))

            dir := vec.Minus(vec.Add(lowerLeftCorner, hU, vV), origin)
            r := ray.Ray{origin, dir}

            t := hitSphere(sphereCenter, 0.5, r)
            //fmt.Printf("!! t es %f", t)
            if t > 0.0 {
                normal := vec.Minus(r.At(t), vec.Vector{0, 0, -1})
                sphereColor := vec.Add(normal, vec.Vector{1, 1, 1})
                sphereColorHalved := vec.Scale(sphereColor, 0.5)
                fmt.Printf(sphereColorHalved.ToColorString())
                //fmt.Printf(colorRed.ToColorString())
                continue
            }

            unit := vec.Unit(r.Dir)
            t  = 0.5 * (unit.Y + 1.0)
            c1 := vec.Scale(vec.Vector{1, 1, 1}, 1.0 - t)
            c2 := vec.Scale(vec.Vector{0.5, 0.7, 1.0}, t)
            c  := vec.Add(c1, c2)

            fmt.Printf(c.ToColorString())
        }
    }
}
