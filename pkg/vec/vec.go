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

    scale := 1.0 / float64(numOfSamples)

    r, g, b := v.unpack()
    r = math.Sqrt(r * scale)
    g = math.Sqrt(g * scale)
    b = math.Sqrt(b * scale)

    rs := int(256 * util.Clamp(r, 0.0, 0.999))
    gs := int(256 * util.Clamp(g, 0.0, 0.999))
    bs := int(256 * util.Clamp(b, 0.0, 0.999))

    return fmt.Sprintf("%d %d %d\n", rs, gs, bs)
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

/*vec3 refract(const vec3& uv, const vec3& n, double etai_over_etat) {
    auto cos_theta = fmin(dot(-uv, n), 1.0);
    vec3 r_out_perp =  etai_over_etat * (uv + cos_theta*n);
    vec3 r_out_parallel = -sqrt(fabs(1.0 - r_out_perp.length_squared())) * n;
    return r_out_perp + r_out_parallel;
}*/
func Refract(uv, n Vector, etaiOverEtat float64) Vector {
    cosTheta := math.Min(Dot(Scale(uv, -1), n), 1.0)
    s := Add(uv, Scale(n, cosTheta))
    rOutPrep := Scale(s, etaiOverEtat)
    sq := -math.Sqrt(math.Abs(1.0 - rOutPrep.LengthSquared()))
    rOutParallel := Scale(n, sq)
    return Add(rOutPrep, rOutParallel)
}

/*func (v *Vector) ToColorString() string {
    ri := int(255 * v.X)
    gi := int(255 * v.Y)
    bi := int(255 * v.Z)

    return fmt.Sprintf("%d %d %d\n", ri, gi, bi)
}*/
