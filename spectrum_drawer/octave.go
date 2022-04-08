package spectrum_drawer

import (
	"fmt"
	mn "github.com/michaelhugi/go-hugipipes-musical-notes"
	"math"
)

type Octave struct {
	mn.MOctave
}

func newOctave(o mn.MOctave) Octave {
	return Octave{
		MOctave: o,
	}
}

func (o *Octave) drawFrequency() bool {
	return o.Octave() > 0
}

func (o *Octave) frequencyReadable() string {
	l := fmt.Sprint(math.Round(o.LowerFrequency()*100) / 100)
	if o.Octave() <= 1 {
		return l
	}
	return fmt.Sprintf("%sHz", l)
}

func (o *Octave) getNotes() []Note {
	if o.Octave() < 3 {
		return make([]Note, 0)
	}
	nt := o.AllNotes()

	notes := make([]Note, len(nt))
	for i, n := range nt {
		notes[i] = newNote(n)
	}
	return notes
}

type Note struct {
	mn.MNote
}

func newNote(n mn.MNote) Note {
	return Note{
		MNote: n,
	}
}
