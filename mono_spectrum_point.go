package hugipipes_sample

import "fmt"
import dt "github.com/informaticon/lib.go.base.data-types"

type MonoSpectrumPoint struct {
	Frequency float64
	Power     float64
	Amplitude float64
	Phase     dt.Option[float64]
}

func newMonoSpectrumPoint(Power float64, Frequency float64, Amplitude float64, Phase dt.Option[float64]) *MonoSpectrumPoint {
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
