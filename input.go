package gang

import (
	"fmt"
	"strings"
)

func GetInputOfSelectIndex(list []string) int {
	var selected int

	for {
		fmt.Printf("Select run command [1-%d]: ", len(list))
		fmt.Scanf("%d", &selected)
		if selected >= 1 && selected <= len(list) {
			return selected
		}
		fmt.Println("Selected command is out of range.")
	}
}

func GetInputOfYesNo(msg string) bool {
	var yn string

	fmt.Print(msg)
	for {
		fmt.Scanf("%s", &yn)
		switch strings.ToLower(yn) {
		case "y":
			return true
		case "n":
			return false
		}
		fmt.Print("Please input \"y\" or \"n\": ")
	}
}
