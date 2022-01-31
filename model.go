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

	// Setup list of loaded items.
	delegate := newItemDelegate(delegateKeys)
	loadedList := list.New(items, delegate, 0, 0)
	loadedList.Title = "📺 wotchy"
	loadedList.StatusMessageLifetime = statusMessageLifetime
	loadedList.Styles.Title = listTitleStyle

	return model{
		loadedList:   loadedList,
		delegateKeys: delegateKeys,
	}
}
