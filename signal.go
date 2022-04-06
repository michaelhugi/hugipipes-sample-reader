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

type AudioChannel uint8

const (
	Left  AudioChannel = 0
	Right AudioChannel = 1
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

func (s *Signal) GetSamples(c AudioChannel) []float64 {
	if c == Left {
		return s.samplesL
	}
	return s.samplesR
}

func NewSignalFromWav(wavPath string) (s *Signal, e error) {
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
	return newSignal(float64(w.SampleRate), w, SampleCount, samplesL, samplesR), nil
}

func newSignal(SampleRate float64, Wav *wav.Wav, SampleCount float64, samplesL []float64, samplesR []float64) *Signal {
	return &Signal{
		SampleRate:  SampleRate,
		Wav:         Wav,
		SampleCount: SampleCount,
		samplesL:    samplesL,
		samplesR:    samplesR,
		fftSize:     calcFFTSize(2, len(samplesL)),
	}
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

func (s *Signal) MonoSpectrum(c AudioChannel) (*MonoSpectrum, error) {
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

		amplitude := cmplx.Abs(spectrum2[i])
		phase := datatype.None[float64]()
		if freqs[i] <= minSpectrumFreq {
			amplitude = 0.0
		} else {
			if amplitude > maxAmpl {
				maxAmpl = amplitude
			}
			if amplitude >= amplThreshold {
				phase = datatype.Some(cmplx.Phase(spectrum2[i]))
			}
		}
		spectrumPoints[i] = *newMonoSpectrumPoint(p, freqs[i], amplitude, phase)
	}

	test := false
	for _, m := range spectrumPoints {
		if m.Amplitude >= maxAmpl {
			test = true
		}
	}
	if !test {
		panic("wtf?")
	}

	return newMonoSpectrum(spectrumPoints, maxAmpl), nil
}

func (s *Signal) BandpassAtBaseFrequency() (*Signal, error) {
	spec, err := s.MonoSpectrum(Left)
	if err != nil {
		return nil, err
	}
	f := spec.EstimatedBaseFreq.Frequency

	return s.BandPassFilter(f-50.0, f+50.0, 6)
}

func (s *Signal) calcExactBaseFrequency() (float64, error) {
	sig, waveCount, err := s.GetPeakToPeakSignal()
	if err != nil {
		return 0, err
	}
	dt := 1 / sig.SampleRate
	dur := dt * sig.SampleCount
	println("Dur ", dur)
	durPerWave := dur / waveCount

	println("DurW ", durPerWave)

	f := 1 / durPerWave
	println(f, "Hz")
	return f, nil

}

func (s *Signal) GetPeakToPeakSignal() (*Signal, float64, error) {

	maxAmpl := 0.0

	for _, a := range s.samplesL {
		if maxAmpl < a {
			maxAmpl = a
		}
	}
	validAmpl := 0.6 * maxAmpl

	firstValidPeak := 0
	curveRising := false
	for i, curr := range s.samplesL {
		if i > 0 && curr > validAmpl {
			prev := s.samplesL[i-1]
			if curveRising {
				if curr < prev {
					firstValidPeak = i - 1
					break
				}
			}
			curveRising = curr > prev
		}
	}

	lastValidPeak := len(s.samplesL) - 2
	curveRising = false
	for i := lastValidPeak; i >= 0; i-- {
		curr := s.samplesL[i]
		if curr > validAmpl {
			prev := s.samplesL[i+1]
			if curveRising {
				if curr <= prev {
					lastValidPeak = i + 1
					break
				}
			}
			curveRising = curr > prev
		}

	}
	samplesL := s.samplesL[firstValidPeak : lastValidPeak+1]
	samplesR := s.samplesR[firstValidPeak : lastValidPeak+1]

	waveCount := 0

	curveRising = false

	for i, curr := range samplesL {
		if i != 0 {
			prev := samplesL[i-1]
			r := curr > prev

			if curveRising && !r {
				waveCount++
			}
			curveRising = r
		}
	}

	waveCount++

	return newSignal(s.SampleRate, s.Wav, float64(len(samplesL)), samplesL, samplesR), float64(waveCount), nil
}
