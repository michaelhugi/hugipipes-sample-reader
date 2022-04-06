package wave_drawer

import (
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

type waveDrawer struct {
	plotHeight            int
	plotWidth             int
	labelSpace            int
	imageWidth            int
	imageHeight           int
	spacePart             int
	maxAmplBase           float64
	minAmplBase           float64
	amplOffsetBase        float64
	amplFactorBase        float64
	scaleAmplitudePerDraw bool
	img                   *image.RGBA
}

func newWaveDrawer(amplitudes [][]float64, scaleAmplitudePerDraw bool) *waveDrawer {
	plotHeight := 300
	plotWidth := 0
	maxAmpl := -300000.0
	minAmpl := 3000000.0
	for _, a := range amplitudes {
		if len(a) > plotWidth {
			plotWidth = len(a)
		}
		for _, amp := range a {
			if amp > maxAmpl {
				maxAmpl = amp
			}
			if amp < minAmpl {
				minAmpl = amp
			}
		}
	}
	amplRange := maxAmpl - minAmpl
	amplFactor := float64(plotHeight) / amplRange
	labelSpace := 80
	imageWidth := plotWidth + 2*labelSpace
	imageHeight := len(amplitudes)*plotHeight + len(amplitudes)*2*labelSpace
	spacePart := labelSpace / 6

	return &waveDrawer{
		plotHeight:            plotHeight,
		plotWidth:             plotWidth,
		labelSpace:            labelSpace,
		imageWidth:            imageWidth,
		imageHeight:           imageHeight,
		spacePart:             spacePart,
		maxAmplBase:           maxAmpl,
		minAmplBase:           minAmpl,
		amplFactorBase:        amplFactor,
		amplOffsetBase:        -minAmpl,
		scaleAmplitudePerDraw: scaleAmplitudePerDraw,
	}
}

func DrawWaveSignal(path string, amplitudes [][]float64, scaleAmplitudePerDraw bool) error {
	drawer := newWaveDrawer(amplitudes, scaleAmplitudePerDraw)
	img := image.NewRGBA(image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: drawer.imageWidth, Y: drawer.imageHeight}})
	drawer.img = img
	drawer.drawBackground()

	y := 0

	for _, ampl := range amplitudes {
		d := newWaveItemDrawer(y, ampl, drawer)
		d.draw()
		y += 2*drawer.labelSpace + drawer.plotHeight
		drawer.drawDivider(y)
	}

	if !strings.HasSuffix(path, "png") {
		path = path + ".png"
	}
	f, _ := os.Create(path)
	return png.Encode(f, img)
}

func (s *waveDrawer) drawBackground() {
	for x := 0; x <= s.imageWidth; x++ {
		for y := 0; y <= s.imageHeight; y++ {
			s.img.Set(x, y, image.Black)
		}
	}
}
func (s *waveDrawer) drawXAxis(y int) {
	y += s.labelSpace + s.plotHeight
	for x := s.labelSpace - s.spacePart; x <= s.imageWidth; x++ {
		s.img.Set(x, y, image.White)
	}

}

func (s *waveDrawer) drawDivider(y int) {
	for x := 0; x <= s.imageWidth; x++ {
		s.img.Set(x, y, image.White)
	}
}

func (s *waveDrawer) drawPlotTitle(title string, lineTop int) {
	x := s.labelSpace
	y := lineTop + 3*s.spacePart
	s.drawText(x, y, title, color.White)
}

func (s *waveDrawer) drawText(x int, y int, text string, color color.Color) {
	point := fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)}

	fd := &font.Drawer{
		Dst:  s.img,
		Src:  image.NewUniform(color),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	fd.DrawString(text)
}

func (s *waveDrawer) drawYAxis(top int, labels []xlabel) {
	top += s.labelSpace
	bottom := top + s.plotHeight + s.spacePart
	x := s.labelSpace

	for y := top; y <= bottom; y++ {
		s.img.Set(x, y, image.White)
	}
	if labels != nil {
		for _, label := range labels {
			for x := s.labelSpace - s.spacePart; x < s.labelSpace; x++ {
				s.img.Set(x, label.Y+top, image.White)
			}
			s.drawText(s.labelSpace-4*s.spacePart, label.Y+top-3, label.Text, image.White)
		}
	}
}

type xlabel struct {
	Y    int
	Text string
}
