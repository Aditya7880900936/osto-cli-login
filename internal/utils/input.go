package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

var Reader = bufio.NewReader(os.Stdin)

func ReadLine(prompt string) string {
	fmt.Print(prompt)

	input, _ := Reader.ReadString('\n')

	return strings.TrimSpace(input)
}

func ReadPassword(prompt string) string {
	fmt.Print(prompt)

	password, _ := term.ReadPassword(int(os.Stdin.Fd()))

	fmt.Println()

	return strings.TrimSpace(string(password))
}