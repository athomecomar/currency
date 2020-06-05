package currency

type ARS int64

// Float64 converts a ARS to float64
func (a ARS) Float64() float64 {
	x := float64(a)
	x = x / 100
	return x
}

func (a ARS) Multiply(f float64) ARS {
	x := (float64(a) * f) + 0.5
	return ARS(x)
}
