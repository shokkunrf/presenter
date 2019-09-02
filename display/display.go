package display

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

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
	fmt.Println(displays)
}
