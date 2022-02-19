package main

import (
    "fmt"
    "math"
    "math/rand"
    "time"
    "os"

    sphere "github.com/ikerlb/ray-tracer/cmd/sphere"
    lamb "github.com/ikerlb/ray-tracer/cmd/lambertian"
    metal "github.com/ikerlb/ray-tracer/cmd/metal"

    ray "github.com/ikerlb/ray-tracer/pkg/ray"
    vec "github.com/ikerlb/ray-tracer/pkg/vec"
    model "github.com/ikerlb/ray-tracer/pkg/model"
    world "github.com/ikerlb/ray-tracer/pkg/world"
    cam "github.com/ikerlb/ray-tracer/pkg/camera"
)

func rayColor(rng *rand.Rand, r *ray.Ray, h model.Hittable, depth int) vec.Vector {
    if depth <= 0 {
        return vec.Vector{0, 0, 0}
    }

    doesHit, hRecordP := h.Hit(r, 0.001, math.Inf(1))
    if doesHit {
        doesScatter, attenuation, scatteredRay := hRecordP.Material.Scatter(r, hRecordP, rng)
        if doesScatter {
            return vec.Prod(*attenuation, rayColor(rng, scatteredRay, h, depth-1))
        }
        return vec.Vector{0, 0, 0}
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

    /*ground := lamb.Lambertian{vec.Vector{0.8, 0.8, 0.0}}
    center := metal.Metal{vec.Vector{0.8, 0.8, 0.8}}

    sphere1 := sphere.Sphere{vec.Vector{0, 0, -1}, 0.5, center}
    sphere2 := sphere.Sphere{vec.Vector{0, -100.5, -1}, 100, ground}


    auto material_ground = make_shared<lambertian>(color(0.8, 0.8, 0.0));
    auto material_center = make_shared<lambertian>(color(0.7, 0.3, 0.3));
    auto material_left   = make_shared<metal>(color(0.8, 0.8, 0.8));
    auto material_right  = make_shared<metal>(color(0.8, 0.6, 0.2));

    world.add(make_shared<sphere>(point3( 0.0, -100.5, -1.0), 100.0, material_ground));
    world.add(make_shared<sphere>(point3( 0.0,    0.0, -1.0),   0.5, material_center));
    world.add(make_shared<sphere>(point3(-1.0,    0.0, -1.0),   0.5, material_left));
    world.add(make_shared<sphere>(point3( 1.0,    0.0, -1.0),   0.5, material_right));*/

    materialGround := lamb.Lambertian{vec.Vector{0.8, 0.8, 0.0}}
    materialCenter := lamb.Lambertian{vec.Vector{0.7, 0.3, 0.3}}
    materialLeft := metal.Metal{vec.Vector{0.8, 0.8, 0.8}, 0.3}
    materialRight := metal.Metal{vec.Vector{0.8, 0.6, 0.2}, 1.0}

    sphereGround := sphere.Sphere{vec.Vector{0.0, -100.5, -1.0}, 100.0, materialGround}
    sphereCenter := sphere.Sphere{vec.Vector{0.0, 0.0, -1.0}, 0.5, materialCenter}
    sphereLeft := sphere.Sphere{vec.Vector{-1.0, 0.0, -1.0}, 0.5, materialLeft}
    sphereRight := sphere.Sphere{vec.Vector{1.0, 0.0, -1.0}, 0.5, materialRight}

    sphereList := make([]model.Hittable, 0)
    sphereList = append(sphereList, sphereGround)
    sphereList = append(sphereList, sphereCenter)
    sphereList = append(sphereList, sphereLeft)
    sphereList = append(sphereList, sphereRight)

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
