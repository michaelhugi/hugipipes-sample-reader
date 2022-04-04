package hugipipes_sample

import (
	du "github.com/eripe970/go-dsp-utils"
	"github.com/mjibson/go-dsp/fft"
)

type Spectrum struct {
	Points []SpectrumPoint
}

func newSpectrum(realSamplesDenormalized []float64, sampleRate float64) (*Spectrum, error) {

	signal := du.Signal{
		SampleRate: sampleRate,
		Signal:     realSamplesDenormalized,
	}

	realSamples, err := signal.Normalize()
	if err != nil {
		return nil, err
	}

	complexSpectrumTwoSided := fft.FFTReal(realSamples.Signal)

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
	N := float64(len(realSamples.Signal))
	Fs := sampleRate
	powFactor := 1 / (Fs * N)
	for i := range points {
		points[i] = *newSpectrumPoint(complexSpectrumOneSided[i], powFactor, &deltaF, &deltaFHalf, &i, &maxFrequency)
	}
	return &Spectrum{
		Points: points,
	}, nil
}
