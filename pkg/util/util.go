package util

import (
    "math/rand"
    "math"
    "fmt"
    "io"
)

func Clamp(x, mn, mx float64) float64 {
    if x < mn {
        return mn
    } else if x > mx {
        return mx
    } else {
        return x
    }
}

func RandomRange(r *rand.Rand, mn, mx float64) float64{
    return mn + (mx - mn) * r.Float64()
}

func Reflectance(cosine, refIdx float64) float64 {
    r0 := (1 - refIdx) / (1 + refIdx)
    r0 = r0 * r0
    return r0 + (1 - r0) * math.Pow((1 - cosine), 5)
}

func DegreesToRadians(degrees float64) float64 {
    return degrees * math.Pi / 180;
}

func EncodeToPpm(pixels []int, w, h int, target io.Writer) {
    fmt.Fprintf(target, "P3\n%d %d\n255\n", w, h)
    rMask := 0xFF << 16
    gMask := 0xFF << 8
    bMask := 0xFF
    for y := (h - 1); y >= 0; y -= 1 {
        for x := 0; x < w; x += 1 {
            i := w * y + x
            r := (pixels[i] & rMask) >> 16
            g := (pixels[i] & gMask) >> 8
            b := (pixels[i] & bMask)
            fmt.Fprintf(target, "%d %d %d\n", r, g, b)
        }
    }
}
