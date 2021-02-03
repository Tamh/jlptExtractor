package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

// Function to open the Kanjidic2 file
func readKanjidic() ([]KanjidicBasicEntry, string, error) {
	fmt.Println("Abriendo el archivo Kanjidic2...")
	xmlFile, err := os.Open("kanjidic2.xml")
	defer xmlFile.Close()
	if err != nil {
		return nil, "", errors.New("no se pudo abrir el archivo kanjidic2.xml")
	}
	fmt.Println("Satisfactoriamente abierto kanjidic2.xml")

	fmt.Println("Leyendo el archivo Kanjidic2...")
	byteValue, _ := ioutil.ReadAll(xmlFile)

	var kanjidic Kanjidic
	fmt.Println("Interpretando el archivo Kanjidic2...")
	err = xml.Unmarshal(byteValue, &kanjidic)
	xmlFile.Close()

	if err != nil {
		return nil, "", errors.New("no se pudo intertpretar el archivo kanjidic2.xml")
	}

	fmt.Println("Extrayendo los kanji JLPT5...")
	var jlpt4Chars = filterCharacters(kanjidic.Characters, jlptIs4)
	fmt.Println("Extrayendo los kanji JLPT4...")
	var jlpt3Chars = filterCharacters(kanjidic.Characters, jlptIs3)

	fmt.Println("Mezclando y ordenando las listas...")
	var allChars = append(jlpt4Chars, jlpt3Chars...)
	sort.SliceStable(allChars, func(i, j int) bool {
		return compareDictNum(allChars[i], allChars[j])
	})

	var allCharsList = convertToKanjidicBasicEntry(allChars)
	var allCharsFlatMap = printCharacterLiterals(allCharsList)

	fmt.Println("Encontrados", len(kanjidic.Characters), "kanji en total en Kanjidic.")
	fmt.Println("Filtrados", len(jlpt4Chars), "kanji JLPT N5.")
	fmt.Println("Filtrados", len(jlpt3Chars), "kanji JLPT N4.")

	return allCharsList, allCharsFlatMap, nil
}

func readJmdict(allCharsList []KanjidicBasicEntry, allCharsFlatMap string) ([]KanjidicBasicEntry, error) {
	fmt.Println("Abriendo el archivo JMdict...")
	xmlFile2, err := os.Open("JMdict")
	defer xmlFile2.Close()
	if err != nil {
		return nil, errors.New("no se pudo abrir el archivo JMdict")
	}
	fmt.Println("Satisfactoriamente abierto JMdict")

	fmt.Println("Leyendo e interpretando el archivo JMdict...")
	var jmdict, _, err2 = LoadJmdict(xmlFile2)
	if err2 != nil {
		return nil, errors.New("no se pudo leer o interpretar el archivo JMdict")
	}
	xmlFile2.Close()

	var allSpanishEntries = filterDictEntries(jmdict.Entries, hasSpanishDefinition)

	fmt.Println("Encontradas", len(jmdict.Entries), "entradas de diccionario en total.")
	fmt.Println("Filtradas", len(allSpanishEntries), "entradas en Español.")

	finalMap := extractEntries(allCharsList, allSpanishEntries, allCharsFlatMap)
	fmt.Println("Encontradas", countAllEntries(finalMap), "entradas de diccionario en español para los kanji N4 y N5.")

	return finalMap, nil
}

func writeResultsFile(finalMap []KanjidicBasicEntry, printEntries bool) error {
	resultFile, err := os.Create("kanjiResults.tsv")
	if err != nil {
		return errors.New("no se pudo crear el archivo kanjiResults.tsv")
	}
	defer resultFile.Close()

	fmt.Fprintln(resultFile, dictionaryToString(finalMap, printEntries))
	fmt.Println("Archivo kanjiResults.tsv escrito.")
	return nil
}

func main() {
	var allCharsList, allCharsFlatMap, err = readKanjidic()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	var finalMap, err2 = readJmdict(allCharsList, allCharsFlatMap)
	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(2)
	}

	var err3 = writeResultsFile(finalMap, false)
	if err3 != nil {
		fmt.Println(err3.Error())
		os.Exit(2)
	}

}
