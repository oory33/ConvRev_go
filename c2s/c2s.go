package cmplxs

import (
	"math"

	"github.com/youpy/go-wav"
)

func Spl(in []complex128, bit uint16) []wav.Sample {
	l := len(in)
	out := make([]wav.Sample, l)
	bits := float64(bit)
	for i := range out {
		value := real(in[i])
		out[i].Values[0] = int(value / math.Pow(2, bits))
		out[i].Values[1] = int(value / math.Pow(2, bits))
	}
	return out
}

func Flt(in []complex128) []float64 {
	l := len(in)
	out := make([]float64, l)
	for i := range out {
		out[i] = real(in[i])
	}
	return out
}

func Add(in1, in2 []complex128) []complex128 {
	l := len(in1)
	out := make([]complex128, l)
	for i := range out {
		out[i] = in1[i] + in2[i]
	}
	return out
}

func Zeros(l int) []complex128 {
	out := make([]complex128, l)
	for i := range out {
		out[i] = 0
	}
	return out
}

func TimesComplx(in1, in2 []complex128) []complex128 {
	l := len(in1)
	out := make([]complex128, l)
	for i := range out {
		out[i] = in1[i] * in2[i]
	}
	return out
}
