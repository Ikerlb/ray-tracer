package main

import (
    "fmt"
    "math"
    "math/rand"
    "time"
    "os"

    ray "github.com/ikerlb/ray-tracer/pkg/ray"
    vec "github.com/ikerlb/ray-tracer/pkg/vec"
    sphere "github.com/ikerlb/ray-tracer/pkg/sphere"
    hittable "github.com/ikerlb/ray-tracer/pkg/hittable"
    world "github.com/ikerlb/ray-tracer/pkg/world"
    cam "github.com/ikerlb/ray-tracer/pkg/camera"
)

func rayColor(rng *rand.Rand, r *ray.Ray, h hittable.Hittable, depth int) vec.Vector {
    if depth <= 0 {
        return vec.Vector{0, 0, 0}
    }

    doesHit, hRecordP := h.Hit(r, 0.001, math.Inf(1))
    if doesHit {
        rInUnit := vec.RandomUnitVector(rng)
        target := vec.Add(hRecordP.Point, hRecordP.Normal, rInUnit)
        nRay := ray.Ray{hRecordP.Point, vec.Minus(target, hRecordP.Point)}
        return vec.Scale(rayColor(rng, &nRay, h, depth - 1), 0.5)
    }

    unit := vec.Unit(r.Dir)
    t  := 0.5 * (unit.Y + 1.0)
    c1 := vec.Scale(vec.Vector{1, 1, 1}, 1.0 - t)
    c2 := vec.Scale(vec.Vector{0.5, 0.7, 1.0}, t)
    return vec.Add(c1, c2)
}

func main() {
    // Image data
    w := 400
    aspectRatio := 16.0 / 9.0
    h := int64(float64(w) / aspectRatio)


    // Camera data
    camera := cam.Init(aspectRatio, 2.0, 1.0)

    sphere1 := sphere.Sphere{vec.Vector{0, 0, -1}, 0.5}
    sphere2 := sphere.Sphere{vec.Vector{0, -100.5, -1}, 100}

    sphereList := make([]hittable.Hittable, 0)
    sphereList = append(sphereList, sphere1)
    sphereList = append(sphereList, sphere2)

    world := world.World{sphereList}

    numOfSamples := 100
    maxDepth := 50

    randSource := rand.NewSource(time.Now().UnixNano())
    rng := rand.New(randSource)

    fmt.Printf("P3\n%d %d\n255\n", w, h)

    for y := (h - 1); y >= 0; y = y - 1 {
        fmt.Fprintf(os.Stderr, "starting scanline #%d\n", (h - 1) - y)
        for x := 0; x < w; x++ {
            c := vec.Vector{0, 0, 0}
            for sample := 0; sample < numOfSamples; sample += 1 {
                u := (float64(x) + rng.Float64()) / float64(w - 1)
                v := (float64(y) + rng.Float64()) / float64(h - 1)
                r := camera.GetRay(u, v)
                c.AddInPlace(rayColor(rng, &r, world, maxDepth))
            }
            // fmt.Fprintf(os.Stderr, "ended color as: %v\n", c)

            fmt.Printf("%v", c.ToColorString(numOfSamples))
        }
    }
}
