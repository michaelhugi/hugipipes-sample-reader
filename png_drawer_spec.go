package hugipipes_sample

import (
	dt "github.com/informaticon/lib.go.base.data-types"
	mn "github.com/michaelhugi/go-hugipipes-musical-notes"
)

type SpectrumPngDrawer struct {
	frequencies            []float64
	amplitudes             []float64
	phases                 []float64
	estimatedBaseFrequency float64
	title                  string
	temperament            dt.Option[mn.MTemperament]
	lowNote                dt.Option[mn.MNote]
	highNote               dt.Option[mn.MNote]
}

func NewSpectrumPngDrawer(spec *MonoSpectrum, title string) *SpectrumPngDrawer {
	return &SpectrumPngDrawer{
		title:                  title,
		frequencies:            spec.getFrequencies(),
		amplitudes:             spec.getAmplitudes(),
		phases:                 spec.getPhases(),
		estimatedBaseFrequency: spec.EstimatedBaseFreq.Frequency,
	}
}
func (s *SpectrumPngDrawer) Temperament(t mn.MTemperament) *SpectrumPngDrawer {
	s.temperament = dt.Some(t)
	return s
}
func (s *SpectrumPngDrawer) LowNote(n mn.MNote) *SpectrumPngDrawer {
	s.lowNote = dt.Some(n)
	return s
}
func (s *SpectrumPngDrawer) HighNote(n mn.MNote) *SpectrumPngDrawer {
	s.highNote = dt.Some(n)
	return s
}
