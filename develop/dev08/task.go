package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"syscall"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*


Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

/*
PS C:\Users\vplotnik\localTruecoderRepos\l2-wb-tasks\develop\dev08> golint
PS C:\Users\vplotnik\localTruecoderRepos\l2-wb-tasks\develop\dev08> go vet
*/

func main() {
	// 0 - need start infinity reader
	reader := bufio.NewReader(os.Stdin)

	for {
		curDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s$ ", curDir) //output current working dir after each command. like welcome to input

		// read command || commands chain (check if supported)
		commands, err := reader.ReadString('\n') // delim by next line
		if err != nil {
			log.Fatal(err)
		}

		err = ExecuteCommand(commands)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

//ExecuteCommand method. Process input
func ExecuteCommand(commands string) error {
	// 0 parse commands :: it can be pipline. split |
	cmds := strings.Split(commands, "|")
	pipes := len(cmds) > 1 // do we process pipe command or single
	var input *bytes.Buffer
	var err error

	// iteratively execute each command
	for _, val := range cmds {
		input, err = RunCommand(val, input, pipes)
		if err != nil {
			return err
		}
	}

	if input == nil || len(input.String()) == 0 {
		fmt.Fprint(os.Stdout, "")
	} else {
		fmt.Fprint(os.Stdout, input.String())
	}

	return nil //err
}

//RunCommand real execution of single command. IN:: command buffer bytes, if pipes
func RunCommand(cmd string, input *bytes.Buffer, pipes bool) (*bytes.Buffer, error) {
	argsSplit := strings.Split(cmd, " ")
	var args []string
	for _, val := range argsSplit {
		if val != "" {
			args = append(args, val) // create array of commands
		}
	}

	// use shell interface to exec some commands
	var sh iShell
	switch args[0] {
	case "cd":
		sh = &changeDirStruct{}
	case "pwd":
		sh = &presentWorkingDirStruct{}
	case "echo":
		sh = &echoStruct{}
	case "kill":
		sh = &killStruct{}
	case "ps":
		sh = &processStatusStruct{}
	case "exec":
		sh = &executeStruct{}
	case "nc":
		sh = &netCatStruct{}
	case "quit":
		sh = &quitStruct{}
	}

	if sh != nil {
		return sh.run(args[1:], input, pipes)
	}

	return nil, nil

}

//shell interface
type iShell interface {
	run([]string, *bytes.Buffer, bool) (*bytes.Buffer, error)
}

//struct for each command. inherit shell interface
type changeDirStruct struct{}
type presentWorkingDirStruct struct{}
type echoStruct struct{}
type killStruct struct{}
type processStatusStruct struct{}
type executeStruct struct{}
type netCatStruct struct{}
type quitStruct struct{}

//run for Change Direcoty command.
func (sh *changeDirStruct) run(args []string, input *bytes.Buffer, pipes bool) (*bytes.Buffer, error) {
	if !pipes {
		if len(args) > 0 {
			return nil, os.Chdir(args[0]) // simply change dir.
		}
	}
	return nil, nil
}

//run for get current work dir
func (sh *presentWorkingDirStruct) run(args []string, input *bytes.Buffer, pipes bool) (*bytes.Buffer, error) {
	var output bytes.Buffer

	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	output.WriteString(dir + "\n")
	return &output, nil
}

//run for echo
func (sh *echoStruct) run(args []string, input *bytes.Buffer, pipes bool) (*bytes.Buffer, error) {
	var output bytes.Buffer
	var err error

	if len(args) > 0 {
		_, err = output.WriteString(args[0] + "\n")
	} else {
		_, err = output.WriteString("\n")
	}

	if err != nil {
		return nil, err
	}

	return &output, nil
}

func (sh *killStruct) run(args []string, input *bytes.Buffer, pipes bool) (*bytes.Buffer, error) {
	if len(args) > 0 {
		pid, err := strconv.Atoi(args[0])
		if err != nil {
			return nil, err
		}
		p, err := os.FindProcess(pid)

		if err != nil {
			return nil, err
		}
		p.Signal(syscall.SIGTERM)
		//syscall.Kill(pid, syscall.SIGKILL)

	}
	return nil, nil
}

type proc struct {
	name string
	pid  int
}

func (sh *processStatusStruct) run(args []string, input *bytes.Buffer, pipes bool) (*bytes.Buffer, error) {
	var output bytes.Buffer

	ps, err := filepath.Glob("/proc/*/exe")
	if err != nil {
		return nil, err
	}

	var procs []proc
	for _, file := range ps {
		link, _ := os.Readlink(file)
		if len(link) > 0 {
			name := filepath.Base(link)
			pid, err := strconv.Atoi(strings.Split(file, "/")[2])

			if err == nil {
				procs = append(procs, proc{name, pid})
			}
		}
	}

	sort.Slice(procs, func(i, j int) bool { return procs[i].pid < procs[j].pid })

	output.WriteString("PID Name\n")
	for _, val := range procs {
		output.WriteString(fmt.Sprintf("%d %s\n", val.pid, val.name))
	}
	return &output, nil
}

func (sh *executeStruct) run(args []string, input *bytes.Buffer, pipes bool) (*bytes.Buffer, error) {
	com := exec.Command(args[0], args[1:]...)
	var output bytes.Buffer

	if input != nil {
		com.Stdin = bytes.NewReader((*input).Bytes())
	}

	com.Stdout = &output
	com.Stderr = os.Stderr
	if err := com.Run(); err != nil {
		return nil, err
	}

	return &output, nil
}

func (shell *netCatStruct) run(args []string, input *bytes.Buffer, pipes bool) (*bytes.Buffer, error) {
	if len(args) < 3 {
		return nil, errors.New("too few args (type, ip, port)")
	}

	conn, err := net.Dial(args[0], args[1]+":"+args[2])
	if err != nil {
		return nil, err
	}

	var str string
	fmt.Printf(">")
	fmt.Scan(&str)

	_, err = conn.Write([]byte(str))

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (shell *quitStruct) run(args []string, input *bytes.Buffer, pipes bool) (*bytes.Buffer, error) {
	if !pipes {
		os.Exit(0)
	}
	return nil, nil
}
