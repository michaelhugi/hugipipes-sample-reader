package hugipipes_sample

import "fmt"

type MonoSpectrumPoint struct {
	Frequency float64
	Power     float64
	Abs       float64
}

func newMonoSpectrumPoint(Power float64, Frequency float64, Abs float64) *MonoSpectrumPoint {
	return &MonoSpectrumPoint{
		Power:     Power,
		Frequency: Frequency,
		Abs:       Abs,
	}
}

func (s *MonoSpectrumPoint) String() string {
	return fmt.Sprintf("---------------------------\n"+
		"MonoSpectrumPoint:\n"+
		"Frequency: %f Hz\n"+
		"Power:     %f dB\n"+
		"Absolute:  %f\n"+
		"---------------------------\n",
		s.Frequency,
		s.Power,
		s.Abs)
}
