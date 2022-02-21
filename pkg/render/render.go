package render

import (
    "fmt"
    "math"
    "math/rand"
    "time"
    "os"
    "io"
    //"github.com/pkg/profile"

    sphere "github.com/ikerlb/ray-tracer/pkg/objects/sphere"

    lamb "github.com/ikerlb/ray-tracer/pkg/materials/lambertian"
    metal "github.com/ikerlb/ray-tracer/pkg/materials/metal"
    dielectric "github.com/ikerlb/ray-tracer/pkg/materials/dielectric"

    ray "github.com/ikerlb/ray-tracer/pkg/ray"
    util "github.com/ikerlb/ray-tracer/pkg/util"
    vec "github.com/ikerlb/ray-tracer/pkg/vec"
    model "github.com/ikerlb/ray-tracer/pkg/model"
    world "github.com/ikerlb/ray-tracer/pkg/world"
    cam "github.com/ikerlb/ray-tracer/pkg/camera"
)

type renderOptions struct {
    width int;
    height int;
    samples int;
    maxDepth int;
    pixels []int;
    cam cam.Camera;
    world world.World;
}

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
                    // Diffuse
                    albedo :=  vec.Random(rng, 0, 1.0)
                    sphereMat := lamb.Lambertian{albedo}
                    sph := sphere.Sphere{center, 0.2, sphereMat}
                    sphereList = append(sphereList, sph)
                } else if chooseMat < 0.95 {
                    // metal
                    albedo := vec.Random(rng, 0, 1.0)
                    fuzz := util.RandomRange(rng, 0.5, 1)
                    sphereMat := metal.Metal{albedo, fuzz}
                    sph := sphere.Sphere{center, 0.2, sphereMat}
                    sphereList = append(sphereList, sph)
                } else {
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

func render(rng *rand.Rand, scanline int, rO renderOptions) {
    for x := 0; x < rO.width; x++ {
        c := vec.Vector{0, 0, 0}
        for sample := 0; sample < rO.samples; sample += 1 {
            u := (float64(x) + rng.Float64()) / float64(rO.width - 1)
            v := (float64(scanline) + rng.Float64()) / float64(rO.height - 1)
            r := rO.cam.GetRay(rng, u, v)
            c.AddInPlace(rayColor(rng, &r, rO.world, rO.maxDepth))
        }
        i := rO.width * scanline + x
        rO.pixels[i] = c.PackToInt(rO.samples)
    }
}

func worker(id int, rng *rand.Rand, jobs <-chan int, results chan<- int, rO renderOptions) {
    for y := range jobs {
        render(rng, y, rO)
        results <- y
    }
}

func Render(cpus int, target io.Writer) {
    //defer profile.Start(profile.ProfilePath("/tmp")).Stop()
    // Image data
    w := 200
    aspectRatio := 3.0 / 2.0
    h := int(float64(w) / aspectRatio)

    lookFrom := vec.Vector{13, 2, 3}
    lookAt := vec.Vector{0, 0, 0}
    vUp := vec.Vector{0, 1, 0}
    aperture := 0.1
    distToFocus := 10.0
    camera := cam.Init(lookFrom, lookAt, vUp, 20, aspectRatio, aperture, distToFocus)

    numOfSamples := 100
    maxDepth := 50


    pixels := make([]int, w * h, w * h)

    randSourceScene := rand.NewSource(time.Now().UnixNano() + int64(cpus))
    rngScene := rand.New(randSourceScene)
    world := randomScene(rngScene)

    rO := renderOptions{w, h, numOfSamples, maxDepth, pixels, camera, world}

    jobs := make(chan int, h)
    results := make(chan int, h)

    for i := 0; i < cpus; i += 1 {
        randSource := rand.NewSource(time.Now().UnixNano() + int64(i))
        rng := rand.New(randSource)
        go worker(i, rng, jobs, results, rO)
    }

    for y := 0; y < h; y += 1 {
        jobs <- y
    }
    close(jobs)

    for y := 0; y < h; y += 1 {
        fmt.Fprintf(os.Stderr, "finished line %d \n", <-results)
    }
    util.EncodeToPpm(pixels, w, h, target)
}
