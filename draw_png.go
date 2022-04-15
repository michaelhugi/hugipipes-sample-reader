package hugipipes_sample

import (
	"fmt"
	mn "github.com/michaelhugi/go-hugipipes-musical-notes"
	sd "github.com/michaelhugi/go-hugipipes-signal-drawer"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
	"time"
)

var green = color.RGBA{A: 255, R: 0, G: 128, B: 0}
var red = color.RGBA{A: 255, R: 128, G: 0, B: 0}
var blue = color.RGBA{A: 255, R: 0, G: 0, B: 128}
var yellow = color.RGBA{A: 255, R: 255, G: 255, B: 0}
var gray = color.RGBA{A: 255, R: 128, G: 128, B: 128}

func DrawSpectrumAndWaves(spectrum *MonoSpectrum, path string, temperament mn.MTemperament, lowNote mn.MNote, highNote mn.MNote, signals []Signal) error {
	drawer := sd.NewDrawer()

	spec1 := sd.NewSpectrumDrawer(drawer, spectrum.getFrequencies(), "Spectrum").Temperament(temperament).StartNote(lowNote).EndNote(highNote)

	spec1.SetItems(sd.NewSpectrumDrawerItems(spectrum.getAmplitudes(), true, blue))
	spec1.SetItems(sd.NewSpectrumDrawerItems(spectrum.getPhases(), false, red))
	spec1.SetMark(sd.NewSpectrumDrawerMark(spectrum.EstimatedBaseFreq.Frequency, green))
	drawer.AddPlot(spec1)
	if signals != nil {
		for i, signal := range signals {
			sig1 := sd.NewWaveDrawer(drawer, signal.GetTimes(), fmt.Sprintf("Wave %d", i+1)).StartTime(time.Millisecond * 1000).EndTime(time.Millisecond * 1100)
			sig1.SetItems(sd.NewWaveDrawerItems(signal.GetSamples(Left), image.White.C))
			drawer.AddPlot(sig1)
		}
	}

	img := image.NewRGBA(image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: drawer.GetWidth(), Y: drawer.GetHeight()}})

	drawer.SetDrawable(sd.NewImageDrawable(img))
	drawer.Build().Draw()

	if !strings.HasSuffix(path, "png") {
		path = path + ".png"
	}
	f, _ := os.Create(path)
	return png.Encode(f, img)
}
