package camera

import (
    vec "github.com/ikerlb/ray-tracer/pkg/vec"
    ray "github.com/ikerlb/ray-tracer/pkg/ray"
    util "github.com/ikerlb/ray-tracer/pkg/util"

    //"os"
    //"fmt"
    "math"
    "math/rand"
)

type Camera struct {
    origin vec.Vector;
    lowerLeftCorner vec.Vector;
    horizontal vec.Vector;
    vertical vec.Vector;
    u, v, w vec.Vector;
    lensRadius float64;
}


func Init(lookFrom, lookAt, vUp vec.Vector,  vfov, aspectRatio, aperture, focusDist float64) Camera {
    theta := util.DegreesToRadians(vfov)
    h := math.Tan(theta / 2)

    viewportHeight := 2.0 * h
    viewportWidth := aspectRatio * viewportHeight

    w := vec.Unit(vec.Minus(lookFrom, lookAt))
    u := vec.Unit(vec.Cross(vUp, w))
    v := vec.Cross(w, u)

    //fmt.Fprintf(os.Stderr, "w: %v\n", w)
    //fmt.Fprintf(os.Stderr, "u: %v\n", u)
    //fmt.Fprintf(os.Stderr, "v: %v\n", v)

    origin := lookFrom
    horizontal := vec.Scale(u, viewportWidth * focusDist)
    vertical := vec.Scale(v, viewportHeight * focusDist)

    horizontalS := vec.Scale(horizontal, 0.5)
    verticalS := vec.Scale(vertical, 0.5)
    depth := vec.Scale(w, focusDist)
    lowerLeftCorner := vec.Minus(origin, horizontalS, verticalS, depth)

    lensRadius := aperture / 2.0

    return Camera{origin, lowerLeftCorner, horizontal, vertical, u, v, w, lensRadius}
}


func (c Camera) GetRay(rng *rand.Rand, u, v float64) ray.Ray {
    rd := vec.Scale(vec.RandomInUnitDisk(rng), c.lensRadius)
    offset := vec.Add(vec.Scale(c.u, rd.X), vec.Scale(c.v, rd.Y))

    uH := vec.Scale(c.horizontal, u)
    uV := vec.Scale(c.vertical, v)
    s := vec.Add(c.lowerLeftCorner, uH, uV)
    return ray.Ray{vec.Add(c.origin, offset), vec.Minus(s, c.origin, offset)}
}
