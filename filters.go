package hugipipes_sample

import (
	"github.com/mattetti/audio"
)
import "github.com/mattetti/audio/transforms/filters"

func (s *Signal) LowPassFilter(cutOffFreq float64, order int) (*Signal, error) {
	if order < 1 {
		order = 1
	}
	bufferL := audio.NewPCMFloatBuffer(s.GetSamples(Left), &audio.Format{SampleRate: int(s.SampleRate)})
	bufferR := audio.NewPCMFloatBuffer(s.GetSamples(Right), &audio.Format{SampleRate: int(s.SampleRate)})
	for i := 0; i < order; i++ {
		if err := filters.LowPass(bufferL, cutOffFreq); err != nil {
			return nil, err
		}

		if err := filters.LowPass(bufferR, cutOffFreq); err != nil {
			return nil, err
		}
	}
	return newSignal(s.SampleRate, s.Wav, float64(len(bufferL.Floats)), bufferL.Floats, bufferR.Floats, s.times), nil
}

func (s *Signal) HighPassFilter(cutOffFreq float64, order int) (*Signal, error) {
	if order < 1 {
		order = 1
	}
	bufferL := audio.NewPCMFloatBuffer(s.GetSamples(Left), &audio.Format{SampleRate: int(s.SampleRate)})
	bufferR := audio.NewPCMFloatBuffer(s.GetSamples(Right), &audio.Format{SampleRate: int(s.SampleRate)})
	for i := 0; i < order; i++ {

		if err := filters.HighPass(bufferL, cutOffFreq); err != nil {
			return nil, err
		}

		if err := filters.HighPass(bufferR, cutOffFreq); err != nil {
			return nil, err
		}
	}
	return newSignal(s.SampleRate, s.Wav, float64(len(bufferL.Floats)), bufferL.Floats, bufferR.Floats, s.times), nil
}

func (s *Signal) BandPassFilter(lower, upper float64, order int) (*Signal, error) {
	signal1, err := s.LowPassFilter(upper, order)

	if err != nil {
		return nil, err
	}

	signal2, err := signal1.HighPassFilter(lower, order)

	if err != nil {
		return nil, err
	}

	return signal2, nil
}
