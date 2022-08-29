package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.

Утилита sort сортирует строки текстового файла
По умолчанию сортировка - по всей строке,
можно разбить каждую строку по разделителю на столбцы и отсортировать строки по выбранному столбцу
*/

func main() {
	// [0] base options
	var nFlag, rFlag, uFlag bool
	var kFlag int

	flag.BoolVar(&nFlag, "n", false, "sort by numeric value")
	flag.BoolVar(&rFlag, "r", false, "sort in reverse order")
	flag.BoolVar(&uFlag, "u", false, "do not output duplicate lines")
	flag.IntVar(&kFlag, "k", 0, "sortable column")

	// [1] extra options
	var MFlag, bFlag, cFlag bool
	var hFlag int

	flag.BoolVar(&MFlag, "M", false, "month")
	flag.BoolVar(&bFlag, "b", false, "number")
	flag.BoolVar(&cFlag, "c", false, "check")
	flag.IntVar(&hFlag, "h", 0, "suffix")

	flag.Parse()

	// [1] - get pattern and file
	//args := flag.Args()
	filePath := os.Args[len(os.Args)-1] //0 arg - get file to parse

	// [2] read file
	text := ReadFile(filePath)
	//fmt.Println(text)

	// [3] process
	text = Sort(text, kFlag, nFlag, rFlag, uFlag, bFlag, cFlag)

	// [4] output
	Output(text)

}

//array of lines
type Line struct {
	idx   int
	line  string
	print bool
}

//ReadFile  IN::file path, OUT::text structure  [0] -
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
	idx := 0
	for fileScanner.Scan() {
		//txt = append(txt, Line{idx, fileScanner.Text(), false})
		txt = append(txt, fileScanner.Text())
		idx++
	}
	//Error handling
	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	return txt
}

//Sort :: sort by col, numbers, reverse, duplicates. ignore, check
func Sort(text []string, sortByColumn int, sortByNumbers, sortReverse, noDuplicates, ignoreEndSpace, checkSort bool) []string {
	// [0] Если checkSort, проверяем, отсортированы ли данные
	if checkSort {
		res := text
		sorted := text
		sort.Strings(sorted)
		if reflect.DeepEqual(res, sorted) { // Сравним отсортированный с таким же, но несортированным, и если они равны, то:
			fmt.Println("sorted")
		} else { // Иначе:
			fmt.Println("unsorted")
		}
		return nil
	}

	// [1] обрезаем пробельные символы в конце строки
	if ignoreEndSpace {
		for idx, line := range text {
			text[idx] = strings.TrimSpace(line)
		}
	}

	// [2] сортируем
	if sortByColumn >= -1 {
		text = ColumnBySort(text, sortByColumn)
	}

	//[3] НЕ выводим повторяющиеся строки
	if noDuplicates {
		un := map[string]struct{}{}
		for _, r := range text {
			un[r] = struct{}{}
		}
		text = make([]string, 0, len(un))
		for k := range un {
			text = append(text, k)
		}
	}

	//[4] sort by numbers
	if sortByNumbers {
		nums := make([]int, 0, len(text)) // Создаём слайс целых чисел с cap равной количеству строк для сортировки
		if checkSort {                    // Проверяет, отсортировано ли содержимое файла
			if !sort.IntsAreSorted(nums) {
				return []string{"Need to sort"}
			} else {
				return []string{"No need to sort"}
			}
		}
		for k := range text { // Перебираем индексы строк
			n, err := strconv.Atoi(text[k])
			if err != nil {
				fmt.Println("error sorting by number:", err)
				os.Exit(1)
			}
			nums = append(nums, n)
		}
		if sortReverse {
			sort.Sort(sort.Reverse(sort.IntSlice(nums)))
		} else {
			sort.Ints(nums)
		}

		text = make([]string, 0, len(nums))
		for _, n := range nums {
			text = append(text, strconv.Itoa(n))
		}
	}

	//[5] sort reversed mod
	if sortReverse {
		SortReverse(text)
	}

	return text
}

//ColumnSort - sort text be exact column. If in line less columns, it will be sorted by last value in it and moved DOWN
func ColumnBySort(text []string, columnNum int) []string {
	if columnNum == -1 { // Номер столбца по которому сортируем. По умолчанию - вся строка (до перевода строки)
		sort.Strings(text)
		return text
	}

	sort.Slice(text, func(i, j int) bool {
		lhs := strings.Split(text[i], " ")
		rhs := strings.Split(text[j], " ")
		if len(lhs) <= columnNum || len(rhs) <= columnNum { // Если количество столбцов в двух соседних строках меньше, чем н
			return lhs[0] < rhs[0] // Сортируем по возрастанию слова в столбце
		}
		return strings.Split(text[i], " ")[columnNum] <
			strings.Split(text[j], " ")[columnNum]
	})
	return text
}

//SortReverse
func SortReverse(text []string) []string {
	for i, j := 0, len(text)-1; i < j; i, j = i+1, j-1 {
		text[i], text[j] = text[j], text[i]
	}
	return text
}

//Output
func Output(text []string) {
	for _, v := range text {
		fmt.Println(v)
	}
}
