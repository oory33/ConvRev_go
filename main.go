package main

import (
	"log"
	"math"
	"os"

	"github.com/oory33/ConvRev_go/c2s"
	"github.com/youpy/go-wav"
	"gonum.org/v1/gonum/dsp/window"
)

func main() {
	overlap := 1024

	var datareader *wav.Reader
	var irreader *wav.Reader

	ir, err1 := os.Open("./input/ir_1.wav")
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
	longer := math.Max(float64(irl), float64(datal))
	shorter := math.Min(float64(irl), float64(datal))

	if longer == float64(irl) {
		log.Fatal()
	} //IRの長さがデータの長さより長い場合は暫定的にエラー

	samples, _ := irreader.ReadSamples(uint32(shorter))
	irsample := make([]complex128, uint(shorter))
	for i, sample := range samples {
		irsample[i] = complex(float64(irreader.IntValue(sample, 0)), 0)
	}

	samples, _ = datareader.ReadSamples(uint32(longer))
	datasample := make([]complex128, uint(longer))
	for i, sample := range samples {
		datasample[i] = complex(float64(datareader.IntValue(sample, 0)), 0)
	}

	outcmplx := make([]complex128, uint32(longer))
	outcmplx2 := make([]complex128, uint32(longer))

	for i := 0; i < int(longer)-int(shorter); i += int(shorter) * int(overlap) {
		windowedir := window.HammingComplex(irsample[i:int(shorter)])
		windoweddata := window.HammingComplex(datasample[i:int(shorter)])

		outcmplx = windoweddata + windowedir

	}

	output := c2s.Cplx2spl(outcmplx)

	outfile, _ := os.Create("./output/out.wav")
	defer outfile.Close()

	srate, _ := irreader.Samprate()
	bit, _ := irreader.BitsPerSample()

	writer := wav.NewWriter(outfile, uint32(longer), 2, srate, bit)
	writer.WriteSamples(output)
}
