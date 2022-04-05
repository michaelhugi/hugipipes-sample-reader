package spectrum_drawer

type spectrumDrawerAmplitude struct {
	*spectrumDrawer
	amplFactor  float64
	maxAmpl     float64
	amplitudes  []float64
	yStart      int
	yPlotTop    int
	yPlotBottom int
}

func newSpectrumDrawerAmplitude(spectrumDrawer *spectrumDrawer, amplitudes []float64, yStart int) *spectrumDrawerAmplitude {
	maxAmpl := 0.0
	for _, a := range amplitudes {
		if a > maxAmpl {
			maxAmpl = a
		}
	}

	amplFactor := float64(spectrumDrawer.plotHeight) / maxAmpl

	return &spectrumDrawerAmplitude{
		spectrumDrawer: spectrumDrawer,
		amplFactor:     amplFactor,
		amplitudes:     amplitudes,
		yStart:         yStart,
		yPlotTop:       yStart + spectrumDrawer.labelSpace,
		yPlotBottom:    yStart + spectrumDrawer.labelSpace + spectrumDrawer.plotHeight,
	}
}

func (s *spectrumDrawerAmplitude) draw() {
	s.drawXAxis(s.yStart)
	s.drawPlotTitle("Amplitudes / Phases", s.yStart)
	labels := make([]xlabel, 3)
	labels[0] = xlabel{Y: 0, Text: "PI/2"}
	labels[1] = xlabel{Y: s.plotHeight / 2, Text: "0"}
	labels[2] = xlabel{Y: s.plotHeight, Text: "-PI/2"}
	s.drawYAxis(s.yStart, labels)
	for i, amplitude := range s.amplitudes {
		s.drawAmplitudeLine(amplitude, s.frequencies[i])
	}
}

func (s *spectrumDrawerAmplitude) amplToY(ampl float64) int {
	return s.yPlotBottom - int(s.amplFactor*ampl)
}
func (s *spectrumDrawerAmplitude) drawAmplitudeLine(ampl float64, freq float64) {
	x := s.freqToX(freq)
	for y := s.amplToY(ampl); y < s.yPlotBottom; y++ {
		s.img.Set(x, y, green)
	}
}
