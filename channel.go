package epg

// Channel represents a single channel containing programmes
type Channel struct {
	ID          string `xml:"id,attr"`
	DisplayName struct {
		Text string `xml:",chardata"`
		Lang string `xml:"lang,attr"`
	} `xml:"display-name"`
	URL  string `xml:"url"`
	Icon *struct {
		Text string `xml:",chardata"`
		Src  string `xml:"src,attr"`
	} `xml:"icon,omitempty"`

	epg        *EPG
	Programmes []*Programme
}

// Delete deletes channel with programmes from the associated EPG
func (ch *Channel) Delete() {
	if ch.epg == nil {
		return
	}

	for _, pr := range ch.Programmes {
		ch.epg.deleteProgramme(pr)
	}

	ch.epg.deleteChannel(ch)
}
