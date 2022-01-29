package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	items := []list.Item{
		item{title: "Raspberry Pi's", description: "Test description", videoId: "STsI6IbPbGQ"},
		item{title: "Nutella", description: "Test description", videoId: "JY-q3tNSCNw"},
		item{title: "Bitter melon", description: "Test description", videoId: "K10tdjEGbJg"},
		item{title: "Chillhop Radio", description: "Jazzy & Lo-Fi hip hop beats", videoId: "5yx6BWlEVcY"},
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = titleStyle.Render(" ►  wotchy.fyi")

	if err := tea.NewProgram(m, tea.WithAltScreen()).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
