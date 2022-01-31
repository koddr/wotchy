package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if err := tea.NewProgram(newModel()).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
