package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	loadedList   list.Model
	delegateKeys *keyMap
}

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		topGap, rightGap, bottomGap, leftGap := appStyle.GetPadding()
		m.loadedList.SetSize(msg.Width-leftGap-rightGap, msg.Height-topGap-bottomGap)

	case tea.KeyMsg:
		// Don't match any of the keys below if we're actively filtering.
		if m.loadedList.FilterState() == list.Filtering {
			break
		}
	}

	// This will also call our delegate's update function.
	newListModel, cmd := m.loadedList.Update(msg)
	m.loadedList = newListModel
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return appStyle.Render(m.loadedList.View())
}
