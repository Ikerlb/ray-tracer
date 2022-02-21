package vec

import (
    "math"
    "fmt"
    "math/rand"

    util "github.com/ikerlb/ray-tracer/pkg/util"
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

func prod(v, u Vector) Vector {
    return Vector{v.X * u.X, v.Y * u.Y, v.Z * u.Z}
}

func Prod(vs ...Vector) Vector {
    res := Vector{1, 1, 1}
    for _, v := range vs {
        res = prod(res, v)
    }
    return res
}

func Add(vs ...Vector) Vector {
    res := Vector{0, 0, 0} 
    for _, v := range vs {
        res = add(res, v)
    }
    return res
}

func Minus(vs ...Vector) Vector {
    res := vs[0]
    for i := 1; i < len(vs); i += 1 {
       res = minus(res, vs[i])
    }
    return res
}

func Neg(v Vector) Vector {
    return Scale(v, -1)
}

func Dot(v, u Vector) float64 {
    return v.X * u.X + v.Y * u.Y + v.Z * u.Z
}


/*vec3 cross(const vec3 &u, const vec3 &v) {
    return vec3(u.e[1] * v.e[2] - u.e[2] * v.e[1],
                u.e[2] * v.e[0] - u.e[0] * v.e[2],
                u.e[0] * v.e[1] - u.e[1] * v.e[0]);*/
func Cross(u,v Vector) Vector {
    x := (u.Y * v.Z) - (u.Z * v.Y)
    y := (u.Z * v.X) - (u.X * v.Z)
    z := (u.X * v.Y) - (u.Y * v.X)
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

func (v *Vector) Length() float64{
    return math.Sqrt(v.LengthSquared())
}

func (v Vector) ToColorString(numOfSamples int) string {
    r, g, b := v.unpackColor(numOfSamples)
    return fmt.Sprintf("%d %d %d\n", r, g, b)
}

func (v *Vector) PackToInt(samples int) int {
    r, g, b := v.unpackColor(samples)
    return r << 16 | g << 8 | b
}

// unpacks gamma corrected color
func (v *Vector) unpackColor(samples int) (int, int, int){
    r, b, g := v.unpack()
    s := 1.0 / float64(samples)

    ri := int(256 * util.Clamp(math.Sqrt(r * s), 0.0, 0.999))
    gi := int(256 * util.Clamp(math.Sqrt(g * s), 0.0, 0.999))
    bi := int(256 * util.Clamp(math.Sqrt(b * s), 0.0, 0.999))

    return ri, gi, bi
}

func (v *Vector) unpack() (float64, float64, float64) {
    return v.X, v.Y, v.Z
}

func (v *Vector) NearZero() bool {
    e := 1e-8
    return (math.Abs(v.X)<e) && (math.Abs(v.Y)<e) && (math.Abs(v.Z)<e)
}

// n is a unit vector!
func Reflect(v, n Vector) Vector {
    d := Dot(v, n)
    c := Scale(n, 2 * d)
    return Minus(v, c)
}

func Random(r *rand.Rand, mn, mx float64) Vector {
    x := util.RandomRange(r, mn, mx)
    y := util.RandomRange(r, mn, mx)
    z := util.RandomRange(r, mn, mx)
    return Vector{x, y, z}
}

func RandomInUnitSphere(r *rand.Rand) Vector {
    v := Random(r, 0.0, 1.0)
    for v.LengthSquared() > 1 {
        v = Random(r, 0.0, 1.0)
    }
    return v
}

/*vec3 random_in_unit_disk() {
    while (true) {
        auto p = vec3(random_double(-1,1), random_double(-1,1), 0);
        if (p.length_squared() >= 1) continue;
        return p;
    }
}*/
func RandomInUnitDisk(r *rand.Rand) Vector {
    x := util.RandomRange(r, -1.0, 1.0)
    y := util.RandomRange(r, -1.0, 1.0)
    v := Vector{x, y, 0}
    for v.LengthSquared() >= 1 {
        x = util.RandomRange(r, -1.0, 1.0)
        y = util.RandomRange(r, -1.0, 1.0)
        v = Vector{x, y, 0}
    }
    return v
}


func RandomUnitVector(r *rand.Rand) Vector {
    return Unit(RandomInUnitSphere(r))
}

func Refract(uv, n Vector, etaiOverEtat float64) Vector {
    cosTheta := math.Min(Dot(Scale(uv, -1), n), 1.0)
    s := Add(uv, Scale(n, cosTheta))
    rOutPrep := Scale(s, etaiOverEtat)
    sq := -math.Sqrt(math.Abs(1.0 - rOutPrep.LengthSquared()))
    rOutParallel := Scale(n, sq)
    return Add(rOutPrep, rOutParallel)
}
