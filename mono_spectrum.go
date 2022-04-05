package hugipipes_sample

import datatype "github.com/informaticon/lib.go.base.data-types"

type MonoSpectrum struct {
	Points  []MonoSpectrumPoint
	MaxAmpl float64
}

func newMonoSpectrum(Points []MonoSpectrumPoint, MaxAmpl float64) *MonoSpectrum {
	return &MonoSpectrum{
		Points:  Points,
		MaxAmpl: MaxAmpl,
	}
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
