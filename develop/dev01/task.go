package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.


[1] - INIT PACKAGE
> go mod init github.com/truecoder34/l2-wb-tasks/develop/dev01
go: creating new go.mod: module github.com/truecoder34/l2-wb-tasks/develop/dev01
go: to add module requirements and sums:
        go mod tidy

[2] download ntp
> go get github.com/beevik/ntp

[3] go vet https://golang-blog.blogspot.com/2019/07/vet-command-golang.html
> go vet -json ..\dev01\
# github.com/truecoder34/l2-wb-tasks/develop/dev01
{}


[4] golint
> go get -u golang.org/x/lint/golint
> golint
task.go:44:1: comment on exported function GetTime should be of the form "GetTime ..."
task.go:56:1: comment on exported function GetCurrentTime should be of the form "GetCurrentTime ..."
*/

/*
GetTime - Call ntp.Time("0.beevik-ntp.pool.ntp.org") and return Time pointer on value back
*/
func GetTime() (*time.Time, error) {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Printf("Error occurred on time request : %v\n", err)
		return nil, err
	}
	return &time, nil
}

/*
GetCurrentTime - do it via query with extension of default settings
*/
func GetCurrentTime() (*time.Time, error) {
	options := ntp.QueryOptions{Timeout: 5 * time.Second, TTL: 25}
	response, err := ntp.QueryWithOptions("0.beevik-ntp.pool.ntp.org", options)
	time := time.Now().Add(response.ClockOffset)
	if err != nil {
		fmt.Printf("Error occurred on time request : %v\n", err)
		return nil, err
	}
	return &time, nil
}

func main() {
	timeRec, err := GetTime()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("[TIME] : %v", timeRec)

	timeRec2, err := GetCurrentTime()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("[TIME] : %v", timeRec2)

}
