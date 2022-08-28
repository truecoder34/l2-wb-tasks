package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// [0] - process input files
	fFlag := flag.Int("f", 1, "fields")
	dFlag := flag.String("d", "\t", "delimiter")
	sFlag := flag.Bool("s", false, "separated")
	flag.Parse()

	// [1] - get pattern and file
	args := flag.Args()
	filePath := args[0] //0 arg - get file to parse

	//[2] read file
	text := ReadFile(filePath)

	res, err := Cut(text, *fFlag, *dFlag, *sFlag)
	if err != nil {
		log.Fatal(err)
	}

	Output((res))
}

//Read ile. IN::file path, OUT::text structure [0]
func ReadFile(filePath string) []string {
	var txt []string

	//open file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//read line by line
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		txt = append(txt, fileScanner.Text())
	}
	//Error handling
	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	return txt
}

/*
Cut method based on options [2]
*/
func Cut(text []string, fields int, delimiter string, separated bool) ([]string, error) {
	var cut []string
	var temp []string

	for _, val := range text {
		temp = strings.Split(val, delimiter)
		if separated && len(temp) < 2 {
			continue
		}
		if len(temp) >= fields {
			cut = append(cut, temp[fields-1])
		}
	}
	return cut, nil
}

//Output method. [3]
func Output(text []string) {
	for _, val := range text {
		fmt.Println(val)
	}
}
