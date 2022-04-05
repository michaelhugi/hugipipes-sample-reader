package hugipipes_sample

import (
	"errors"
	"fmt"
	datatype "github.com/informaticon/lib.go.base.data-types"
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

const amplThreshold = -1

type Signal struct {
	samplesL    []float64
	samplesR    []float64
	SampleRate  float64
	Wav         *wav.Wav
	SampleCount float64
	fftSize     int
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
		fftSize:     calcFFTSize(2, len(samplesL)),
	}, nil
}

func calcFFTSize(cur int, sampleSize int) int {
	next := cur * 2
	if next > sampleSize {
		return cur
	}
	return calcFFTSize(next, sampleSize)
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
	if true {
		return input
	}
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
		"FftSize:     %d\n"+
		"------------------------------\n",
		s.SampleRate,
		s.SampleCount,
		len(s.samplesL),
		len(s.samplesR),
		s.Wav.NumChannels,
		s.fftSize)
}

func (s *Signal) pwelchOptions() *spectral.PwelchOptions {
	return &spectral.PwelchOptions{
		NFFT: s.fftSize,
	}
}

func (s *Signal) MonoSpectrum(c Channel) (*MonoSpectrum, error) {
	samples := s.GetSamples(c)
	samples = samples[0:s.fftSize]
	pxx, freqs := spectral.Pwelch(samples, s.SampleRate, s.pwelchOptions())

	spectrumPoints := make([]MonoSpectrumPoint, len(pxx))

	///twoSidedSpectrum
	spectrum2 := fft2.FFTReal(samples) //fft.FFT(dsputils.ToComplex(samples), s.fftSize) //dftNative(samples) // fft2.FFTReal(samples)

	spectrum1 := spectrum2[0:len(pxx)]

	if len(spectrum1) != len(pxx) {
		panic(fmt.Sprintf("should have same length!\n  %d\n  %d", len(spectrum1), len(pxx)))
	}
	//testFft(spectrum2)
	maxAmpl := 0.0
	for i, p := range pxx {

		amplitudes := cmplx.Abs(spectrum2[i])
		if amplitudes > maxAmpl {
			maxAmpl = amplitudes
		}

		phase := datatype.None[float64]()

		if amplitudes >= amplThreshold {
			phase = datatype.Some(cmplx.Phase(spectrum2[i]))
		}

		spectrumPoints[i] = *newMonoSpectrumPoint(p, freqs[i], amplitudes, phase)
	}
	ampl := make([]float64, len(spectrumPoints))
	freq := make([]float64, len(spectrumPoints))
	for i := range ampl {
		ampl[i] = spectrumPoints[i].Amplitude
		freq[i] = spectrumPoints[i].Frequency
	}

	return newMonoSpectrum(spectrumPoints, maxAmpl), nil
}

func testFft(spectrum []complex128) {
	for i := 0; i < len(spectrum); i++ {

		l1 := cmplx.Abs(spectrum[i])
		r1 := cmplx.Abs(spectrum[len(spectrum)-1-i])
		if l1 != r1 {
			panic(fmt.Sprintf("Should be equal\nl:%v\nr:%v\ni:%d", spectrum[i], spectrum[len(spectrum)-1-i], i))
		}
	}
}
