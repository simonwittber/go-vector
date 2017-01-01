package vector

import "math"

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

func (a Vector3) Magnitude() float32 {
	return float32(math.Sqrt(float64(a.X*a.X + a.Y*a.Y + a.Z*a.Z)))
}

func (a Vector3) SqrMagnitude() float32 {
	return a.X*a.X + a.Y*a.Y + a.Z*a.Z
}

func (a Vector3) Cross(rhs Vector3) Vector3 {
	return Vector3{a.Y*rhs.Z - a.Z*rhs.Y, a.Z*rhs.X - a.X*rhs.Z, a.X*rhs.Y - a.Y*rhs.X}
}

func (a Vector3) Dot(rhs Vector3) float32 {
	return a.X*rhs.X + a.Y*rhs.Y + a.Z*rhs.Z
}

func (a Vector3) Sub(b Vector3) Vector3 {
	return Vector3{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func (a Vector3) Add(b Vector3) Vector3 {
	return Vector3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (a Vector3) Mul(b float32) Vector3 {
	return Vector3{a.X * b, a.Y * b, a.Z * b}
}

func (a Vector3) Div(b float32) Vector3 {
	return Vector3{a.X / b, a.Y / b, a.Z / b}
}

func (a Vector3) Scale(b Vector3) Vector3 {
	return Vector3{a.X * b.X, a.Y * b.Y, a.Z * b.Z}
}

func (a Vector3) Project(onNormal Vector3) Vector3 {
	n := onNormal.Dot(onNormal)
	if n < 1.e-5 {
		return Vector3{0, 0, 0}
	}
	return onNormal.Mul(a.Dot(onNormal)).Div(n)
}

func (a Vector3) ProjectOnPlane(planeNormal Vector3) Vector3 {
	return a.Sub(a.Project(planeNormal))
}

func (a Vector3) Reflect(inNormal Vector3) Vector3 {
	return inNormal.Mul(-2 * inNormal.Dot(a)).Add(a)
}

func (a Vector3) Normalized() Vector3 {
	magnitude := a.Magnitude()
	if magnitude > 1.e-5 {
		return a.Div(magnitude)
	} else {
		return Vector3{0, 0, 0}
	}
}

func (a Vector3) Distance(b Vector3) float32 {
	return a.Sub(b).Magnitude()
}

func (a Vector3) Lerp(b Vector3, d float32) Vector3 {
	return Vector3{a.X + (b.X-a.X)*d, a.Y + (b.Y-a.Y)*d, a.Z + (b.Z-a.Z)*d}
}
