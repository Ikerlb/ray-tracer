package main

import (
    "github.com/ikerlb/ray-tracer/pkg/render"
    "runtime"
    "os"
    "fmt"
)

func main(){
    f, err := os.OpenFile("/tmp/spheres.ppm", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Printf("Error opening file %v\n", err)
    } else {
        render.Render(runtime.NumCPU(), f)
        defer f.Close()
    }


}
