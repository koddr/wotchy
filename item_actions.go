package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type item struct {
	title, description, videoId string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.description }
func (i item) FilterValue() string { return i.title }

func newItemDelegate(keys *keyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var title, videoId, status string

		if i, ok := m.SelectedItem().(item); ok {
			title = i.title
			videoId = i.videoId
		} else {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.playVideo), key.Matches(msg, keys.playAudio):
				go startPlaying(videoId, key.Matches(msg, keys.playAudio)) // goroutine
				status = fmt.Sprintf("%s  '%s' playback has started", tilde, title)
				return m.NewStatusMessage(listStatusMessageStyle(status))

			case key.Matches(msg, keys.togglePause):
				go togglePause() // goroutine
				status = fmt.Sprintf("%s  playback has been paused/resumed", tilde)
				return m.NewStatusMessage(listStatusMessageStyle(status))

			case key.Matches(msg, keys.stopPlaying):
				go stopPlaying() // goroutine
				status = fmt.Sprintf("%s  ok, playback has been stopped", tilde)
				return m.NewStatusMessage(listStatusMessageStyle(status))
			}
		}

		return nil
	}

	// Show short help at the start.
	d.ShortHelpFunc = func() []key.Binding {
		return []key.Binding{
			keys.playVideo,
			keys.playAudio,
			keys.stopPlaying,
		}
	}

	// Show full help (by clicking `?` key).
	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{
			{
				keys.playVideo,
				keys.playAudio,
				keys.togglePause,
				keys.stopPlaying,
			},
		}
	}

	return d
}
