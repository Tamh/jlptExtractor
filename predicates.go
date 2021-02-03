package main

func jlptIs4(v KanjidicCharacter) bool {
	if v.Misc.JlptLevel != nil {
		return *v.Misc.JlptLevel == "4"
	}
	return false
}

func jlptIs3(v KanjidicCharacter) bool {
	if v.Misc.JlptLevel != nil {
		return *v.Misc.JlptLevel == "3"
	}
	return false
}

func readingIsJPOn(v KanjidicReading) bool {
	return v.Type == "ja_on"
}

func readingIsJPKun(v KanjidicReading) bool {
	return v.Type == "ja_kun"
}

func meaningInES(v KanjidicMeaning) bool {
	if v.Language != nil {
		return *v.Language == "es"
	}
	return false
}

func meaningNoLanguage(v KanjidicMeaning) bool {
	return v.Language == nil
}

func dictNumIsNelson(v KanjidicDicNumber) bool {
	return v.Type == "nelson_n"
}

func dictNumIsHeisig(v KanjidicDicNumber) bool {
	return v.Type == "heisig6"
}

/**
 *
 */

func hasSpanishDefinition(v JmdictEntry) bool {
	numGlossary := 0
	for _, sense := range v.Sense {
		for _, glossary := range sense.Glossary {
			if glossary.Language != nil && *glossary.Language == "spa" {
				numGlossary++
			}
		}
	}
	if numGlossary > 0 {
		return true
	}
	return false
}
