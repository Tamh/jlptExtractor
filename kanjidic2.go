package main

import "encoding/xml"

// Kanjidic Structure for the whole Kanjidic
type Kanjidic struct {
	XMLName    xml.Name            `xml:"kanjidic2"`
	Header     KanjidicHeader      `xml:"header"`
	Characters []KanjidicCharacter `xml:"character"`
}

// KanjidicHeader Structure for the Kanjidic header
type KanjidicHeader struct {
	XMLName         xml.Name `xml:"header"`
	FileVersion     string   `xml:"file_version"`
	Databaseversion string   `xml:"database_version"`
	DateOfCreation  string   `xml:"date_of_creation"`
}

// KanjidicCharacter Structure for the Kanjidic characters
type KanjidicCharacter struct {
	XMLName           xml.Name                `xml:"character"`
	Literal           string                  `xml:"literal"`
	Codepoint         []KanjidicCodepoint     `xml:"codepoint>cp_value"`
	Radical           []KanjidicRadical       `xml:"radical>rad_value"`
	Misc              KanjidicMisc            `xml:"misc"`
	DictionaryNumbers []KanjidicDicNumber     `xml:"dic_number>dic_ref"`
	QueryCode         []KanjidicQueryCode     `xml:"query_code>q_code"`
	ReadingMeaning    *KanjidicReadingMeaning `xml:"reading_meaning"`
}

// KanjidicCodepoint Structure for the Kanjidic kanji codepoints
type KanjidicCodepoint struct {
	Value string `xml:",chardata"`
	Type  string `xml:"cp_type,attr"`
}

// KanjidicRadical Structure for the Kanjidic kanji radicals
type KanjidicRadical struct {
	Value string `xml:",chardata"`
	Type  string `xml:"rad_type,attr"`
}

// KanjidicMisc Structure for the Kanjidic miscelaneous fields
type KanjidicMisc struct {
	Grade        *string           `xml:"grade"`
	StrokeCounts []string          `xml:"stroke_count"`
	Variants     []KanjidicVariant `xml:"variant"`
	Frequency    *string           `xml:"freq"`
	RadicalName  []string          `xml:"rad_name"`
	JlptLevel    *string           `xml:"jlpt"`
}

// KanjidicVariant Structure for the Kanjidic kanji variants
type KanjidicVariant struct {
	Value string `xml:",chardata"`
	Type  string `xml:"var_type"`
}

// KanjidicDicNumber Structure for the Kanjidic dictionary numbers
type KanjidicDicNumber struct {
	Value  string `xml:",chardata"`
	Type   string `xml:"dr_type,attr"`
	Volume string `xml:"m_vol,attr"`
	Page   string `xml:"m_page,attr"`
}

// KanjidicQueryCode Structure for the Kanjidic query codes
type KanjidicQueryCode struct {
	Value             string `xml:",chardata"`
	Type              string `xml:"qc_type,attr"`
	Misclassification string `xml:"skip_misclass,attr"`
}

// KanjidicReadingMeaning Structure for the Kanjidic kanji readings/meanings
type KanjidicReadingMeaning struct {
	Readings []KanjidicReading `xml:"rmgroup>reading"`
	Meanings []KanjidicMeaning `xml:"rmgroup>meaning"`
	Nanori   []string          `xml:"nanori"`
}

// KanjidicReading Structure for a Kanjidic kanji reading
type KanjidicReading struct {
	Value        string  `xml:",chardata"`
	Type         string  `xml:"r_type,attr"`
	OnType       *string `xml:"on_type"`
	JouyouStatus *string `xml:"r_status"`
}

// KanjidicMeaning Structure for a Kanjidic kanji meaning
type KanjidicMeaning struct {
	Meaning  string  `xml:",chardata"`
	Language *string `xml:"m_lang,attr"`
}

// KanjidicBasicEntry basic entry
type KanjidicBasicEntry struct {
	Kanji    string
	Onyomi   string
	Kunyomi  string
	Meaning  string
	Examples []JmdictBasicEntry
}
