package main

import (
	"fmt"
	"sort"
	"strings"
)

func printReadings(vs []KanjidicReading) string {
	var rv = ""
	for _, v := range vs {
		rv += v.Value + "、"
	}
	return strings.TrimRight(rv, "、")
}

func printMeanings(vs []KanjidicMeaning) string {
	var rv = ""
	for _, v := range vs {
		rv += v.Meaning + "; "
	}
	return strings.TrimRight(rv, "; ")
}

func printCharacterLiterals(chars []KanjidicBasicEntry) string {
	fullstr := ""
	for _, char := range chars {
		fullstr += char.Kanji
	}
	return fullstr
}

func doesntContain(kanji string, charMap string) bool {
	for _, c := range kanji {
		if !strings.ContainsRune(charMap, c) {
			return false
		}
	}
	return true
}

func extractEntries(characters []KanjidicBasicEntry, entries []JmdictEntry, flatMap string) []KanjidicBasicEntry {
	for index, key := range characters {
		for _, entry := range entries {
			for _, kanjiEntry := range entry.Kanji {
				if strings.Contains(kanjiEntry.Expression, key.Kanji) && doesntContain(kanjiEntry.Expression, flatMap) {
					characters[index].Examples = append(characters[index].Examples, convertToJmdictBasicEntry(kanjiEntry.Expression, entry))
				}
			}
		}
		sort.SliceStable(characters[index].Examples, func(i, j int) bool {
			return compareEntryLength(characters[index].Examples[i], characters[index].Examples[j])
		})
	}
	return characters
}

func countAllEntries(fullList []KanjidicBasicEntry) int {
	numEntries := 0
	for _, baseKanji := range fullList {
		numEntries += len(baseKanji.Examples)
	}
	return numEntries
}

func convertToKanjidicBasicEntry(kanjiList []KanjidicCharacter) []KanjidicBasicEntry {
	var rl = make([]KanjidicBasicEntry, 0)
	for _, kanji := range kanjiList {
		var basicEntry KanjidicBasicEntry
		basicEntry.Kanji = kanji.Literal
		basicEntry.Onyomi = printReadings(filterReadings(kanji.ReadingMeaning.Readings, readingIsJPOn))
		basicEntry.Kunyomi = printReadings(filterReadings(kanji.ReadingMeaning.Readings, readingIsJPKun))
		basicEntry.Meaning = printMeanings(filterMeanings(kanji.ReadingMeaning.Meanings, meaningInES))
		rl = append(rl, basicEntry)
	}
	return rl
}

func convertToJmdictBasicEntry(jukugo string, entry JmdictEntry) JmdictBasicEntry {
	var r JmdictBasicEntry
	r.Jukugo = jukugo

	var strReadings = ""
	for _, reading := range entry.Readings {
		strReadings += reading.Reading + "、"
	}
	r.Reading = strings.TrimRight(strReadings, "、")

	var strMeanings = ""
	for _, sense := range entry.Sense {
		for _, glossary := range sense.Glossary {
			if glossary.Language != nil && *glossary.Language == "spa" {
				strMeanings += glossary.Content + "; "
			}
		}
	}
	r.Glossary = strings.TrimRight(strMeanings, "; ")

	return r
}

func entryToString(entry JmdictBasicEntry) string {
	return entry.Jukugo + "\t" + entry.Reading + "\t" + entry.Glossary
}

func characterToString(char KanjidicBasicEntry, printEntries bool) string {
	rs := fmt.Sprintln(char.Kanji + "\t" + char.Onyomi + "\t" + char.Kunyomi + "\t" + char.Meaning)
	if printEntries {
		for _, entry := range char.Examples {
			rs += fmt.Sprintln("\t" + entryToString(entry))
		}
	}
	return rs
}

func dictionaryToString(dict []KanjidicBasicEntry, printEntries bool) string {
	rs := ""
	for _, char := range dict {
		rs += fmt.Sprint(characterToString(char, printEntries))
	}
	return rs
}
