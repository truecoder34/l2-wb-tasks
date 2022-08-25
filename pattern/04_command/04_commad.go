package main

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
	Использование:
	1) когда вы хотите параметризировать объекты выполняемым действием
	2) когда вы хотите ставить операции в очередь, выполнять их по расписанию или передовать по сети
	3) когда вам нужна операция отмены
	+:
	1) убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют
	2) позволяет реализовать простую отмену и повтор операций
	3) позволяет реализовать отложенный запуск операций
	4) позволяет собирать сложные команды из простых
	5) реализует принцип открытости/закрытости
	-:
	1) усложняет код программы из-за введения множества дополнительных классов
*/

/*
[1] https://levelup.gitconnected.com/the-command-pattern-with-go-fd5dabc84c7
[2] https://www.sohamkamani.com/golang/command-pattern/
[3] https://golangbyexample.com/command-design-pattern-in-golang/
*/

/*
[0] - we have an interface of job scheduler
it can Run() , Stop(), Kill(), GetInfo() about abstract job\process
*/
type JobScheduler interface {
	Run()
	Stop()
	Kill()
	GetInfo()
}

/*
[1] - интерфейс нашей команды. Главный метод - выполнить команду
*/
type Command interface {
	Execute()
}

//============================================================

/*
[2] - создаем под КАЖДУЮ КОМАНДУ структуру
Каждая структура-клманда реализует ИНТЕРФЕЙС комманды
*/
type RunCommand struct {
	scheduler JobScheduler
}

func (rc *RunCommand) Execute() {
	rc.scheduler.Run()
}

type StopCommand struct {
	scheduler JobScheduler
}

func (sc *StopCommand) Execute() {
	sc.scheduler.Stop()
}

type KillCommand struct {
	scheduler JobScheduler
}

func (kc *KillCommand) Execute() {
	kc.scheduler.Kill()
}

type GetInfoCommand struct {
	scheduler JobScheduler
}

func (gic *GetInfoCommand) Execute() {
	gic.scheduler.GetInfo()
}

//============================================================

// [3] - реализуем структуру шедулера . реализуем методы  интерфеса JobScheduler
type ClusterJobScheduler struct {
	jobsRun     int
	jobsStopped int
	jobsKilled  int
}

// реализуем интерфейс шедулера
func (clSch *ClusterJobScheduler) Run() {
	clSch.jobsRun += 1
	fmt.Printf("One more job started\n")
}

func (clSch *ClusterJobScheduler) Stop() {
	clSch.jobsStopped += 1
	fmt.Printf("One more job stopped\n")
}

func (clSch *ClusterJobScheduler) Kill() {
	clSch.jobsKilled += 1
	fmt.Printf("One more job killed\n")
}

func (clSch *ClusterJobScheduler) GetInfo() {
	fmt.Printf("Running: %d , Stopped: %d, Killed: %d", clSch.jobsRun, clSch.jobsStopped, clSch.jobsKilled)
}

//============================================================

// [4] Реализуем структуру Receiver . Которая реалихзует интерйфейм команды. У нас Ресивер это ПОЛЬЗОВАТЕЛЬ, который хочет управлять задачей
// класс который содержит Бизнес Логику. Команда
type Receiver struct {
	command Command
}

func (r *Receiver) manageJob() {
	r.command.Execute()
}

//============================================================

func main() {
	// 0 - init ClusterJobScheduler
	clusteScheduler := &ClusterJobScheduler{
		jobsRun:     0,
		jobsStopped: 0,
		jobsKilled:  0,
	}

	// 1 - создаем экземпляры команд
	runCmd := &RunCommand{clusteScheduler}
	stopCmd := &StopCommand{clusteScheduler}
	killCmd := &KillCommand{clusteScheduler}
	infoCmd := &GetInfoCommand{clusteScheduler}

	// 2 - init receiver for each command
	runRec := Receiver{runCmd}
	stopRec := Receiver{stopCmd}
	killRec := Receiver{killCmd}
	infoRec := Receiver{infoCmd}

	runRec.manageJob()
	runRec.manageJob()
	stopRec.manageJob()
	stopRec.manageJob()
	killRec.manageJob()
	infoRec.manageJob()
}
