package main

import "io"

// Jmdict dict
type Jmdict struct {
	Entries []JmdictEntry `xml:"entry"`
}

// JmdictEntry entry
type JmdictEntry struct {
	Sequence int             `xml:"ent_seq"`
	Kanji    []JmdictKanji   `xml:"k_ele"`
	Readings []JmdictReading `xml:"r_ele"`
	Sense    []JmdictSense   `xml:"sense"`
}

// JmdictKanji kanji
type JmdictKanji struct {
	Expression  string   `xml:"keb"`
	Information []string `xml:"ke_inf"`
	Priorities  []string `xml:"ke_pri"`
}

// JmdictReading reading
type JmdictReading struct {
	Reading      string   `xml:"reb"`
	NoKanji      *string  `xml:"re_nokanji"`
	Restrictions []string `xml:"re_restr"`
	Information  []string `xml:"re_inf"`
	Priorities   []string `xml:"re_pri"`
}

// JmdictSource source
type JmdictSource struct {
	Content  string  `xml:",chardata"`
	Language *string `xml:"lang,attr"`
	Type     *string `xml:"ls_type,attr"`
	Wasei    string  `xml:"ls_wasei,attr"`
}

// JmdictGlossary glossary
type JmdictGlossary struct {
	Content  string  `xml:",chardata"`
	Language *string `xml:"lang,attr"`
	Gender   *string `xml:"g_gend"`
}

// JmdictSense sense
type JmdictSense struct {
	RestrictedKanji    []string         `xml:"stagk"`
	RestrictedReadings []string         `xml:"stagr"`
	References         []string         `xml:"xref"`
	Antonyms           []string         `xml:"ant"`
	PartsOfSpeech      []string         `xml:"pos"`
	Fields             []string         `xml:"field"`
	Misc               []string         `xml:"misc"`
	SourceLanguages    []JmdictSource   `xml:"lsource"`
	Dialects           []string         `xml:"dial"`
	Information        []string         `xml:"s_inf"`
	Glossary           []JmdictGlossary `xml:"gloss"`
}

// JmdictBasicEntry BasicEntry
type JmdictBasicEntry struct {
	Jukugo   string
	Reading  string
	Glossary string
}

// LoadJmdict Loads a Jmdict file
func LoadJmdict(reader io.Reader) (Jmdict, map[string]string, error) {
	var dict Jmdict
	entities, err := parseDict(reader, &dict, true)
	return dict, entities, err
}
