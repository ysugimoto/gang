package gang

import (
	"fmt"
	"strings"
)

func GetInputOfSelectIndex(list []string) int {
	var selected int

	for {
		fmt.Printf(GREEN+"Select run command [1-%d]: "+RESET, len(list))
		Scanf("%d", &selected)
		if selected >= 0 && selected <= len(list) {
			return selected
		}
		Println("Selected command is out of range.")
	}
}

func GetInputOfYesNo(msg string) bool {
	var yn string

	fmt.Print(msg)
	for {
		Scanf("%s", &yn)
		switch strings.ToLower(yn) {
		case "y":
			return true
		case "n":
			return false
		}
		Print("Please input \"y\" or \"n\": ")
	}
}
