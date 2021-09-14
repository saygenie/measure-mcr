package main

import (
	"./experiment"
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var totalContainersList = []int{10}
	var containerRatioList = readDataFromCSV("container-ratio.csv")

	for _, totalContainers := range totalContainersList {
		for j, containerRatio := range containerRatioList {
			fmt.Printf("[총 컨테이너 %d개] %d번째 실험 시작\n", totalContainers, j + 1)
			fmt.Println("containerRatio:", containerRatio)
			exp := experiment.Experiment{}
			exp.Init(totalContainers, containerRatio)
			exp.Run()
			fmt.Println(exp.Result)
			fmt.Printf("[총 컨테이너 %d개] %d번째 실험 종료\n", totalContainers, j + 1)
		}

	}
}

func readDataFromCSV(csvFile string) []map[string]int {
	var containerRatioList []map[string]int
	file, _ := os.Open(csvFile)
	rdr := csv.NewReader(bufio.NewReader(file))
	rows, _ := rdr.ReadAll()
	containerRuntimes := rows[0]

	for i, row := range rows {
		if i == 0 {
			continue
		}
		containerRatio := make(map[string]int)
		for j, containerRuntime := range containerRuntimes {
			ratio, _ := strconv.Atoi(row[j])
			containerRatio[containerRuntime] = ratio
		}
		containerRatioList = append(containerRatioList, containerRatio)
	}

	return containerRatioList
}
