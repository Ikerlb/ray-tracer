package util

import (
    rand "math/rand"
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
