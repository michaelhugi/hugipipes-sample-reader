package hugipipes_sample

import (
	"fmt"
	"github.com/mjibson/go-dsp/fft"
	wv "github.com/mjibson/go-dsp/wav"
	"math"
	"math/cmplx"
)

type SpectrumPoint struct {
	Point           complex128
	LowFrequency    float64
	HighFrequency   float64
	CenterFrequency float64
	Power           float64
	Ampl            float64
	Phase           float64
}

func newSpectrumPoint(point complex128, powFactor float64, deltaF *float64, deltaFHalf *float64, index *int, maxFrequency *float64) *SpectrumPoint {

	ampl := cmplx.Abs(point)
	phase := cmplx.Phase(point)
	pwr := powFactor * math.Pow(ampl, float64(2))
	CenterFrequency := float64(*index) * *deltaF
	LowFrequency := math.Max(CenterFrequency-*deltaFHalf, 0)
	HighFrequency := math.Min(CenterFrequency+*deltaFHalf, *maxFrequency)

	return &SpectrumPoint{
		Point:           point,
		CenterFrequency: CenterFrequency,
		LowFrequency:    LowFrequency,
		HighFrequency:   HighFrequency,
		Power:           pwr,
		Ampl:            ampl,
		Phase:           phase,
	}
}

func (s *SpectrumPoint) String() string {
	return fmt.Sprintf("------------------------------\n"+
		"complex:           %v\n"+
		"low frequency:     %f Hz\n"+
		"center frequency:  %f Hz\n"+
		"high frequency:    %f Hz\n"+
		"ampl:              %f\n"+
		"phase:             %f Rad\n"+
		"power:             %f dB\n------------------------------\n",
		s.Point,
		s.LowFrequency,
		s.CenterFrequency,
		s.HighFrequency,
		s.Ampl,
		s.Phase,
		s.Power)
}

func spectrum(wav *wv.Wav) ([]SpectrumPoint, error) {
	realSamples, err := wav.ReadFloats(wav.Samples)
	if err != nil {
		return nil, err
	}

	realSamples64 := make([]float64, len(realSamples))
	for i, realSample := range realSamples {
		realSamples64[i] = float64(realSample)
	}

	complexSpectrumTwoSided := fft.FFTReal(realSamples64)

	maxFrequency := float64(wav.SampleRate) / 2

	startIndex := len(complexSpectrumTwoSided) / 2
	validSamplePointsN := len(complexSpectrumTwoSided) - startIndex

	complexSpectrumOneSided := make([]complex128, validSamplePointsN)

	for i := range complexSpectrumOneSided {
		complexSpectrumOneSided[i] = complexSpectrumTwoSided[i+startIndex]
	}

	deltaF := maxFrequency / float64(validSamplePointsN-1)
	deltaFHalf := deltaF / 2
	spec := make([]SpectrumPoint, validSamplePointsN)
	N := float64(len(realSamples))
	Fs := float64(wav.SampleRate)
	powFactor := 1 / (Fs * N)
	for i := range spec {
		spec[i] = *newSpectrumPoint(complexSpectrumOneSided[i], powFactor, &deltaF, &deltaFHalf, &i, &maxFrequency)
	}
	return spec, nil
}
