package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	gloss "github.com/charmbracelet/lipgloss"
)

var (
	special    = gloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	appStyle   = gloss.NewStyle().Margin(1, 2)
	titleStyle = gloss.NewStyle().BorderStyle(gloss.RoundedBorder()).Padding(0, 1)
	inputStyle = gloss.NewStyle().Foreground(gloss.Color("62"))
	checkMark  = gloss.NewStyle().SetString("✓").Foreground(special).String()
)

type item struct {
	title, description, videoId string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.description }
func (i item) FilterValue() string { return i.title }

type model struct {
	list        list.Model
	selected    string
	isAudioOnly bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			if m.selected == "" {
				if it, ok := m.list.SelectedItem().(item); ok {
					m.selected = string(it.videoId)
					m.isAudioOnly = msg.String() == " "
				}
			}
		case "w":
			m.selected = ""
			stopPlaying()
			return m, nil
		default:
			m.selected = ""
		}
	case tea.WindowSizeMsg:
		top, right, bottom, left := appStyle.GetMargin()
		m.list.SetSize(msg.Width-left-right, msg.Height-top-bottom)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m model) View() string {
	var state string

	if m.selected != "" {
		go startPlaying(m, m.isAudioOnly)
		state += "Please wait for mpv...\n\n"
		state += fmt.Sprintf("%s Now playing video %s (pid %d)", checkMark, inputStyle.Render(m.selected), <-PID)
		state += "\n\nPress any key to exit to the previous list"
	} else {
		state += m.list.View()
	}

	return appStyle.Render(state)
}
