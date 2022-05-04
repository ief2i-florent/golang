package multithread

import (
	"fmt"
	"runtime"
)

var nbOfCPU = runtime.NumCPU()
var workers []*Worker = make([]*Worker, nbOfCPU)
var nbOfMission int = 0
var nbOfMissionInProgress int = 0
var InProgress chan int = make(chan int)

type Worker struct {
	isRunning bool
	missions  []*Mission
}

func init() {
	for wn, w := range workers {
		if w == nil {
			workers[wn] = &Worker{
				isRunning: false,
				missions:  make([]*Mission, 0),
			}
		}
	}
}

func AddInQueue(m *Mission) {
	id := nbOfMission % nbOfCPU
	worker := workers[id]
	fmt.Println("Add Mission", m, id)
	worker.addMission(m)
	nbOfMission++
}

func Run() {
	for _, w := range workers {
		go w.executeAllMissions(InProgress)
	}

	for c := range InProgress {
		nbOfMissionInProgress = nbOfMissionInProgress + c
		if nbOfMissionInProgress == nbOfMission {
			break
		}
	}

	nbOfMission = 0
}

func (w *Worker) executeAllMissions(c chan int) {
	for _, mission := range w.missions {
		c <- mission.execute()
	}

	w.missions = make([]*Mission, 0)
}

func (w *Worker) addMission(m *Mission) {
	w.missions = append(w.missions, m)
}
