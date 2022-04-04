package hugipipes_sample

type MonoSpectrum struct {
	Points []MonoSpectrumPoint
}

func newMonoSpectrum(Points []MonoSpectrumPoint) *MonoSpectrum {
	return &MonoSpectrum{
		Points: Points,
	}
}
