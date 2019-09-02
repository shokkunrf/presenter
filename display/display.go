package display

import (
	"fmt"
)

func Status() {
	displays := getDisplay()
	fmt.Println(displays)
}
