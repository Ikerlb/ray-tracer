package render

import (
    "fmt"
    "math"
    "math/rand"
    "io"

    ray "github.com/ikerlb/ray-tracer/pkg/ray"
    util "github.com/ikerlb/ray-tracer/pkg/util"
    vec "github.com/ikerlb/ray-tracer/pkg/vec"
    model "github.com/ikerlb/ray-tracer/pkg/model"
    world "github.com/ikerlb/ray-tracer/pkg/world"
    cam "github.com/ikerlb/ray-tracer/pkg/camera"
)

type RayTracerConfig struct {
    Target io.Writer;
    World world.World;
    MaxDepth int;
    Samples int;
    Width int;
    Height int;
    Workers int;
    Rng *rand.Rand;
}

type renderParams struct {
    width int;
    height int;
    samples int;
    maxDepth int;
    pixels []int;
    cam cam.Camera;
    world world.World;
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

func render(rng *rand.Rand, scanline int, rP renderParams) {
    for x := 0; x < rP.width; x++ {
        c := vec.Vector{0, 0, 0}
        for sample := 0; sample < rP.samples; sample += 1 {
            u := (float64(x) + rng.Float64()) / float64(rP.width - 1)
            v := (float64(scanline) + rng.Float64()) / float64(rP.height - 1)
            r := rP.cam.GetRay(rng, u, v)
            c.AddInPlace(rayColor(rng, &r, rP.world, rP.maxDepth))
        }
        i := rP.width * scanline + x
        rP.pixels[i] = c.PackToInt(rP.samples)
    }
}

func worker(id int, rng *rand.Rand, jobs <-chan int, results chan<- int, rP renderParams) {
    for y := range jobs {
        render(rng, y, rP)
        results <- y
    }
}

/*
type renderParams struct {
    width int;
    height int;
    samples int;
    maxDepth int;
    pixels []int;
    cam cam.Camera;
    world world.World;
}
*/
func Render(rtConfig RayTracerConfig) {
    // Image data
    aspectRatio := float64(rtConfig.Width) / float64(rtConfig.Height)

    lookFrom := vec.Vector{13, 2, 3}
    lookAt := vec.Vector{0, 0, 0}
    vUp := vec.Vector{0, 1, 0}
    aperture := 0.1
    distToFocus := 10.0
    camera := cam.Init(lookFrom, lookAt, vUp, 20, aspectRatio, aperture, distToFocus)

    totalPixels := rtConfig.Width * rtConfig.Width

    pixels := make([]int, totalPixels, totalPixels)

    rP := renderParams{
        width: rtConfig.Width,
        height: rtConfig.Height,
        samples: rtConfig.Samples,
        maxDepth: rtConfig.MaxDepth,
        pixels: pixels,
        cam: camera,
        world: rtConfig.World,
    }

    jobs := make(chan int, rtConfig.Height)
    results := make(chan int, rtConfig.Height)

    for i := 0; i < rtConfig.Workers; i += 1 {
        source := rand.NewSource(rtConfig.Rng.Int63())
        rng := rand.New(source)
        go worker(i, rng, jobs, results, rP)
    }

    for y := 0; y < rtConfig.Height; y += 1 {
        jobs <- y
    }
    close(jobs)

    for y := 0; y < rtConfig.Height; y += 1 {
        fmt.Printf("finished line %d \n", <-results)
    }
    util.EncodeToPpm(pixels, rtConfig.Width, rtConfig.Height, rtConfig.Target)
}
