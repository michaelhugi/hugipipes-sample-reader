package hugipipes_sample

import (
	"math"
	"time"
)

const fftWindowSize = 512

type TimeSpectrum struct {
	MaxPower float64
	Items    []TimeSpectrumItem
}

func newTimeSpectrum(samples []float64, sampleRate float64, duration time.Duration) (*TimeSpectrum, error) {
	items := make([]TimeSpectrumItem, len(samples)-fftWindowSize)

	durationDelta := int64(float64(duration.Nanoseconds()) / float64(len(samples)-1))
	MaxPower := 0.0
	for i := range items {
		subSamples := samples[i : i+fftWindowSize]
		spectrum := newSpectrum(subSamples, sampleRate)
		Time := time.Duration(durationDelta*int64(i)) * time.Nanosecond
		items[i] = *newTimeSpectrumItem(Time, spectrum)
		MaxPower = math.Max(MaxPower, items[i].Spectrum.MaxPower)
	}
	return &TimeSpectrum{
		Items:    items,
		MaxPower: MaxPower,
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
