package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	setGOMAXPROCS()
	splitLineCount := 128
	grepString, filePath, processingStartNumberOfLine := getArgsValue()
	lineCount := startScan(grepString, filePath, processingStartNumberOfLine, splitLineCount)
	if processingStartNumberOfLine != 1 {
		fmt.Println("GoGrep finished. All Line: " + strconv.Itoa(lineCount) + " Check Start Line" + strconv.Itoa(processingStartNumberOfLine))
	}
}

func startScan(grepString string, filePath string, processingStartNumberOfLine int, splitLineCount int) int {
	file, e := os.Open(filePath)
	checkError(e)
	defer file.Close()
	lineCount := 1
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // 最終行まで１行づつ読みだす
		if lineCount >= processingStartNumberOfLine {
			lines = append(lines, scanner.Text())
			// 128行で１セットとしてlines配列にぶっ込んでgoroutineで処理させる
			// ほんでlinesを初期化
			if lineCount%splitLineCount == 0 {
				if lineCount != 0 {
					checkLine(lines, grepString)
					lines = []string{}
				}
			}
		}
		lineCount++
	}
	checkLine(lines, grepString)
	return lineCount
}

func checkLine(lines []string, grepString string) {
	channel := make(chan int)
	// 無名goroutine関数即時実行
	go func(sendChannel chan<- int) {
		for i := range lines {
			if strings.Index(lines[i], grepString) > -1 {
				// fmt.Println(lines[i])
			}
		}
		close(sendChannel)
	}(channel)

	for {
		ok := <-channel
		if ok == 0 {
			break
		}
	}
}

func getArgsValue() (string, string, int) {
	processingStartNumberOfLineString := "1"
	grepString := ""
	filePath := ""

	flag.Parse()
	args := flag.Args()

	if len(args) == 2 {
		grepString = args[0]
		filePath = args[1]
	} else if len(args) == 3 {
		grepString = args[0]
		filePath = args[1]
		processingStartNumberOfLineString = args[2]
	} else {
		fmt.Println("args error: need 2 or 3 value. grepString, filePath, processingStartNumberOfLine")
		os.Exit(1)
	}

	processingStartNumberOfLine, e := strconv.Atoi(processingStartNumberOfLineString)
	if e != nil {
		fmt.Println("args error: processingStartNumberOfLine is not number.")
	}
	return grepString, filePath, processingStartNumberOfLine
}

func setGOMAXPROCS() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
