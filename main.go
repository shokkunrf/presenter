package main

import (
	"flag"
	"presenter/display"
)

func main() {
	flag.Parse()
	subCmd := flag.Arg(0)

	switch subCmd {
	case "status":
		display.Status()
		break
	default:
		break
	}
}
