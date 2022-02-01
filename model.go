package main

import (
	"time"

	"github.com/charmbracelet/bubbles/list"
)

func newModel() model {
	// Define options.
	var (
		delegateKeys          = newKeyMap()
		statusMessageLifetime = 3 * time.Second
		items                 = []list.Item{
			item{title: "Lofi Girl", description: "lofi hip hop radio - beats to sleep/chill to", videoId: "DWcJFNfaw9c"},
			item{title: "Chillhop Music", description: "Chillhop Radio - jazzy & lofi hip hop beats 🐾", videoId: "5yx6BWlEVcY"},
			item{title: "Monstafluff Music", description: "24/7 lofi hip hop radio - beats to relax/study to", videoId: "zVqJv_dKUEs"},
		}
	)

	// Create delegate with a new keymap.
	delegate := newItemDelegate(delegateKeys)

	// Setup list of loaded items.
	loadedList := list.New(items, delegate, 0, 0)
	loadedList.Title = "📺 wotchy"
	loadedList.StatusMessageLifetime = statusMessageLifetime

	// Set styles for the list itself.
	loadedList.Styles.Title = listTitleStyle
	loadedList.Styles.FilterPrompt = loadedList.Styles.FilterPrompt.Foreground(stone)
	loadedList.Styles.FilterCursor = loadedList.Styles.FilterCursor.Foreground(rose)

	// Set styles for each item in the list.
	delegate.Styles.NormalTitle = delegate.Styles.SelectedTitle.Foreground(stone)
	delegate.Styles.NormalDesc = delegate.Styles.SelectedDesc.Foreground(stone)
	delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.Foreground(rose).BorderLeftForeground(rose)
	delegate.Styles.SelectedDesc = delegate.Styles.SelectedDesc.Foreground(stone).BorderLeftForeground(rose)

	return model{
		loadedList:   loadedList,
		delegateKeys: delegateKeys,
	}
}
