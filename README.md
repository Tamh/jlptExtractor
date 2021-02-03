# JLPT Extractor
Un proyecto que permite extraer una lista de kanji y ejemplos de uso de los kanji desde kanjidic2 y Jmdict.
El proyecto está creado con una licencia Mozilla Public License 2.0.

El programa genera un archivo de texto con los kanji que están clasificados como recomendados para el examen JLPT N5 y N4, para importar en Anki, o generar memofichas programáticamente.

## Modo de uso
Como las bases de datos Kanjidic2 y JMdict tienen modelos de licenciamiento diferentes, además que se actualizan con cierta frecuencia, no vienen incluidas con el proyecto.
Para más información, ver http://www.edrdg.org/edrdg/licence.html

1. Descargar la distribución más nueva de la base de datos Kanjidic2 desde http://www.edrdg.org/kanjidic/kanjidic2.xml.gz
2. Descargar la distribución más nueva de la base de datos JMdict desde http://ftp.edrdg.org/pub/Nihongo/JMdict.gz
3. Descomprimir ambos archivos en la misma carpeta dónde está el JLPT Extractor
4. Ejecutar `go run .` en la carpeta del JLPT Extractor.

## Generar jukugo
Es posible generar un listado de palabras del diccionario que contienen dichos kanji seleccionados para revisión y estudio del Jukugo (kanji compuesto).
Para ello, editar el archivo main.go, en la línea que contiene la llamada a la función `writeResultsFile` cambiando el segundo parámetro por `true`.

Solo genera lista de jukugo para las palabras que tienen acepciones en español.
