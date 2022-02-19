package util

import (
    "math/rand"
    "math"
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

/*inline double degrees_to_radians(double degrees) {
    return degrees * pi / 180.0;
}*/

func DegreesToRadians(degrees float64) float64 {
    return degrees * math.Pi / 180;
}
