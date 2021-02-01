package melody

import (
	"gitlab.com/gomidi/midi/midimessage/channel"
	"gitlab.com/gomidi/midi/smf"
)

type Note struct {
	Key      uint8
	Duration uint
}

type Melody []Note

var readFunc = func(rd smf.Reader) {
	var melody Melody

	for {
		m, err := rd.Read()

		if err != nil {
			break
		}

		switch msg := m.(type) {
		case channel.NoteOn:
			melody = append(melody, Note{msg.Key(), 1})
		}
	}
}
