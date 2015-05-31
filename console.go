package gang

import (
	"fmt"
	"github.com/mgutz/ansi"
)

var (
	RESET    = ansi.ColorCode("reset")
	RED      = ansi.ColorCode("red")
	REDB     = ansi.ColorCode("red+b")
	GREEN    = ansi.ColorCode("green")
	GREENB   = ansi.ColorCode("green+b")
	MAGENTA  = ansi.ColorCode("magenta")
	MAGENTAB = ansi.ColorCode("magenta+b")
	CYAN     = ansi.ColorCode("cyan")
	CYANB    = ansi.ColorCode("cyan+b")

	GANG = fmt.Sprint(MAGENTAB, "[gang] ", RESET)
)

func EPrintf(format string, args ...interface{}) {
	fmt.Print(GANG, RED, fmt.Sprintf(format, args...), RESET)
}

func EPrintln(message string) {
	fmt.Println(GANG, GREEN, message, RESET)
}

func EPrint(message string) {
	fmt.Print(GANG, RED, message, RESET)
}

func ESprintf(format string, args ...interface{}) string {
	return fmt.Sprint(GANG, RED, fmt.Sprintf(format, args...), RESET)
}

func ESprint(message string) string {
	return fmt.Sprint(GANG, RED, message, RESET)
}

func Printf(format string, args ...interface{}) {
	fmt.Printf(GANG+format, args...)
}

func Println(message string) {
	fmt.Println(GANG, message)
}

func Print(message string) {
	fmt.Print(GANG, message)
}

func QPrintf(format string, args ...interface{}) {
	fmt.Print(GANG, GREEN, fmt.Sprintf(format, args...), RESET)
}

func QPrintln(message string) {
	fmt.Println(GANG, GREEN, message, RESET)
}

func QPrint(message string) {
	fmt.Print(GANG, GREEN, message, RESET)
}

func QSprintf(format string, args ...interface{}) string {
	return fmt.Sprint(GANG, GREEN, fmt.Sprintf(format, args...), RESET)
}

func QSprint(message string) string {
	return fmt.Sprint(GANG, GREEN, message, RESET)
}

func Sprintf(format string, args ...interface{}) string {
	return fmt.Sprintf(GANG+format, args...)
}

func Sprint(message string) string {
	return fmt.Sprint(GANG, message)
}

func Scanf(format string, capture ...interface{}) (int, error) {
	return fmt.Scanf(format, capture...)
}
