package hugipipes_sample

import (
	"github.com/mjibson/go-dsp/fft"
)

type Spectrum struct {
	Points []SpectrumPoint
}

func newSpectrum(realSamples []float64, sampleRate float64) *Spectrum {

	complexSpectrumTwoSided := fft.FFTReal(realSamples)

	maxFrequency := sampleRate / 2

	startIndex := len(complexSpectrumTwoSided) / 2
	validSamplePointsN := len(complexSpectrumTwoSided) - startIndex

	complexSpectrumOneSided := make([]complex128, validSamplePointsN)

	for i := range complexSpectrumOneSided {
		complexSpectrumOneSided[i] = complexSpectrumTwoSided[i+startIndex]
	}

	deltaF := maxFrequency / float64(validSamplePointsN-1)
	deltaFHalf := deltaF / 2
	points := make([]SpectrumPoint, validSamplePointsN)
	N := float64(len(realSamples))
	Fs := sampleRate
	powFactor := 1 / (Fs * N)
	for i := range points {
		points[i] = *newSpectrumPoint(complexSpectrumOneSided[i], powFactor, &deltaF, &deltaFHalf, &i, &maxFrequency)
	}
	return &Spectrum{
		Points: points,
	}
}
