package c2s

import "github.com/youpy/go-wav"

func Cplx2spl(in []complex128) []wav.Sample {
	out := make([]wav.Sample, 0, len(in))
	for i := range out {
		value := real(in[i])
		out[i].Values[0] = int(value)
		out[i].Values[0] = int(value)
	}
	return out
}
