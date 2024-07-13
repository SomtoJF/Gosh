package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/chzyer/readline"
	"github.com/fatih/color"
	"github.com/somtojf/gosh/initializers"
	"github.com/somtojf/gosh/migrate"
	"github.com/somtojf/gosh/models"
)

func init() {
	initializers.ConnectToDb()
}

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
	case "history":
		var history []models.History

		result := initializers.DB.Find(&history)

		if result.Error != nil {
			return result.Error
		}

		for _, historyCommand := range history {
			fmt.Printf("%d\t%s", historyCommand.ID, historyCommand.Command)
		}
		return nil
	}

	command := exec.Command(args[0], args[1:]...)
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout

	return command.Run()
}

func main() {
	migrate.Migrate()

	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		wd, _ := os.Getwd()
		prompt := fmt.Sprintf("%s | gosh 🔱: ", wd)

		c := color.New(color.FgCyan).Add(color.Underline).Add(color.Bold)
		c.Println(prompt)

		input, err := rl.Readline()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err == nil {
			post := models.History{Command: input}
			result := initializers.DB.Create(&post)

			if result.Error != nil {
				color.Red("could not persist command to history")
			}
		}
	}
}
