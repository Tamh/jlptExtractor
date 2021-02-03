package main

import "strconv"

func compareDictNum(a, b KanjidicCharacter) bool {
	nelsonNum1, _ := strconv.Atoi(obtainDictNum(a.DictionaryNumbers, dictNumIsHeisig))
	nelsonNum2, _ := strconv.Atoi(obtainDictNum(b.DictionaryNumbers, dictNumIsHeisig))
	return nelsonNum1 < nelsonNum2
}

func compareEntryLength(a, b JmdictBasicEntry) bool {
	return len(a.Jukugo) < len(b.Jukugo)
}
