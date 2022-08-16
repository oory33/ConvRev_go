package flts

import "github.com/youpy/go-wav"

func Add(in1, in2 []float64) []float64 {
	l := len(in1)
	out := make([]float64, l)
	for i := range out {
		out[i] = in1[i] + in2[i]
	}
	return out
}

func Spl(in []float64) []wav.Sample {
	l := len(in)
	out := make([]wav.Sample, l)
	for i := range out {
		value := in[i]
		out[i].Values[0] = int(value)
		out[i].Values[0] = int(value)
	}
	return out
}

func Cmplx(in []float64) []complex128 {
	l := len(in)
	out := make([]complex128, l)
	for i := range out {
		out[i] = complex(in[i], 0)
	}
	return out
}
