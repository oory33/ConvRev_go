package cmplx

import "github.com/youpy/go-wav"

func Cplx2spl(in []complex128) []wav.Sample {
	l := len(in)
	out := make([]wav.Sample, l)
	for i := range out {
		value := real(in[i])
		out[i].Values[0] = int(value)
		out[i].Values[0] = int(value)
	}
	return out
}
