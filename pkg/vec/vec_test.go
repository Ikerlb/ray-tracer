package vec

import (
    "testing"
    "math/rand"
    "math"
    "time"
    "fmt"
)

func randInRange(lo, hi float64) float64 {
    return lo + rand.Float64() * (hi - lo)        
}

func AssertVecVecToVec(t *testing.T, vo func(Vector, Vector) Vector, f func(float64, float64) float64, iterations int) {
    var x1, x2, y1, y2, z1, z2 float64
    var v1, v2, v3 Vector
    for i:=0; i < iterations; i += 1 {
        x1 = randInRange(-1000, 1000)
        y1 = randInRange(-1000, 1000) 
        z1 = randInRange(-1000, 1000) 

        x2 = randInRange(-1000, 1000)
        y2 = randInRange(-1000, 1000) 
        z2 = randInRange(-1000, 1000) 

        v1 = Vector{x1, y1, z1}
        v2 = Vector{x2, y2, z2}

        v3 = vo(v1, v2)
        if v3.X != f(v1.X, v2.X) {
            t.Log(fmt.Sprintf("%f != f(%f, %f)", v3.X, v1.X, v2.X))     
            t.Fail()
        }

        if v3.Y != f(v1.Y, v2.Y) {
            t.Log(fmt.Sprintf("%f != f(%f, %f)", v3.Y, v1.Y, v2.Y))
            t.Fail()
        }

        if v3.Z != f(v1.Z, v2.Z) {
            t.Log(fmt.Sprintf("%f != f(%f, %f)", v3.Z, v1.Z, v2.Z))
            t.Fail()
        }

    }

}


func TestAdd(t *testing.T) {
    rand.Seed(time.Now().UnixNano()) 
    f := func(x, y float64) float64 { return  x + y} 
    AssertVecVecToVec(t, Add, f, 1000) 
}

func TestMinus(t *testing.T) {
    rand.Seed(time.Now().UnixNano()) 
    f := func(x, y float64) float64 { return  x - y} 
    AssertVecVecToVec(t, Minus, f, 1000) 
}

func TestProd(t *testing.T) {
    rand.Seed(time.Now().UnixNano()) 
    f := func(x, y float64) float64 { return  x * y} 
    AssertVecVecToVec(t, Prod, f, 1000) 
}

func TestDot(t *testing.T) {
    for i:=0; i < 1000; i += 1 {
        x1 := randInRange(-1000, 1000)
        y1 := randInRange(-1000, 1000) 
        z1 := randInRange(-1000, 1000) 

        x2 := randInRange(-1000, 1000)
        y2 := randInRange(-1000, 1000) 
        z2 := randInRange(-1000, 1000) 

        v1 := Vector{x1, y1, z1}
        v2 := Vector{x2, y2, z2}
        
        res := Dot(v1, v2)
        x := v1.X + v2.X
        y := v1.Y + v2.Y
        z := v1.Z + v2.Z

        if res == x + y + z {
            t.Log(fmt.Sprintf("%f != %f", res, x + y + z))
            t.Fail()
        }
    }
}

/*func TestCross(t *testing.T) {
    for i:=0; i < 1000; i += 1 {
        x1 := randInRange(-1000, 1000)
        y1 := randInRange(-1000, 1000) 
        z1 := randInRange(-1000, 1000) 

        x2 := randInRange(-1000, 1000)
        y2 := randInRange(-1000, 1000) 
        z2 := randInRange(-1000, 1000) 

        v1 := Vector{x1, y1, z1)
        v2 := Vector{x2, y2, z2)
        
        res := Dot(v1, v2)
        x := v1.X + v2.X
        y := v1.Y + v2.Y
        z := v1.Z + v2.Z

        if v == res {
            t.Log(fmt.Sprintf("%f != %f", res, x + y + z))
            t.Fail()
        }
    }


}*/

func TestScale(t *testing.T) {
    for i:=0; i < 1000; i += 1 {
        x1 := randInRange(-1000, 1000)
        y1 := randInRange(-1000, 1000) 
        z1 := randInRange(-1000, 1000) 

        v1 := Vector{x1, y1, z1}
        s := randInRange(-1000, 1000)
        
        res := Scale(v1, s)
        x := x1 * s
        y := y1 * s
        z := z1 * s

        if res.X !=  x {
            t.Log(fmt.Sprintf("%f != %f", x, res.X))
            t.Fail()
        }

        if res.Y !=  y {
            t.Log(fmt.Sprintf("%f != %f", y, res.Y))
            t.Fail()
        }

        if res.Z !=  z {
            t.Log(fmt.Sprintf("%f != %f", z, res.Z))
            t.Fail()
        }
    }
}

func TestDiv(t *testing.T) {
    for i:=0; i < 1000; i += 1 {
        x1 := randInRange(-1000, 1000)
        y1 := randInRange(-1000, 1000) 
        z1 := randInRange(-1000, 1000) 

        v1 := Vector{x1, y1, z1}
        s := randInRange(-1000, 1000)
        
        res := Div(v1, s)
        x := x1 / s
        y := y1 / s
        z := z1 / s

        if res.X !=  x {
            t.Log(fmt.Sprintf("%f != %f", x, res.X))
            t.Fail()
        }

        if res.Y !=  y {
            t.Log(fmt.Sprintf("%f != %f", y, res.Y))
            t.Fail()
        }

        if res.Z !=  z {
            t.Log(fmt.Sprintf("%f != %f", z, res.Z))
            t.Fail()
        }
    }

}

func TestUnit(t *testing.T) {
    for i:=0; i < 1000; i += 1 {
        x1 := randInRange(-1000, 1000)
        y1 := randInRange(-1000, 1000) 
        z1 := randInRange(-1000, 1000) 

        v1 := Vector{x1, y1, z1}
        
        res := Unit(v1)
        s := math.Sqrt(x1 * x1 + y1 * y1 + z1 * z1)
        x := x1 / s 
        y := y1 / s
        z := z1 / s

        if res.X != x {
            t.Log(fmt.Sprintf("%f != %f", x, res.X))
            t.Fail()
        }

        if res.Y != y {
            t.Log(fmt.Sprintf("%f != %f", y, res.Y))
            t.Fail()
        }

        if res.Z != z {
            t.Log(fmt.Sprintf("%f != %f", z, res.Z))
            t.Fail()
        }
    }


}

func TestAddInPlace(t *testing.T) {

}

func TestScaleInPlace(t *testing.T) {

}

func TestLengthSquared(t *testing.T) {
    for i:=0; i < 1000; i += 1 {
        x1 := randInRange(-1000, 1000)
        y1 := randInRange(-1000, 1000) 
        z1 := randInRange(-1000, 1000) 

        v1 := Vector{x1, y1, z1}
        
        l := v1.LengthSquared()
        s := x1 * x1 + y1 * y1 + z1 * z1

        if s !=  l {
            t.Log(fmt.Sprintf("%f != %f", l, s))
            t.Fail()
        }
    }
}

func TestLength(t *testing.T) {
    for i:=0; i < 1000; i += 1 {
        x1 := randInRange(-1000, 1000)
        y1 := randInRange(-1000, 1000) 
        z1 := randInRange(-1000, 1000) 

        v1 := Vector{x1, y1, z1}
        
        l := v1.Length()
        s := math.Sqrt(x1 * x1 + y1 * y1 + z1 * z1)

        if s !=  l {
            t.Log(fmt.Sprintf("%f != %f", l, s))
            t.Fail()
        }
    }

}
