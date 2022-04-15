package hugipipes_sample

import (
	sd "github.com/michaelhugi/go-hugipipes-signal-drawer"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

var green = color.RGBA{A: 255, R: 0, G: 128, B: 0}
var red = color.RGBA{A: 255, R: 128, G: 0, B: 0}
var blue = color.RGBA{A: 255, R: 0, G: 0, B: 128}

type PngDrawer struct {
	drawer *sd.DrawerBuilder
}

func (s *PngDrawer) AddSpectrum(spec *SpectrumPngDrawer) {
	sp := sd.NewSpectrumDrawer(s.drawer, spec.frequencies, spec.title)
	if spec.temperament.IsSome() {
		sp = sp.Temperament(spec.temperament.Unwrap())
	}
	if spec.lowNote.IsSome() {
		sp = sp.StartNote(spec.lowNote.Unwrap())
	}
	if spec.highNote.IsSome() {
		sp = sp.EndNote(spec.highNote.Unwrap())
	}
	sp.SetItems(sd.NewSpectrumDrawerItems(spec.amplitudes, true, blue))
	sp.SetItems(sd.NewSpectrumDrawerItems(spec.phases, false, red))
	sp.SetMark(sd.NewSpectrumDrawerMark(spec.estimatedBaseFrequency, green))
	s.drawer.AddPlot(sp)
}

func (s *PngDrawer) AddWave(wve *WavePngDrawer) {

	sig := sd.NewWaveDrawer(s.drawer, wve.getTimes(), wve.title)

	if wve.startTime.IsSome() {
		sig = sig.StartTime(wve.startTime.Unwrap())
	}
	if wve.endTime.IsSome() {
		sig = sig.EndTime(wve.endTime.Unwrap())
	}
	for _, w := range wve.waves {
		sig.SetItems(sd.NewWaveDrawerItems(w.samples, w.color))
	}
	s.drawer.AddPlot(sig)
}

func NewPngDrawer() *PngDrawer {
	return &PngDrawer{
		drawer: sd.NewDrawer(),
	}
}

func (s *PngDrawer) PlotWidth(plotWidth int) *PngDrawer {
	s.drawer.PlotWidth(plotWidth)
	return s
}

func (s *PngDrawer) PlotHeight(plotHeight int) *PngDrawer {
	s.drawer.PlotHeight(plotHeight)
	return s
}

func (s *PngDrawer) Draw(path string) error {
	img := image.NewRGBA(image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: s.drawer.GetWidth(), Y: s.drawer.GetHeight()}})

	s.drawer.SetDrawable(sd.NewImageDrawable(img))
	s.drawer.Build().Draw()

	if !strings.HasSuffix(path, "png") {
		path = path + ".png"
	}
	f, _ := os.Create(path)
	return png.Encode(f, img)
}
