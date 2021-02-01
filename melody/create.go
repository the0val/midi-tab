package melody

import (
	"gitlab.com/gomidi/midi/midimessage/channel"
	"gitlab.com/gomidi/midi/smf"
	"gitlab.com/gomidi/midi/smf/smfreader"
)

type Note struct {
	Key      uint8
	Duration uint
}

type Melody []Note

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

func ReadFile(path string) (Melody, error) {
	var r melodyReader
	err := smfreader.ReadFile(path, r.callback)
	if err != nil {
		return nil, err
	}
	return Melody(r), nil
}
