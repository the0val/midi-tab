package melody

import (
	"io"

	"gitlab.com/gomidi/midi/midimessage/channel"
	"gitlab.com/gomidi/midi/smf"
	"gitlab.com/gomidi/midi/smf/smfreader"
)

type Note struct {
	Key      uint8
	Duration uint
}

type Melody []Note

func (m Melody) String() string {
	out := ""
	noteNames := [12]string{
		"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B",
	}
	for _, v := range m {
		note := noteNames[v.Key%12]
		out += note + " "
	}
	// trim last space
	out = out[:len(out)-1]
	return out
}

type melodyReader Melody

func (mr *melodyReader) callback(rd smf.Reader) {
	for {
		m, err := rd.Read()

		if err != nil {
			break
		}

		switch msg := m.(type) {
		case channel.NoteOn:
			*mr = append(*mr, Note{msg.Key(), 1})
		}
	}
}

func ReadReader(midi io.Reader) (Melody, error) {
	var r melodyReader
	smfr := smfreader.New(midi)
	r.callback(smfr)
	return Melody(r), nil
}

func ReadFile(path string) (Melody, error) {
	var r melodyReader
	err := smfreader.ReadFile(path, r.callback)
	if err != nil {
		return nil, err
	}
	return Melody(r), nil
}
