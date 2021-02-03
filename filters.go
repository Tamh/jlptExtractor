package main

func filterCharacters(vs []KanjidicCharacter, f func(KanjidicCharacter) bool) []KanjidicCharacter {
	vsf := make([]KanjidicCharacter, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func filterReadings(vs []KanjidicReading, f func(KanjidicReading) bool) []KanjidicReading {
	vsf := make([]KanjidicReading, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func filterMeanings(vs []KanjidicMeaning, f func(KanjidicMeaning) bool) []KanjidicMeaning {
	vsf := make([]KanjidicMeaning, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func obtainDictNum(vs []KanjidicDicNumber, f func(KanjidicDicNumber) bool) string {
	for _, v := range vs {
		if f(v) {
			return v.Value
		}
	}
	return ""
}

func filterDictEntries(vs []JmdictEntry, f func(JmdictEntry) bool) []JmdictEntry {
	vsf := make([]JmdictEntry, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
