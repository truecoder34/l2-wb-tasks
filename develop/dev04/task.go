package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

/*
Search Anagrams in vocabulary.
[IN] : vocabulary array
[OUT]: map of anagrams . key - word ; value -
*/
func searchAnagrams(vocabulary []string) *map[string][]string {
	mapResult := make(map[string][]string)
	mapDuplicate := make(map[string][]int)

	// [1] - sort  EACH word in vacab by leters - O(n)
	for idx, word := range vocabulary {
		sortdWord := sortLettersInWord(word)
		mapDuplicate[sortdWord] = append(mapDuplicate[sortdWord], idx)
	}

	// [2]
	for key := range mapDuplicate {
		if len(mapDuplicate[key]) > 1 {
			for _, numberWord := range mapDuplicate[key] {
				mapResult[vocabulary[mapDuplicate[key][0]]] = append(mapResult[vocabulary[mapDuplicate[key][0]]], vocabulary[numberWord])
			}
		}
	}

	//[3]
	for mapR := range mapResult {
		mapResult[mapR] = sortArrayOfStrings(mapResult[mapR])
		mapResult[mapR] = removeDuplicates(mapResult[mapR])
	}

	return &mapResult
}

//=====================================================================

/*
[0] - Read Anagrams file: .list format
*/
func loadAnagrams(path2file string) *[]string {
	//[0] - open file
	file, err := os.Open(path2file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var vocabulary []string

	//[1] - read file line by line
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		vocabulary = append(vocabulary, strings.ToLower(fileScanner.Text()))
	}

	// [2] -
	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("[READING]: Anagrams file ...")

	return &vocabulary
}

//=====================================================================

/*
[1] - Sort words in array by Alphabetic order
*/
func sortArrayOfStrings(words []string) []string {
	// via implementation of sort interface
	sort.Slice(words, func(i, j int) bool { return words[i] < words[j] })
	return words
}

//=====================================================================

/*
[2] - Sort elements (letters\runes) in word in Alphabetic order
*/
func sortLettersInWord(word string) string {
	toRune := []rune(word)
	sort.Slice(toRune, func(i, j int) bool { return toRune[i] < toRune[j] })
	return string(toRune)
}

//=====================================================================

/*
[3] - remove duplicates from array
*/
func removeDuplicates(vocab []string) []string {
	if len(vocab) == 0 {
		return nil
	}
	var cleanedVocab []string
	for i := 1; i < len(vocab); i++ {
		if vocab[i] != vocab[i-1] {
			cleanedVocab = append(cleanedVocab, vocab[i-1])
		}
	}
	cleanedVocab = append(cleanedVocab, vocab[len(vocab)-1])

	return cleanedVocab
}

//=====================================================================

func main() {
	vcb := loadAnagrams("C:\\Users\\vplotnik\\localTruecoderRepos\\l2-wb-tasks\\develop\\dev04\\anagrams.list")
	fmt.Printf("[INPUT] : %v\n", vcb)

	anagrams := searchAnagrams(*vcb)
	fmt.Printf("[OUTPUT] : %v\n", anagrams)
}
