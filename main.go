package main

import (
	"log"
	"math"
	"os"

	"github.com/mjibson/go-dsp/fft"
	cmplxs "github.com/oory33/ConvRev_go/c2s"
	flts "github.com/oory33/ConvRev_go/f2s"

	"github.com/youpy/go-wav"
)

func main() {
	overlap := 0.1

	var datareader *wav.Reader
	var irreader *wav.Reader

	ir, err1 := os.Open("./input/ir.wav")
	if err1 != nil {
		log.Fatal(err1)
	} else {
		irreader = wav.NewReader(ir)
	}
	defer ir.Close()

	data, err2 := os.Open("./input/test.wav")
	if err2 != nil {
		log.Fatal(err2)
	} else {
		datareader = wav.NewReader(data)
	}

	irl, _ := irreader.Datasize()
	datal, _ := datareader.Datasize()
	srate, _ := irreader.Samprate()
	bit, _ := irreader.BitsPerSample()
	longer := math.Max(float64(irl), float64(datal))
	shorter := math.Min(float64(irl), float64(datal))

	if longer == float64(irl) {
		panic("IR data is longer")
	} //IRの長さがデータの長さより長い場合は暫定的にエラー

	samples, _ := irreader.ReadSamples(uint32(shorter))
	irsample := make([]float64, uint(shorter))
	for i, sample := range samples {
		irsample[i] = float64(irreader.IntValue(sample, 0))
	}

	samples, _ = datareader.ReadSamples(uint32(longer))
	datasample := make([]float64, uint(longer))
	for i, sample := range samples {
		datasample[i] = float64(datareader.IntValue(sample, 0))
	}

	ircmplx := flts.Cmplx(irsample)
	out := make([]complex128, int(longer))

	for i := 0; i < int(longer-shorter); i += int(shorter * overlap) {
		datacmplx := flts.Cmplx(datasample[i : i+int(shorter)])
		out2 := fft.Convolve(ircmplx, datacmplx)

		out2 = append(make([]complex128, i), out2...)
		out2 = append(out2, make([]complex128, int(longer)-i-int(shorter))...)
		out = cmplxs.Add(out, out2)
	}

	output := cmplxs.Spl(out, bit+3)

	outfile, _ := os.Create("./output/out.wav")
	defer outfile.Close()

	writer := wav.NewWriter(outfile, uint32(longer), 2, srate, bit)
	writer.WriteSamples(output)
}
