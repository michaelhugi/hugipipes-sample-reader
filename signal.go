package hugipipes_sample

import (
	"errors"
	"fmt"
	fft2 "github.com/mjibson/go-dsp/fft"
	"github.com/mjibson/go-dsp/spectral"
	"github.com/mjibson/go-dsp/wav"
	"math/cmplx"
	"os"
)

type Channel uint8

const (
	Left  Channel = 0
	Right Channel = 1
)

type Signal struct {
	samplesL    []float64
	samplesR    []float64
	SampleRate  float64
	Wav         *wav.Wav
	SampleCount float64
}

func (s *Signal) GetSamples(c Channel) []float64 {
	if c == Left {
		return s.samplesL
	}
	return s.samplesR
}

func NewSignal(wavPath string) (s *Signal, e error) {
	f, err := os.Open(wavPath)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = f.Close()
		if err != nil {
			e = err
		}
	}()
	w, err := wav.New(f)
	if err != nil {
		return nil, err
	}
	if w.NumChannels > 2 || w.NumChannels == 0 {
		return nil, errors.New("only mono or stereo files are supported")
	}
	var samplesL []float64
	var samplesR []float64
	samples, err := w.ReadFloats(w.Samples)
	var SampleCount float64
	if w.NumChannels == 2 {
		samplesL = make([]float64, w.Samples/2)
		samplesR = make([]float64, w.Samples/2)
		SampleCount = float64(w.Samples / 2)
		if err != nil {
			return nil, err
		}
		for i := range samples {
			if i%2 == 0 {
				samplesL[i/2] = float64(samples[i])
			} else {
				samplesR[i/2] = float64(samples[i])
			}
		}
	} else {
		samplesL = make([]float64, w.Samples)
		samplesR = make([]float64, w.Samples)
		SampleCount = float64(w.Samples)
		for i := range samples {
			samplesL[i] = float64(samples[i])
			samplesR[i] = float64(samples[i])
		}
	}
	minL, maxL := calcMinAndMax(samplesL)
	minR, maxR := calcMinAndMax(samplesR)
	return &Signal{
		SampleRate:  float64(w.SampleRate),
		Wav:         w,
		SampleCount: SampleCount,
		samplesL:    normalize(samplesL, minL, maxL),
		samplesR:    normalize(samplesR, minR, maxR),
	}, nil
}

func calcMinAndMax(samples []float64) (float64, float64) {
	min := samples[0]
	max := samples[0]
	for _, s := range samples {
		if min > s {
			min = s
		}
		if max < s {
			max = s
		}
	}
	return min, max
}

func normalize(input []float64, min float64, max float64) []float64 {

	// We can't normalize a flat signal where min == max
	if min == max {
		panic("cannot normalize signal. Max == Min")
	}

	normalized := make([]float64, len(input))

	for i, val := range input {
		normalized[i] = 2*((val-min)/(max-min)) - 1
	}
	return normalized
}

func (s *Signal) String() string {
	return fmt.Sprintf("------------------------------\n"+
		"Sample:\n"+
		"SampleRate:  %f Hz\n"+
		"SampleCount: %f\n"+
		"SamplesL:    %d\n"+
		"SamplesR:    %d\n"+
		"Channels:    %d\n"+
		"------------------------------\n",
		s.SampleRate,
		s.SampleCount,
		len(s.samplesL),
		len(s.samplesR),
		s.Wav.NumChannels)
}

func (s *Signal) pwelchOptions() *spectral.PwelchOptions {
	return &spectral.PwelchOptions{
		NFFT: s.Wav.Samples,
	}
}

func (s *Signal) MonoSpectrum(c Channel) (*MonoSpectrum, error) {
	samples := s.GetSamples(c)
	pxx, freqs := spectral.Pwelch(samples, s.SampleRate, s.pwelchOptions())

	spectrumPoints := make([]MonoSpectrumPoint, len(pxx))

	spectrum2 := fft2.FFTReal(samples)

	spectrum1 := spectrum2[len(spectrum2)/2-1 : len(spectrum2)]

	if len(spectrum1) != len(pxx) {
		panic(fmt.Sprintf("should have same length!\n  %d\n  %d", len(spectrum1), len(pxx)))
	}

	for i, p := range pxx {
		abs := cmplx.Abs(spectrum1[i]) * 2
		spectrumPoints[i] = *newMonoSpectrumPoint(p, freqs[i], abs)
	}

	return newMonoSpectrum(spectrumPoints), nil
}
