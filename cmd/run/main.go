package main

import (
    "fmt"
    ray "github.com/ikerlb/ray-tracer/pkg/ray"
    vec "github.com/ikerlb/ray-tracer/pkg/vec"
)

func main() {
    // Image data
    w := 400
    aspectRatio := 16.0 / 9.0
    h := int64(400 / aspectRatio)

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

    fmt.Printf("P3\n%d %d\n255\n", w, h)

    for y := (h - 1); y >= 0; y = y - 1 {
        for x := 0; x < w; x++ {
            hU := vec.Scale(horizontalS, float64(x) / float64(w - 1))
            vV := vec.Scale(verticalS, float64(y) / float64(h - 1))

            dir := vec.Minus(vec.Add(lowerLeftCorner, hU, vV), origin)
            r := ray.Ray{origin, dir}

            unit := vec.Unit(r.Dir)
            t  := 0.5 * (unit.Y + 1.0)
            c1 := vec.Scale(vec.Vector{1, 1, 1}, 1.0 - t)
            c2 := vec.Scale(vec.Vector{0.5, 0.7, 1.0}, t)
            c  := vec.Add(c1, c2)

            ir := int(255.999 * c.X)
            ig := int(255.999 * c.Y)
            ib := int(255.999 * c.Z)

            fmt.Printf("%d %d %d\n", ir, ig, ib)
        }
    }
}
