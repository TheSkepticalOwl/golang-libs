package readinput

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadFloat reads a float64 value and outputs it
func ReadFloat() float64 {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		in := scanner.Text()

		n, err := strconv.ParseFloat(in, 64)

		if err == nil {
			return n
		} else {
			fmt.Println("ERROR:", err)
			fmt.Print("\nPlease enter a valid number: ")
		}
	}
}

// ReadString reads a string and outputs it
func ReadString() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	return text
}
