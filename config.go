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

func DataHariIni(c Config) ([]data.Menu, []data.Meja) {
	menuFromFile := FileRead(c.Menu)
	sliceMenu := []data.Menu{}
	sliceMeja := []data.Meja{}
	for _, menuNya := range menuFromFile {
		menuTemp := strings.Split(menuNya, "/")
		priceTemp,_ := strconv.Atoi(menuTemp[1])
		stockTemp,_ := strconv.Atoi(menuTemp[2])
		
		menuStruct := data.Menu{
			Name: menuTemp[0],
			Price: priceTemp,
			Stock: stockTemp,
		}

		sliceMenu = append(sliceMenu, menuStruct)
	}
	return sliceMenu, sliceMeja
}

// func DataHariIni(c Config) (string, int) {
// 	FileRead(c.Menu)
// 	return c.Menu, c.Meja
// }

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
