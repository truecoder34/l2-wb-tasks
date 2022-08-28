package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

/*
Ex.
$> go run .\task.go -A=1000  .\grep_man.txt "00"
*/

func main() {
	// [0] - init flags
	var AFlag int
	flag.IntVar(&AFlag, "A", 0, "after")
	BFlag := flag.Int("B", 10, "before")
	CFlag := flag.Int("C", 10, "context")

	cFlag := flag.Bool("c", false, "count")
	iFlag := flag.Bool("i", false, "ignore-case")
	vFlag := flag.Bool("v", false, "invert")
	FFlag := flag.Bool("F", false, "fixed")
	nFlag := flag.Bool("n", false, "line num")

	tFlag := flag.String("t", "", "template")
	flag.Parse()

	// [1] - get pattern and file
	args := flag.Args()
	filePath := args[0] //0 arg - get file to parse
	pattern := args[1]  //1 arg - pattern

	// [2] - read text
	fileStruct := ReadFile(filePath)

	fmt.Printf("[file]::%s\n", filePath)
	fmt.Printf("[pattern]::%s\n", pattern)
	//fmt.Println(*AFlag)
	//fmt.Println(fileStruct)
	fmt.Printf("[INPUT]:: A = %v, B = %v, C = %v, c = %v, i = %v, v = %v, F = %v, n = %v, t = %v\n", AFlag, *BFlag, *CFlag, *cFlag, *iFlag, *vFlag, *FFlag, *nFlag, *tFlag)

	// if *tFlag == "" {
	// 	log.Fatal("need template (-t)")
	// }

	resText, err := LookForMatch(fileStruct, *tFlag, *FFlag, *iFlag) //match
	if err != nil {
		log.Fatal(err)
	}
	countStrs := Count(resText) //count strings with match

	resText = BeforeAfterMatch(resText, AFlag, *BFlag, *CFlag) //if need more strings for output

	if *vFlag { //invert
		resText = Inversion(resText)
	}

	if *cFlag { //print result
		fmt.Println(countStrs)
	} else {
		Output(resText, *nFlag)
	}

}

type Line struct {
	line  string
	print bool
}

// [1] - Init text structure. OUT:: Text struct
func NewText() []Line {
	var txt []Line
	return txt
}

// [0] - read file. IN::file path, OUT::text structure
func ReadFile(filePath string) []Line {
	txt := NewText()

	//open file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//read line by line
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		txt = append(txt, Line{fileScanner.Text(), false})
	}
	//Error handling
	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	return txt
}

//[2] - handle template match. Update array[] of Line in place
func LookForMatch(resText []Line, template string, fixed, noCase bool) ([]Line, error) {
	// 0 - if ignore case. template to lower,
	if noCase {
		template = strings.ToLower(template)
	}

	// 1 - compile regexp
	regex := regexp.MustCompile(template)

	// 2 - fixed case
	if fixed {
		for i, line := range resText {
			if noCase {
				if strings.ToLower(line.line) == template {
					resText[i].print = true
				}
			} else {
				if line.line == template {
					resText[i].print = true
				}
			}
		}
	} else {
		var temp string
		for i, line := range resText {
			if noCase {
				temp = strings.ToLower(line.line)
				if regex.Match([]byte(temp)) {
					resText[i].print = true
				}
			} else {
				if regex.Match([]byte(line.line)) {
					resText[i].print = true
				}
			}
		}
	}
	return resText, nil
}

//[3] - process before|after|context options
func BeforeAfterMatch(resText []Line, after, before, context int) []Line {
	//0 - process context
	if context > 0 {
		after = context
		before = context
	}
	// 1 - after process
	if after > 0 {
		for i, val := range resText {
			if val.print {
				for j, c := 1, after; c > 0; j, c = j+1, c-1 {
					if i-j < 0 {
						break
					}
					resText[i-j].print = true
				}
			}
		}
	}
	// 2 - before process
	if before > 0 {
		for i := len(resText) - 1; i >= 0; i-- {
			if resText[i].print {
				for j, c := 1, before; c > 0; j, c = j+1, c-1 {
					if i+j > len(resText)-1 {
						break
					}
					resText[i+j].print = true
				}
			}
		}
	}

	return resText
}

//[4] - count marked to print lines
func Count(resText []Line) int {
	var counter int
	for _, line := range resText {
		if line.print {
			counter++
		}
	}
	return counter
}

//[5] - output method. Process line numbering case
func Output(resText []Line, numbersLine bool) {
	for i, line := range resText {
		if line.print {
			if numbersLine {
				fmt.Println(i+1, line.line)
			} else {
				fmt.Println(line.line)
			}
		}
	}
}

//[6] - inversion handler
func Inversion(resText []Line) []Line {
	for _, line := range resText {
		if line.print {
			line.print = false
		} else {
			line.print = true
		}
	}
	return resText
}
