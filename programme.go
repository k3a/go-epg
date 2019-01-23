package epg

// Programme contains as single programme of a channel
type Programme struct {
	Start   string `xml:"start,attr"`
	Stop    string `xml:"stop,attr"`
	Channel string `xml:"channel,attr"`
	Title   struct {
		Text string `xml:",chardata"`
		Lang string `xml:"lang,attr"`
	} `xml:"title"`
	SubTitle struct {
		Text string `xml:",chardata"`
		Lang string `xml:"lang,attr"`
	} `xml:"sub-title"`
	Desc struct {
		Text string `xml:",chardata"`
		Lang string `xml:"lang,attr"`
	} `xml:"desc"`
	Category []struct {
		Text string `xml:",chardata"`
		Lang string `xml:"lang,attr"`
	} `xml:"category"`
	Icon *struct {
		Text string `xml:",chardata"`
		Src  string `xml:"src,attr"`
	} `xml:"icon,omitempty"`
	Subtitles  string `xml:"subtitles"`
	Premiere   string `xml:"premiere"`
	EpisodeNum struct {
		Text   string `xml:",chardata"`
		System string `xml:"system,attr"`
	} `xml:"episode-num"`
	Rating *struct {
		Text   string `xml:",chardata"`
		System string `xml:"system,attr"`
		Value  string `xml:"value"`
	} `xml:"rating,omitempty"`
	Date       string `xml:"date"`
	StarRating *struct {
		Text  string `xml:",chardata"`
		Value string `xml:"value"`
	} `xml:"star-rating,omitempty"`
	PreviouslyShown string `xml:"previously-shown"`

	epg *EPG
}
