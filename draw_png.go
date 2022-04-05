package hugipipes_sample

import (
	"hugipipes-sample/spectrum_drawer"
)

func DrawSpectrum(spectrum *MonoSpectrum, path string, freqLogarithmic bool) error {
	return spectrum_drawer.DrawSpectrum(freqLogarithmic, path, spectrum.getFrequencies(), spectrum.getAmplitudes(), spectrum.getPhases(), spectrum.getPowers())

}
