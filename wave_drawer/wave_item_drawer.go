package wave_drawer

import "image"

type waveItemDrawer struct {
	*waveDrawer
	top        int
	bottomPlot int
	amplitudes []float64
	maxAmpl    float64
	minAmpl    float64
	amplOffset float64
	amplFactor float64
}

func newWaveItemDrawer(top int, amplitudes []float64, wd *waveDrawer) *waveItemDrawer {
	maxAmpl := wd.maxAmplBase
	minAmpl := wd.minAmplBase
	amplOffset := wd.amplOffsetBase
	amplFactor := wd.amplFactorBase

	if wd.scaleAmplitudePerDraw {
		maxAmpl = -300000.0
		minAmpl = 3000000.0
		for _, amp := range amplitudes {
			if amp > maxAmpl {
				maxAmpl = amp
			}
			if amp < minAmpl {
				minAmpl = amp
			}
		}
		amplRange := maxAmpl - minAmpl
		amplFactor = float64(wd.plotHeight) / amplRange
	}

	amplOffset = -minAmpl
	return &waveItemDrawer{
		waveDrawer: wd,
		top:        top,
		bottomPlot: top + wd.plotHeight + wd.labelSpace,
		amplitudes: amplitudes,
		maxAmpl:    maxAmpl,
		minAmpl:    minAmpl,
		amplOffset: amplOffset,
		amplFactor: amplFactor,
	}
}

func (s *waveItemDrawer) draw() {
	s.drawPlotTitle("WaveSignal", s.top)
	s.drawXAxis(s.top)
	s.drawYAxis(s.top, nil)

	x := s.labelSpace

	for _, amp := range s.amplitudes {
		s.img.Set(x, s.amplToY(amp), image.White)
		x++
	}
}

func (s *waveItemDrawer) amplToY(ampl float64) int {
	return s.bottomPlot - int((ampl+s.amplOffset)*s.amplFactor)
}
