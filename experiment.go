package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"sync"
	"time"
)

type Experiment struct {
	totalContainers int
	containerRatio  map[string]int
	Result          map[string]float32
	requestOrder    []string
}

func (e *Experiment) Init(totalContainers int, containerRatio map[string]int) {
	e.totalContainers = totalContainers
	e.containerRatio = containerRatio
	var requestOrder []string

	for runtime, ratio := range containerRatio {
		// 해당 컨테이너 런타임으로 실행시켜야 하는 컨테이너 개수
		num := e.totalContainers * ratio / 100
		for i := 0; i < num; i++ {
			requestOrder = append(requestOrder, runtime)
		}
	}
	rand.Shuffle(e.totalContainers, func(i, j int) {
		requestOrder[i], requestOrder[j] = requestOrder[j], requestOrder[i]
	})
	e.requestOrder = requestOrder
}

func (e *Experiment) Run() {
	var ch = make(chan time.Duration, e.totalContainers)
	var wg sync.WaitGroup

	for i, runtime := range e.requestOrder {
		wg.Add(1)
		go func(workerId int, runtime string) {
			defer wg.Done()
			runContainer(ch, runtime)
		}(i, runtime)
	}

	wg.Wait()

	// 모든 컨테이너 실행 완료

	var avg time.Duration
	for j := 1; j <= e.totalContainers; j++ {
		avg += <-ch
	}
	avg /= time.Duration(e.totalContainers)
	fmt.Println("평균 실행시간:", avg)
}

// 입력 받은 런타임으로 컨테이너를 실행시키고 그 실행 시간을 출력한다
func runContainer(ch chan time.Duration, runtime string) {
	startTime := time.Now()
	cmd := exec.Command("./script.sh", runtime)
	var _, err = cmd.Output()

	if err != nil {
		fmt.Println(err)
	}

	elapsedTime := time.Since(startTime)
	//fmt.Println(elapsedTime)
	ch <- elapsedTime
}
