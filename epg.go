package epg

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// EPG holds program guide records
type EPG struct {
	XMLName    xml.Name     `xml:"tv"`
	Channels   []*Channel   `xml:"channel"`
	Programmes []*Programme `xml:"programme"`

	channelsByID map[string]*Channel
}

func (e *EPG) Read(rdr io.Reader) (err error) {
	dec := xml.NewDecoder(rdr)

	if err = dec.Decode(e); err != nil {
		return
	}

	// assign epg to channels and programmes
	for _, ch := range e.Channels {
		ch.epg = e
	}
	for _, pr := range e.Programmes {
		pr.epg = e
	}

	// create id->channel mapping
	e.channelsByID = make(map[string]*Channel)
	for _, ch := range e.Channels {
		e.channelsByID[ch.ID] = ch
	}

	// assign programmes to channels
	for _, ch := range e.Channels {
		for _, pr := range e.Programmes {
			if pr.Channel == ch.ID {
				ch.Programmes = append(ch.Programmes, pr)
			}
		}
	}

	return
}

func (e *EPG) deleteProgramme(pr *Programme) {
	var newArr []*Programme

	for _, p := range e.Programmes {
		if p != pr {
			newArr = append(newArr, p)
		}
	}

	e.Programmes = newArr
}

func (e *EPG) deleteChannel(ch *Channel) {
	var newArr []*Channel

	for _, c := range e.Channels {
		if c != ch {
			newArr = append(newArr, c)
		}
	}

	e.Channels = newArr
}

// GetChannelByID returns channel by ID or nil if not found
func (e *EPG) GetChannelByID(id string) *Channel {
	return e.channelsByID[id]
}

func (e *EPG) Write(wr io.Writer) (err error) {
	dec := xml.NewEncoder(wr)

	if err = dec.Encode(e); err != nil {
		return
	}

	return
}

func (e *EPG) String() string {
	dbg := ""

	dbg += "CHANNELS:\n"
	for _, ch := range e.Channels {
		dbg += fmt.Sprintf("- %s\n", ch.DisplayName.Text)
	}

	dbg += "\nPROGRAMMES\n"
	for _, pr := range e.Programmes {
		dbg += fmt.Sprintf("- %s\n", pr.Title.Text)
	}

	return dbg
}

// ReadFromFile attempts to read and parse EPG from a file path
func ReadFromFile(path string) (*EPG, error) {
	e := new(EPG)

	file, err := os.Open(path) //nolint
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err := e.Read(file); err != nil {
		return nil, err
	}

	return e, nil
}
