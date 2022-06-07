package main

import (
	"bufio"
	"fmt"
	"live-code-3-1/data"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Menu string
	Meja int
}

func DataHariIni(c Config) ([]data.Service, []data.Service) {
	menuFromFile := FileRead(c.Menu)
	sliceMenu := []data.Service{}
	sliceMeja := []data.Service{}
	for _, menuNya := range menuFromFile {
		menuTemp := strings.Split(menuNya, "/")
		priceTemp,_ := strconv.Atoi(menuTemp[1])
		stockTemp,_ := strconv.Atoi(menuTemp[2])
		
		menuStruct := data.Menu{
			Name: menuTemp[0],
			Price: priceTemp,
			Stock: stockTemp,
		}
		sliceMenu = append(sliceMenu, &menuStruct)
	}
	for i := 0; i < c.Meja; i++ {
		mejaStruct := data.Meja{
			Nomor: i+1,
			Availability: true,
		}
		sliceMeja = append(sliceMeja, &mejaStruct)
	}
	return sliceMenu, sliceMeja
}

func FileRead(filePath string) []string {
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines[1:]
}
