package hugipipes_sample

import "fmt"

type MonoSpectrumPoint struct {
	Frequency float64
	Power     float64
	Amplitude float64
	Phase     float64
}

func newMonoSpectrumPoint(Power float64, Frequency float64, Amplitude float64, Phase float64) *MonoSpectrumPoint {
	return &MonoSpectrumPoint{
		Power:     Power,
		Frequency: Frequency,
		Amplitude: Amplitude,
		Phase:     Phase,
	}
}

func (s *MonoSpectrumPoint) String() string {
	return fmt.Sprintf("---------------------------\n"+
		"MonoSpectrumPoint:\n"+
		"Frequency: %f Hz\n"+
		"Power:     %f dB\n"+
		"Absolute:  %f\n"+
		"Phase:     %v Rad\n"+
		"---------------------------\n",
		s.Frequency,
		s.Power,
		s.Amplitude,
		s.Phase)
}
