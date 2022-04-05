package spectrum_drawer

type spectrumDrawerPower struct {
	*spectrumDrawer
	pwrFactor   float64
	maxPwr      float64
	powers      []float64
	yStart      int
	yPlotTop    int
	yPlotBottom int
}

func newSpectrumDrawerPower(spectrumDrawer *spectrumDrawer, powers []float64, yStart int) *spectrumDrawerPower {
	maxPwr := 0.0
	for _, a := range powers {
		if a > maxPwr {
			maxPwr = a
		}
	}

	pwrFactor := float64(spectrumDrawer.plotHeight) / maxPwr

	return &spectrumDrawerPower{
		spectrumDrawer: spectrumDrawer,
		pwrFactor:      pwrFactor,
		powers:         powers,
		yStart:         yStart,
		yPlotTop:       yStart + spectrumDrawer.labelSpace,
		yPlotBottom:    yStart + spectrumDrawer.labelSpace + spectrumDrawer.plotHeight,
	}
}

func (s *spectrumDrawerPower) draw() {
	s.drawXAxis(s.yStart)
	s.drawPlotTitle("Powers", s.yStart)
	s.drawYAxis(s.yStart, nil)
	for i, power := range s.powers {
		s.drawPowerLine(power, s.frequencies[i])
	}
}

func (s *spectrumDrawerPower) pwrToY(pwr float64) int {
	return s.yPlotBottom - int(s.pwrFactor*pwr)
}
func (s *spectrumDrawerPower) drawPowerLine(pwr float64, freq float64) {
	x := s.freqToX(freq)
	for y := s.pwrToY(pwr); y < s.yPlotBottom; y++ {
		s.img.Set(x, y, blue)
	}
}
