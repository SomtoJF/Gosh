package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func execInput(cmd string) error {
	input := strings.TrimSuffix(cmd, "\n")
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("path required")
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)

	}

	command := exec.Command(args[0], args[1:]...)
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout

	return command.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		wd, _ := os.Getwd()
		prompt := fmt.Sprintf("%s | gosh 🔱: ", wd)

		c := color.New(color.FgCyan).Add(color.Underline).Add(color.Bold)
		i := color.New(color.FgGreen)
		c.Println(prompt)
		i.Print("> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
