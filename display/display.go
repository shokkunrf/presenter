package display

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

type display struct {
	name        string
	sizes       []string
	isConnected bool
	isPrimary   bool
}

func getDisplayInfo(line string) display {
	dp := display{}

	words := strings.Split(line, " ")
	dp.name = words[0]

	if words[1] == "connected" {
		dp.isConnected = true
		if words[2] == "primary" {
			dp.isPrimary = true
		}
	}
	return dp
}

func getDisplaySize(line string) string {
	rep := regexp.MustCompile(`\d+x\d+i?`)
	size := rep.FindAllStringSubmatch(line, -1)

	return size[0][0]
}

func getDisplay() []display {
	out, err := exec.Command("xrandr").Output()
	if err != nil {
		log.Fatalln(err)
	}

	text := strings.Split(string(out), "\n")

	dps := []display{}
	for i := 0; i < len(text); i++ {
		rep := regexp.MustCompile(`connected`)
		if rep.MatchString(text[i]) {
			dp := display{}

			dp = getDisplayInfo(text[i])
			sizes := []string{}

			for j := i + 1; j < len(text); j++ {
				rep = regexp.MustCompile(`\s*\d+x\d+i?\s`)
				if rep.MatchString(text[j]) {
					size := getDisplaySize(text[j])
					sizes = append(sizes, size)
					continue
				}
				i = j - 1
				break
			}
			dp.sizes = sizes
			dps = append(dps, dp)
		}
	}
	return dps
}

func Status() {
	displays := getDisplay()
	for _, display := range displays {
		status := "\x1b[31m x \x1b[0m"
		if display.isConnected {
			status = "\x1b[32m o \x1b[0m"
		}
		msg := display.name + "\t[" + status + "]"
		fmt.Println(msg)
	}
}

func Mirroring() {
	displays := getDisplay()
	primary := display{}
	for _, display := range displays {
		if display.isPrimary == true {
			primary = display
		}
	}

	for _, display := range displays {
		if display.isPrimary == true {
			continue
		}
		exec.Command("xrandr --output " + display.name + " --same-as " + primary.name)
	}
}
