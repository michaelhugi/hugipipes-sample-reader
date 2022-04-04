package hugipipes_sample

import (
	wv "github.com/mjibson/go-dsp/wav"
	"os"
)

type Sample struct {
	Wav          *wv.Wav
	Spectrum     *Spectrum
	TimeSpectrum *TimeSpectrum
}

func LoadSample(sample *os.File) (*Sample, error) {
	wav, err := wv.New(sample)
	if err != nil {
		return nil, err
	}
	samples, err := wav.ReadFloats(wav.Samples)
	if err != nil {
		return nil, err
	}
	samples64 := make([]float64, len(samples))
	for i, sample := range samples {
		samples64[i] = float64(sample)
	}
	timeSpectrum, err := newTimeSpectrum(samples64, float64(wav.SampleRate), wav.Duration)
	if err != nil {
		return nil, err
	}
	return &Sample{
		Wav:          wav,
		Spectrum:     newSpectrum(samples64, float64(wav.SampleRate)),
		TimeSpectrum: timeSpectrum,
	}, nil
}
