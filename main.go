package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

var nameCount = map[string]int{}

func main() {

	files, err := ioutil.ReadDir("./")

	if err != nil {
		logError(err)
	}

	for _, f := range files {
		fileName := f.Name()
		if strings.HasSuffix(fileName, ".html") {
			processFile(fileName)
		}
	}

	output := []string{}

	for key, val := range nameCount {
		if !strings.HasPrefix(key, "The account has disconnected") {
			output = append(output, key+" "+strconv.Itoa(val))
		}
	}

	err = ioutil.WriteFile("output.txt", []byte(strings.Join(output, "\n")), 0755)

	if err != nil {
		logError(err)
	}
}

func processFile(name string) {

	file, err := ioutil.ReadFile(name)

	if err != nil {
		logError(err)
	}

	fileContent := string(file)

	lines := strings.Split(fileContent, "\n")

	for _, line := range lines {
		i := strings.Split(line, "<b>")

		if len(i) > 1 {
			suf := strings.Split(i[1], "</b>")
			if len(suf) > 1 {
				out := strings.TrimSpace(suf[0])
				out = strings.Trim(out, "\n")
				nameCount[out]++
			}
		}
	}
}

func logError(err error) {
	log.Println(err)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
