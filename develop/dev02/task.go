package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку.
Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

/*
=== IMPLEMENT ===
[1]
> go mod init github.com/truecoder34/l2-wb-tasks/develop/dev02
PS C:\Users\vplotnik\localTruecoderRepos\l2-wb-tasks\develop\dev02> golint
PS C:\Users\vplotnik\localTruecoderRepos\l2-wb-tasks\develop\dev02> go vet

TODO : переделать решение и разбирать строку с конца.

*/

/*
Unpack - method to unpack input string. Return unzipped string
1 - convert to run
2 - go range rune
*/
func Unpack(inputLine string) (*string, error) {
	var resUnpacked []rune
	rns := []rune(inputLine)
	isSlash := false
	var count int

	// 0 - check input line for emptiness or consist only of spaces
	if len(rns) == 0 || len(strings.TrimSpace(inputLine)) == 0 {
		return &inputLine, nil
	}
	// 1 - validate that first elem is not num, if num - error
	if ValidateString(rns) {
		return nil, errors.New("incorrect string, first element is number")
	}
	// 2 - go through range of runes
	for idx, ru := range rns {
		fmt.Printf("[%d] - rune in byte = %x in symbol = %c\n", idx, ru, ru)
		if ru == '\\' && !isSlash {
			isSlash = true
			fmt.Printf("[%d] - faced double slash rune %v\n", idx, ru)
			continue
		}
		if isSlash {
			resUnpacked = append(resUnpacked, ru)
			isSlash = false
			fmt.Printf("[%d] - faced single slash rune %v\n", idx, ru)
			continue
		}
		if unicode.IsNumber(ru) {
			fmt.Printf("[%d] - rune is number; in byte = %x in symbol = %c\n", idx, ru, ru)
			count = int(ru - '0')
			if count == 0 {
				continue
			}
			for j := 0; j < count-1; j++ {
				resUnpacked = append(resUnpacked, rns[idx-1])
			}
			continue
		}

		resUnpacked = append(resUnpacked, ru)
	}

	res := string(resUnpacked)
	return &res, nil
}

/*
ValidateString - check first element. if num - return true, if not num - false
*/
func ValidateString(inputRunes []rune) bool {
	return unicode.IsNumber(inputRunes[0])
}

func main() {
	// res1, _ := Unpack("a4bc2d5e") // aaaabccddddde
	// fmt.Printf("%v , len = %d\n", res1, len(res1))

	// res2, _ := Unpack("") // ""
	// fmt.Printf("%v , len = %d\n", res2, len(res2))

	// res3, _ := Unpack("    ") // "    "
	// fmt.Printf("%v , len = %d\n", res3, len(res3))

	res4, _ := Unpack(`qwe\4\5`) // "    "
	fmt.Printf("%v , len = %d\n", *res4, len(*res4))

	res5, _ := Unpack(`qwe\45`) // "    "
	fmt.Printf("%v , len = %d\n", *res5, len(*res5))

	res6, _ := Unpack(`qwe\\5`) // "    "
	fmt.Printf("%v , len = %d\n", *res6, len(*res6))

	res7, err := Unpack(`5`) // "    "
	if err != nil {
		fmt.Printf("%v - %v", res7, err)
	}
	//fmt.Printf("%v , len = %d\n", res7, len(*res7))

	res8, _ := Unpack(`\\\\4\\5`) // "    "
	fmt.Printf("%v , len = %d\n", *res8, len(*res8))

	res9, _ := Unpack(`\\15\5`) // "    "
	fmt.Printf("%v , len = %d\n", *res9, len(*res9))

}
