package hugipipes_sample

import (
	"fmt"
	datatype "github.com/informaticon/lib.go.base.data-types"
)

const minSpectrumFreq = 15.0

type MonoSpectrum struct {
	Points            []MonoSpectrumPoint
	MaxAmpl           float64
	EstimatedBaseFreq MonoSpectrumPoint
}

func newMonoSpectrum(Points []MonoSpectrumPoint, MaxAmpl float64) *MonoSpectrum {

	return &MonoSpectrum{
		Points:            Points,
		MaxAmpl:           MaxAmpl,
		EstimatedBaseFreq: calcEstimatedBaseFreq(Points, MaxAmpl),
	}
}

func calcEstimatedBaseFreq(Points []MonoSpectrumPoint, MaxAmpl float64) MonoSpectrumPoint {

	minAmplitudeToCalcEstimatedFreq := MaxAmpl * 0.1
	rising := false
	minAmplReached := false
	for i, curr := range Points[1:] {
		if curr.Frequency > minSpectrumFreq {
			if curr.Amplitude >= minAmplitudeToCalcEstimatedFreq {
				minAmplReached = true
			}
			if minAmplReached {
				prev := Points[i-1]
				if rising {
					if curr.Amplitude < prev.Amplitude {
						return prev
					}
				}
				rising = curr.Amplitude > prev.Amplitude
			}
		}
	}
	panic(fmt.Sprintf("could not find out estimated base frequency! %v", minAmplReached))
}

func (s *MonoSpectrum) getFrequencies() []float64 {
	frequencies := make([]float64, len(s.Points))
	for i, p := range s.Points {
		frequencies[i] = p.Frequency
	}
	return frequencies
}

func (s *MonoSpectrum) getPhases() []datatype.Option[float64] {
	phases := make([]datatype.Option[float64], len(s.Points))
	for i, p := range s.Points {
		phases[i] = p.Phase
	}
	return phases
}

func (s *MonoSpectrum) getAmplitudes() []float64 {
	abs := make([]float64, len(s.Points))
	for i, p := range s.Points {
		abs[i] = p.Amplitude
	}
	return abs
}

func (s *MonoSpectrum) getPowers() []float64 {
	frequencies := make([]float64, len(s.Points))
	for i, p := range s.Points {
		frequencies[i] = p.Frequency
	}
	return frequencies
}
