package main

import (
	"log"
	"os"

	"github.com/oory33/ConvRev_go/c2s"

	"github.com/mjibson/go-dsp/fft"
	"github.com/youpy/go-wav"
)

func main() {
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

	samples, _ := irreader.ReadSamples()
	irsample := make([]complex128, 0, len(samples))
	for _, sample := range samples {
		irsample = append(irsample, complex(irreader.FloatValue(sample, 0), 0))
	}

	samples, _ = datareader.ReadSamples()
	datasample := make([]complex128, 0, len(samples))
	for _, sample := range samples {
		datasample = append(datasample, complex(datareader.FloatValue(sample, 0), 0))
	}

	outcmplx := fft.Convolve(irsample, datasample)
	output := c2s.Cplx2spl(outcmplx)

	outfile, _ := os.Create("./output/out.wav")
	defer outfile.Close()

	srate, _ := irreader.Samprate()
	bit, _ := irreader.BitsPerSample()

	writer := wav.NewWriter(outfile, uint32(len(output)), 2, srate, bit)
	writer.WriteSamples(output)
}
