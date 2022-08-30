package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

const (
	WORKERS = 2
	TIMEOUT = 10
)

func main() {
	//[0] - set up connection timeout. Use 10 seconds default
	timeout := flag.Duration("timeout", time.Second*TIMEOUT, "Timeout")
	flag.Parse()

	// get connection data : host first, port second
	host := flag.Arg(0)
	port := flag.Arg(1)

	fmt.Println("[Settings] :: Timeout:", timeout, "Host:", host, "Port:", port)

	// try to connect via net package DialTimout
	conn, err := net.DialTimeout("tcp", host+":"+port, *timeout)
	if err != nil {
		fmt.Println(err)
	}

	// need to init  wg group to handle read and write in parallel
	wg := sync.WaitGroup{}
	wg.Add(WORKERS)
	go Read(conn, &wg)
	go Write(conn, &wg)

	wg.Wait()

	defer fmt.Println("\n==============================================\nConnection is closed")
}

func Read(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	reader := bufio.NewReader(conn) // init reader from connection
	// infinity loop to read
	for {
		message, err := reader.ReadString(' ') // empty space delimiter. Each message wil be parsed
		if err == io.EOF {                     // if server closed connection
			fmt.Println("\n==============================================\nConnection is closed")
			os.Exit(0)
		}
		if netErr, ok := err.(net.Error); ok && !netErr.Timeout() { // if Scanner closed connection
			break
		}
		fmt.Fprint(os.Stdout, message)
	}
}

//Write - read data from STD IN and push to conn
func Write(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	scanner := bufio.NewScanner(os.Stdin) // init new writer to std in

	for scanner.Scan() {
		_, err := conn.Write([]byte(scanner.Text() + "\n"))
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	defer conn.Close() // if ctrl + D
}
