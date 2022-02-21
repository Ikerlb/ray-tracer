package main

import (
    "runtime"
    "os"
    //"github.com/pkg/profile"
    "math/rand"
    "fmt"
    "flag"

    render "github.com/ikerlb/ray-tracer/pkg/render"
    vec "github.com/ikerlb/ray-tracer/pkg/vec"
    util "github.com/ikerlb/ray-tracer/pkg/util"
    model "github.com/ikerlb/ray-tracer/pkg/model"

    sphere "github.com/ikerlb/ray-tracer/pkg/objects/sphere"

    lamb "github.com/ikerlb/ray-tracer/pkg/materials/lambertian"
    metal "github.com/ikerlb/ray-tracer/pkg/materials/metal"
    dielectric "github.com/ikerlb/ray-tracer/pkg/materials/dielectric"

    world "github.com/ikerlb/ray-tracer/pkg/world"
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

func main(){
    //defer profile.Start(profile.ProfilePath("/tmp")).Stop()

    var filePath string;
    flag.StringVar(&filePath, "file", "", "Target path where the image will be saved.")

    w := flag.Int("width", 200, "Width target image will have.")
    h := flag.Int("height", 200, "Height target image will have.")
    samples := flag.Int("samples", 100, "Number of rays per pixels.")
    maxDepth := flag.Int("max-depth", 50, "Number of times a ray is allowed to 'bounce'.")
    seed := flag.Int("seed", 0, "Master seed with which all other seeds will be created.")
    workers := flag.Int("workers", runtime.NumCPU(), "Number of workers to run in parallel.")

    flag.Parse()

    if filePath == "" {
        fmt.Println("No file path was provided.")
        os.Exit(1)
    }

    f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
    defer f.Close()

    if err != nil {
        fmt.Printf("%v", err)
        os.Exit(1)
    } else {
        rngSource := rand.NewSource(int64(*seed))
        rng := rand.New(rngSource)
        world := randomScene(rng)

        config := render.RayTracerConfig{
            Target: f,
            World: world,
            MaxDepth: *maxDepth,
            Samples: *samples,
            Width: *w,
            Height: *h,
            Workers: *workers,
            Rng: rng,
        }
        render.Render(config)
    }
}
