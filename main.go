package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	file = flag.String("file", "input.csv", "csv file path")
)

func ReadCsvFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
		return nil, err
	}
	defer f.Close()
	var record [][]string
	scanner := bufio.NewScanner(f)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		if string(line[0]) == "#" {
			continue
		} else {
			record = append(record, strings.Split(line, ","))
		}
	}

	return record, nil
}

func CalculateEV(actQuan, plannedQuan, plannedMoney float64) (result float64) {
	return actQuan / plannedQuan * plannedMoney
}

func CostVariance(ev, ac float64) (result float64) {
	return ev - ac
}

func ScheduleVariance(ev, pv float64) (result float64) {
	return ev - pv
}

func CostPerformanaceIndex(ev, ac float64) (result float64) {
	return ev / ac
}

func SchedulePerformanceIndex(ev, pv float64) (result float64) {
	return ev / pv
}

func EstimateAtCompletion(bac, cpi float64) (result float64) {
	return bac / cpi
}

func EstimatedTimeToComplete(ote, spi float64) (result float64) {
	return ote / spi
}

func main() {
	flag.Parse()

	data, err := ReadCsvFile(*file)
	if err != nil {
		panic(err)
	}

	var rawData [][]float64
	for i := 0; i < len(data); i++ {
		var row []float64
		for j := 0; j < len(data[i]); j++ {
			s, err := strconv.ParseFloat(data[i][j], 64)
			if err != nil {
				panic(err)
			}
			row = append(row, s)
		}
		rawData = append(rawData, row)
	}

	for i := 0; i < len(rawData); i++ {
		fmt.Println("Day ", i+1)

		fmt.Println(rawData[i])
		ev := CalculateEV(rawData[i][2], rawData[i][0], rawData[i][1])
		fmt.Println("Earned Value: ")
		fmt.Println(ev)

		cv := CostVariance(ev, rawData[i][3])
		fmt.Println("Cost Variance: ")
		fmt.Println(cv)

		sv := ScheduleVariance(ev, rawData[i][1])
		fmt.Println("Schedule Variance: ")
		fmt.Println(sv)

		cpi := CostPerformanaceIndex(ev, rawData[i][3])
		fmt.Println("Cost Performanace Index: ")
		fmt.Println(cpi)

		spi := SchedulePerformanceIndex(ev, rawData[i][1])
		fmt.Println("Schedule Performance Index: ")
		fmt.Println(spi)

		etc := EstimatedTimeToComplete(float64(len(rawData)), spi)
		fmt.Println("Estimated Time To Complete: ")
		fmt.Println(etc)

		fmt.Println("---------------------------------")
	}
}
