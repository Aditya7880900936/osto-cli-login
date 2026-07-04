package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

// Reader is a shared buffered reader
// used for reading user input from the terminal.
var Reader = bufio.NewReader(os.Stdin)

// ReadLine displays a prompt
// and returns the user's input.
func ReadLine(prompt string) string {
	fmt.Print(prompt)

	input, _ := Reader.ReadString('\n')

	return strings.TrimSpace(input)
}

// ReadPassword securely reads a password
// without displaying it on the terminal.
func ReadPassword(prompt string) string {
	fmt.Print(prompt)

	password, _ := term.ReadPassword(int(os.Stdin.Fd()))

	fmt.Println()

	return strings.TrimSpace(string(password))
}
