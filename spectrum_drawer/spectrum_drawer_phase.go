package spectrum_drawer

import (
	datatype "github.com/informaticon/lib.go.base.data-types"
	"math"
)

type spectrumDrawerPhase struct {
	*spectrumDrawer
	phases         []datatype.Option[float64]
	yStart         int
	yPlotTop       int
	yPlotBottom    int
	yCenter        int
	halfPlotHeight float64
	phaseOffset    float64
	phaseFactor    float64
}

func newSpectrumDrawerPhase(spectrumDrawer *spectrumDrawer, phases []datatype.Option[float64], yStart int) *spectrumDrawerPhase {
	yPlotTop := yStart + spectrumDrawer.labelSpace
	yPlotBottom := yStart + spectrumDrawer.labelSpace + spectrumDrawer.plotHeight
	minPhase := 0.0
	maxPhase := 0.0

	for _, p := range phases {
		if p.IsSome() {
			minPhase = math.Min(minPhase, p.Unwrap())
			maxPhase = math.Max(maxPhase, p.Unwrap())
		}
	}
	phaseRange := math.Abs(minPhase) + math.Abs(maxPhase)
	phaseFactor := float64(spectrumDrawer.plotHeight) / phaseRange
	return &spectrumDrawerPhase{
		spectrumDrawer: spectrumDrawer,
		phases:         phases,
		yStart:         yStart,
		yPlotTop:       yPlotTop,
		yPlotBottom:    yPlotBottom,
		yCenter:        (yPlotBottom + yPlotTop) / 2,
		halfPlotHeight: float64(spectrumDrawer.plotHeight / 2),
		phaseOffset:    math.Abs(minPhase),
		phaseFactor:    phaseFactor,
	}
}

func (s *spectrumDrawerPhase) draw() {
	//Labels come from spectrum_drawer_amplitude
	for i, phase := range s.phases {
		if phase.IsSome() {
			s.drawPhasePoint(phase.Unwrap(), s.frequencies[i])
		}
	}
}

func (s *spectrumDrawerPhase) phaseToY(phase float64) int {
	p := phase + s.phaseOffset

	pi := int(s.phaseFactor * p)
	return s.yPlotBottom - pi
}

func (s *spectrumDrawerPhase) drawPhasePoint(phase float64, freq float64) {
	x := s.freqToX(freq)
	s.img.Set(x, s.phaseToY(phase), red)
}
