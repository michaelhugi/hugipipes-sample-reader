package spectrum_drawer

import (
	"fmt"
	datatype "github.com/informaticon/lib.go.base.data-types"
	mn "github.com/michaelhugi/go-hugipipes-musical-notes"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"strings"
)

var green = color.RGBA{A: 255, R: 0, G: 128, B: 0}
var red = color.RGBA{A: 255, R: 128, G: 0, B: 0}
var blue = color.RGBA{A: 255, R: 0, G: 0, B: 128}
var yellow = color.RGBA{A: 255, R: 255, G: 255, B: 0}
var gray = color.RGBA{A: 255, R: 128, G: 128, B: 128}

const xLogScaleFactor = 600
const freqLeftOffset = 2500

type spectrumDrawer struct {
	plotHeight             int
	plotWidth              int
	labelSpace             int
	imageWidth             int
	imageHeight            int
	spacePart              int
	freqFactor             float64
	maxFrequency           float64
	img                    *image.RGBA
	frequencies            []float64
	freqLogarithmic        bool
	estimatedBaseFrequency float64
}

func newSpectrumDrawer(frequencies []float64, freqLogarithmic bool, estimatedBaseFrequency float64) *spectrumDrawer {
	plotHeight := 300
	plotWidth := len(frequencies)
	if freqLogarithmic {
		plotWidth = int(math.Log(float64(len(frequencies)))*xLogScaleFactor) - freqLeftOffset
	}
	labelSpace := 80
	imageWidth := plotWidth + 2*labelSpace
	imageHeight := 2*plotHeight + 4*labelSpace
	spacePart := labelSpace / 6
	maxFrequency := frequencies[len(frequencies)-1]
	freqFactor := float64(plotWidth) / maxFrequency

	return &spectrumDrawer{
		freqFactor:             freqFactor,
		maxFrequency:           maxFrequency,
		freqLogarithmic:        freqLogarithmic,
		frequencies:            frequencies,
		plotHeight:             plotHeight,
		plotWidth:              plotWidth,
		labelSpace:             labelSpace,
		imageWidth:             imageWidth,
		imageHeight:            imageHeight,
		spacePart:              spacePart,
		estimatedBaseFrequency: estimatedBaseFrequency,
	}
}

func DrawSpectrum(freqLogarithmic bool, path string, frequencies []float64, amplitudes []float64, phases []datatype.Option[float64], powers []float64, estimatedBaseFrequency float64) error {
	drawer := newSpectrumDrawer(frequencies, freqLogarithmic, estimatedBaseFrequency)
	img := image.NewRGBA(image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: drawer.imageWidth, Y: drawer.imageHeight}})
	drawer.img = img
	drawer.drawBackground()

	drawerAmplitude := newSpectrumDrawerAmplitude(drawer, amplitudes, 0)
	drawerAmplitude.draw()

	drawerPhase := newSpectrumDrawerPhase(drawer, phases, 0)
	drawerPhase.draw()

	drawer.drawDivider(2*drawer.labelSpace + drawer.plotHeight)

	drawerPower := newSpectrumDrawerPower(drawer, powers, 2*drawer.labelSpace+drawer.plotHeight)
	drawerPower.draw()

	if !strings.HasSuffix(path, "png") {
		path = path + ".png"
	}
	f, _ := os.Create(path)
	return png.Encode(f, img)
}

func (s *spectrumDrawer) freqToX(freq float64) int {
	if !s.freqLogarithmic {
		return int(s.freqFactor*freq) + s.labelSpace
	}
	return int(math.Log(float64(int(s.freqFactor*freq)+s.labelSpace))*xLogScaleFactor) - freqLeftOffset

}

func (s *spectrumDrawer) drawBackground() {
	for x := 0; x <= s.imageWidth; x++ {
		for y := 0; y <= s.imageHeight; y++ {
			s.img.Set(x, y, image.Black)
		}
	}
}
func (s *spectrumDrawer) drawXAxis(y int) {
	y += s.labelSpace + s.plotHeight
	for x := s.labelSpace - s.spacePart; x <= s.imageWidth; x++ {
		s.img.Set(x, y, image.White)
	}
	temp := mn.NewMTemperamentJust(mn.A, mn.A, mn.Octave4, 440)

	s.drawXAxisOctave(newOctave(temp.Octave(mn.OctaveMinus1)), y)
	s.drawXAxisOctave(newOctave(temp.Octave(mn.Octave0)), y)
	s.drawXAxisOctave(newOctave(temp.Octave(mn.Octave1)), y)
	s.drawXAxisOctave(newOctave(temp.Octave(mn.Octave2)), y)
	s.drawXAxisOctave(newOctave(temp.Octave(mn.Octave3)), y)
	s.drawXAxisOctave(newOctave(temp.Octave(mn.Octave4)), y)
	s.drawXAxisOctave(newOctave(temp.Octave(mn.Octave5)), y)
	s.drawXAxisOctave(newOctave(temp.Octave(mn.Octave6)), y)
	s.drawXAxisOctave(newOctave(temp.Octave(mn.Octave7)), y)
	s.drawXAxisOctave(newOctave(temp.Octave(mn.Octave8)), y)
	s.drawXAxisOctave(newOctave(temp.Octave(mn.Octave9)), y)
}
func (s *spectrumDrawer) drawXAxisOctave(oct Octave, lineTop int) {
	//Draw line
	x1 := s.freqToX(oct.LowerFrequency())
	x2 := s.freqToX(oct.UpperFrequency())
	lineBottom := lineTop + 5*s.spacePart
	for y := lineTop; y <= lineBottom; y++ {
		s.img.Set(x1, y, image.White)
		s.img.Set(x2, y, image.Black)
	}
	notes := oct.getNotes()
	if len(notes) == 0 {
		y := lineTop + 2*s.spacePart
		xCenter := s.freqToX((oct.LowerFrequency()+oct.UpperFrequency())/2) - 3
		s.drawText(xCenter, y, oct.String(), color.White)
	}
	if oct.drawFrequency() {
		xFreq := s.freqToX(oct.LowerFrequency()) + 5
		yFreq := lineTop + s.spacePart*5
		s.drawText(xFreq, yFreq, oct.frequencyReadable(), blue)
	}
	for _, note := range notes {
		s.drawXAxisNote(&note, lineTop)
	}
}

func (s *spectrumDrawer) drawDivider(y int) {
	for x := 0; x <= s.imageWidth; x++ {
		s.img.Set(x, y, gray)
	}
}

func (s *spectrumDrawer) drawXAxisNote(n *Note, lineTop int) {
	x1 := s.freqToX(n.ExactFrequency())
	lineBottom := lineTop + s.spacePart
	for y := lineTop; y <= lineBottom; y++ {
		s.img.Set(x1, y, image.White)
	}
	y := lineBottom + s.spacePart + 3
	x := x1 - 2
	if !strings.Contains(n.String(), "#") {
		s.drawText(x, y, n.String(), image.White)
	}

	y += s.spacePart + 3
	s.drawText(x, y, fmt.Sprintf("%d", n.MidiNoteNumber()), image.White)
}

func (s *spectrumDrawer) drawPlotTitle(title string, lineTop int) {
	x := s.labelSpace
	y := lineTop + 3*s.spacePart
	s.drawText(x, y, title, color.White)
}

func (s *spectrumDrawer) drawText(x int, y int, text string, color color.Color) {
	point := fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)}

	fd := &font.Drawer{
		Dst:  s.img,
		Src:  image.NewUniform(color),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	fd.DrawString(text)
}

func (s *spectrumDrawer) drawBaseFrequency(top int) {
	x := s.freqToX(s.estimatedBaseFrequency)
	bottom := top + s.plotHeight + s.labelSpace
	for y := top + s.labelSpace; y <= bottom; y++ {
		s.img.Set(x, y, yellow)
	}
}

func (s *spectrumDrawer) drawYAxis(top int, labels []xlabel) {
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
