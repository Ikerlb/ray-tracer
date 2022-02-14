package vec

import (
    "math"
    "fmt"
)

type Vector struct {
    X float64
    Y float64
    Z float64
}


func add(v, u Vector) Vector {
    return Vector{v.X + u.X, v.Y + u.Y, v.Z + u.Z}
}

func minus(v, u Vector) Vector {
    return Vector{v.X - u.X, v.Y - u.Y, v.Z - u.Z}
}

func Prod(v, u Vector) Vector {
    return Vector{v.X * u.X, v.Y * u.Y, v.Z * u.Z}
}

func Add(vs ...Vector) Vector {
    res := Vector{0, 0, 0} 
    for _, v := range vs {
        res = add(res, v)
    }
    return res
}

func Minus(vs ...Vector) Vector {
    if len(vs) == 0 {
        return Vector{0, 0, 0}
    } else {
        res := vs[0]
        for i := 1; i < len(vs); i += 1 {
           res = minus(res, vs[i])
        }
        return res
    }
    /*res := Vector{0, 0, 0}
    for _, v := range vs {
        res = minus(res, v)
    }
    return res */
}

func Dot(v, u Vector) float64 {
    return v.X * u.X + v.Y * u.Y + v.Z * u.Z
}

func Cross(v, u Vector) Vector {
    x := v.Y * u.Z - v.Z * u.Y
    y := v.Z * u.X - v.X * u.Z
    z := v.X * u.Y - v.Y * u.X
    return Vector{x, y, z}
}

func Scale(v Vector, s float64) Vector {
    return Vector{v.X * s, v.Y * s, v.Z * s}
}

func Div(v Vector, s float64) Vector {
    return Vector{v.X / s, v.Y / s, v.Z / s}
}

func Unit(v Vector) Vector{
    return Div(v, v.Length())
}

func (v *Vector) AddInPlace(u Vector) {
    v.X += u.X
    v.Y += u.Y
    v.Z += u.Z
}

func (v *Vector) ScaleInPlace(s float64) {
    v.X *= s
    v.Y *= s
    v.Z *= s
}

func (v *Vector) LengthSquared() float64{
    return (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z)
}

func (v * Vector) Length() float64{
    return math.Sqrt(v.LengthSquared())
}

func (v *Vector) ToColorString() string {
    ir := int(255.999 * v.X)
    ig := int(255.999 * v.Y)
    ib := int(255.999 * v.Z)
    return fmt.Sprintf("%d %d %d\n", ir, ig, ib)
}

/*func (v *Vector) dot(u Vector) float64 {
    return (v.x * ux) + (v.y * u.y) + (v.z * u.z)
}*/


