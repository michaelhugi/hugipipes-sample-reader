package hugipipes_sample

import (
	du "github.com/eripe970/go-dsp-utils"
	"github.com/mjibson/go-dsp/fft"
	"math"
)

type Spectrum struct {
	Points       []SpectrumPoint
	MaxPower     float64
	MaxPowerFreq float64
}

func newSpectrum(realSamples []float64, sampleRate float64) (*Spectrum, error) {

	signal := du.Signal{
		SampleRate: sampleRate,
		Signal:     realSamples,
	}

	normalized, err := signal.Normalize()
	if err != nil {
		return nil, err
	}
	frequencySpectrum, err := normalized.FrequencySpectrum()
	if err != nil {
		return nil, err
	}

	complexSpectrumTwoSided := fft.FFTReal(realSamples)

	maxFrequency := sampleRate / 2

	startIndex := len(complexSpectrumTwoSided) / 2
	validSamplePointsN := len(complexSpectrumTwoSided) - startIndex

	complexSpectrumOneSided := make([]complex128, validSamplePointsN)

	for i := range complexSpectrumOneSided {
		complexSpectrumOneSided[i] = complexSpectrumTwoSided[i+startIndex]
	}

	MaxPower := 0.0
	MaxPowerFreq := 0.0

	deltaF := maxFrequency / float64(validSamplePointsN-1)
	deltaFHalf := deltaF / 2
	points := make([]SpectrumPoint, validSamplePointsN)
	N := float64(len(realSamples))
	Fs := sampleRate
	powFactor := 1 / (Fs * N)
	for i := range points {
		points[i] = *newSpectrumPoint(complexSpectrumOneSided[i], powFactor, &deltaF, &deltaFHalf, &i, &maxFrequency)
		MaxPower = math.Max(MaxPower, points[i].Power)
		if MaxPower == points[i].Power {
			MaxPowerFreq = points[i].CenterFrequency
		}
	}
	return &Spectrum{
		Points:       points,
		MaxPower:     MaxPower,
		MaxPowerFreq: MaxPowerFreq,
	}
}
