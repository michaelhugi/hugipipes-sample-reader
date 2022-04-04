package hugipipes_sample

import (
	"time"
)

const fftWindowSize = 512

type TimeSpectrum struct {
	Items []TimeSpectrumItem
}

func newTimeSpectrum(samples []float64, sampleRate float64, duration time.Duration) (*TimeSpectrum, error) {
	items := make([]TimeSpectrumItem, len(samples)-fftWindowSize)

	durationDelta := int64(float64(duration.Nanoseconds()) / float64(len(samples)-1))

	for i := range items {
		subSamples := samples[i : i+fftWindowSize]
		spectrum := newSpectrum(subSamples, sampleRate)
		Time := time.Duration(durationDelta*int64(i)) * time.Nanosecond
		items[i] = *newTimeSpectrumItem(Time, spectrum)
	}
	return &TimeSpectrum{
		Items: items,
	}, nil
}

type TimeSpectrumItem struct {
	Time     time.Duration
	Spectrum *Spectrum
}

func newTimeSpectrumItem(Time time.Duration, Spectrum *Spectrum) *TimeSpectrumItem {
	return &TimeSpectrumItem{
		Time:     Time,
		Spectrum: Spectrum,
	}
}
