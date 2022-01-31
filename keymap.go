package main

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	playVideo   key.Binding
	playAudio   key.Binding
	togglePause key.Binding
	stopPlaying key.Binding
}

// Additional short help entries.
// This satisfies the help.KeyMap interface and is entirely optional.
func (d keyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		d.playVideo,
		d.playAudio,
		d.stopPlaying,
	}
}

// Additional full help entries.
// This satisfies the help.KeyMap interface and is entirely optional.
func (d keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			d.playVideo,
			d.playAudio,
			d.togglePause,
			d.stopPlaying,
		},
	}
}

func newKeyMap() *keyMap {
	return &keyMap{
		playVideo: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "play video"),
		),
		playAudio: key.NewBinding(
			key.WithKeys(" "),
			key.WithHelp("space", "play audio only"),
		),
		togglePause: key.NewBinding(
			key.WithKeys("p"),
			key.WithHelp("p", "pause playback"),
		),
		stopPlaying: key.NewBinding(
			key.WithKeys("w"),
			key.WithHelp("w", "stop playback"),
		),
	}
}
