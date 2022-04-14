package hugipipes_sample

import (
	mn "github.com/michaelhugi/go-hugipipes-musical-notes"
	"hugipipes-sample/spectrum_drawer"
	"hugipipes-sample/wave_drawer"
)

func DrawSpectrum(spectrum *MonoSpectrum, path string, temperament mn.MTemperament, lowNote mn.MNote, highNote mn.MNote) error {
	return spectrum_drawer.DrawSpectrum(path, spectrum.getFrequencies(), spectrum.getAmplitudes(), spectrum.getPhases(), spectrum.getPowers(), spectrum.EstimatedBaseFreq.Frequency, temperament, lowNote, highNote)
}

func DrawWaves(signals []Signal, path string, scaleAmplitudePerDraw bool) error {
	amplitudes := make([][]float64, len(signals))
	for i, s := range signals {
		amplitudes[i] = s.GetSamples(Left)
	}
	return wave_drawer.DrawWaveSignal(path, amplitudes, scaleAmplitudePerDraw)
}
func DrawWave(signal *Signal, path string) error {
	amplitudes := make([][]float64, 1)
	amplitudes[0] = signal.GetSamples(Left)
	return wave_drawer.DrawWaveSignal(path, amplitudes, false)
}
