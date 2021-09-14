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
	requestOrder    []string
}

type elapsedTime struct {
	runtime  string
	duration time.Duration
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

func (e *Experiment) Run() map[string]time.Duration {
	var ch = make(chan elapsedTime, e.totalContainers)
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

	// 각 컨테이너 런타임별 컨테이너 실행시간 평균값 계산
	temp := make(map[string]time.Duration)
	result := make(map[string]time.Duration)
	for i := 1; i <= e.totalContainers; i++ {
		t := <-ch
		temp[t.runtime] += t.duration
	}
	for runtime, duration := range temp {
		result[runtime] = duration / time.Duration(e.totalContainers*e.containerRatio[runtime]/100)
	}

	return result
}

// 입력 받은 런타임으로 컨테이너를 실행시키고 그 실행 시간을 출력한다
func runContainer(ch chan elapsedTime, runtime string) {
	startTime := time.Now()
	cmd := exec.Command("./script.sh", runtime)
	var _, err = cmd.Output()

	if err != nil {
		fmt.Println(err)
	}

	duration := time.Since(startTime)

	ch <- elapsedTime{runtime: runtime, duration: duration}
}
