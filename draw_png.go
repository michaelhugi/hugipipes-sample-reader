package hugipipes_sample

import (
	"hugipipes-sample/spectrum_drawer"
)

func DrawSpectrum(spectrum *MonoSpectrum, path string) error {
	return spectrum_drawer.DrawSpectrum(path, spectrum.getFrequencies(), spectrum.getAmplitudes(), spectrum.getPhases(), spectrum.getPowers())

}
