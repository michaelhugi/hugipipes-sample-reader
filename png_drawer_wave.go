package hugipipes_sample

import (
	dt "github.com/informaticon/lib.go.base.data-types"
	"image/color"
	"time"
)

type WavePngDrawer struct {
	startTime dt.Option[time.Duration]
	endTime   dt.Option[time.Duration]
	waves     []WavePngDrawerItem
	title     string
}

func (s *WavePngDrawer) AddWave(item *WavePngDrawerItem) *WavePngDrawer {
	s.waves = append(s.waves, *item)
	return s
}

func (s *WavePngDrawer) StartTime(startTime time.Duration) *WavePngDrawer {
	s.startTime = dt.Some(startTime)
	return s
}

func (s *WavePngDrawer) EndTime(endTime time.Duration) *WavePngDrawer {
	s.endTime = dt.Some(endTime)
	return s
}

func (s *WavePngDrawer) getTimes() []time.Duration {
	if len(s.waves) == 0 {
		return make([]time.Duration, 0)
	}
	maxI := 0
	maxSamples := 0
	for i, w := range s.waves {
		if len(w.samples) > maxSamples {
			maxSamples = len(w.samples)
			maxI = i
		}
	}
	return s.waves[maxI].times
}

func NewWavePngDrawer(title string, wve *WavePngDrawerItem) *WavePngDrawer {
	waves := make([]WavePngDrawerItem, 1)
	waves[0] = *wve
	return &WavePngDrawer{
		title: title,
		waves: waves,
	}
}

type WavePngDrawerItem struct {
	samples    []float64
	duration   time.Duration
	sampleRate uint32
	times      []time.Duration
	color      color.Color
}

func NewWavePngDrawerItem(samp *Signal, c AudioChannel, color color.Color) *WavePngDrawerItem {
	return &WavePngDrawerItem{
		samples:    samp.GetSamples(c),
		duration:   samp.Wav.Duration,
		sampleRate: samp.Wav.SampleRate,
		times:      samp.times,
		color:      color,
	}
}
