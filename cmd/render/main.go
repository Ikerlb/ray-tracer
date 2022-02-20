package main

import (
    "fmt"
    "math"
    "math/rand"
    "time"
    "os"
    "github.com/pkg/profile"

    sphere "github.com/ikerlb/ray-tracer/cmd/sphere"
    lamb "github.com/ikerlb/ray-tracer/cmd/lambertian"
    metal "github.com/ikerlb/ray-tracer/cmd/metal"
    dielectric "github.com/ikerlb/ray-tracer/cmd/dielectric"

    ray "github.com/ikerlb/ray-tracer/pkg/ray"
    util "github.com/ikerlb/ray-tracer/pkg/util"
    vec "github.com/ikerlb/ray-tracer/pkg/vec"
    model "github.com/ikerlb/ray-tracer/pkg/model"
    world "github.com/ikerlb/ray-tracer/pkg/world"
    cam "github.com/ikerlb/ray-tracer/pkg/camera"
)

func randomScene(rng *rand.Rand) world.World {

    sphereList := make([]model.Hittable, 0)

    groundMat := lamb.Lambertian{vec.Vector{0.5, 0.5, 0.5}}
    groundSphere := sphere.Sphere{vec.Vector{0, -1000, 0}, 1000, groundMat}

    sphereList = append(sphereList, groundSphere)

    for a := -11; a < 11; a += 1 {
        for b := -11; b < 11; b += 1 {
            chooseMat := rand.Float64()
            x := float64(a) + 0.9 * rand.Float64()
            y := 0.2
            z := float64(b) + 0.9 * rand.Float64()
            center := vec.Vector{x, y, z}

            minus := vec.Minus(center, vec.Vector{4, 0.2, 0})
            if minus.Length() > 0.9 {

                if chooseMat < 0.9 {
                    fmt.Fprintf(os.Stderr, "%f: Diffuse! \n", chooseMat)
                    // Diffuse
                    albedo :=  vec.Random(rng, 0, 1.0)
                    sphereMat := lamb.Lambertian{albedo}
                    sph := sphere.Sphere{center, 0.2, sphereMat}
                    sphereList = append(sphereList, sph)
                } else if chooseMat < 0.95 {
                    // metal
                    fmt.Fprintf(os.Stderr, "%f: Metal! \n", chooseMat)
                    albedo := vec.Random(rng, 0, 1.0)
                    fuzz := util.RandomRange(rng, 0.5, 1)
                    sphereMat := metal.Metal{albedo, fuzz}
                    sph := sphere.Sphere{center, 0.2, sphereMat}
                    sphereList = append(sphereList, sph)
                } else {
                    fmt.Fprintf(os.Stderr, "%f: Dielectric! \n", chooseMat)
                    sphereMat := dielectric.Dielectric{1.5}
                    sph := sphere.Sphere{center, 0.2, sphereMat}
                    sphereList = append(sphereList, sph)
                }
            }
        }
    }

    material1 := dielectric.Dielectric{1.5}
    sphere1 := sphere.Sphere{vec.Vector{0, 1, 0}, 1.0, material1}

    material2 := lamb.Lambertian{vec.Vector{0.4, 0.2, 0.1}}
    sphere2 := sphere.Sphere{vec.Vector{-4, 1, 0}, 1.0, material2}

    material3 := metal.Metal{vec.Vector{0.7, 0.6, 0.5}, 0.0}
    sphere3 := sphere.Sphere{vec.Vector{4, 1, 0}, 1.0, material3}

    sphereList = append(sphereList, sphere1)
    sphereList = append(sphereList, sphere2)
    sphereList = append(sphereList, sphere3)

    return world.World{sphereList}
}

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
    defer profile.Start(profile.ProfilePath("/tmp")).Stop()
    // Image data
    w := 200
    aspectRatio := 3.0 / 2.0
    h := int64(float64(w) / aspectRatio)


    lookFrom := vec.Vector{13, 2, 3}
    lookAt := vec.Vector{0, 0, 0}
    vUp := vec.Vector{0, 1, 0}
    aperture := 0.1
    distToFocus := 10.0
    camera := cam.Init(lookFrom, lookAt, vUp, 20, aspectRatio, aperture, distToFocus)

    numOfSamples := 100
    maxDepth := 50

    randSource := rand.NewSource(time.Now().UnixNano())
    rng := rand.New(randSource)

    world := randomScene(rng)

    fmt.Printf("P3\n%d %d\n255\n", w, h)

    for y := (h - 1); y >= 0; y = y - 1 {
        fmt.Fprintf(os.Stderr, "starting scanline #%d\n", (h - 1) - y)
        for x := 0; x < w; x++ {
            c := vec.Vector{0, 0, 0}
            for sample := 0; sample < numOfSamples; sample += 1 {
                u := (float64(x) + rng.Float64()) / float64(w - 1)
                v := (float64(y) + rng.Float64()) / float64(h - 1)
                r := camera.GetRay(rng, u, v)
                c.AddInPlace(rayColor(rng, &r, world, maxDepth))
            }
            fmt.Printf("%v", c.ToColorString(numOfSamples))
        }
    }
}
