package main

import gloss "github.com/charmbracelet/lipgloss"

var (
	// Basic colors:
	white = gloss.AdaptiveColor{Light: "#FAFAFA", Dark: "#FAFAFA"}
	black = gloss.AdaptiveColor{Light: "#333333", Dark: "#333333"}

	// Additional colors:
	nordic  = gloss.AdaptiveColor{Light: "#475569", Dark: "#94A3B8"}
	stone   = gloss.AdaptiveColor{Light: "#57534E", Dark: "#A8A29E"}
	emerald = gloss.AdaptiveColor{Light: "#059669", Dark: "#34D399"}
	violet  = gloss.AdaptiveColor{Light: "#7C3AED", Dark: "#A78BFA"}
	rose    = gloss.AdaptiveColor{Light: "#E11D48", Dark: "#FB7185"}
)

var (
	// Main styles:
	appStyle = gloss.NewStyle().Padding(1, 2)

	// List styles:
	listTitleStyle         = gloss.NewStyle().Foreground(rose).UnsetPadding()
	listStatusMessageStyle = gloss.NewStyle().Foreground(stone).Render
)

var (
	// TUI elements:
	checkMark = gloss.NewStyle().SetString("✓").Foreground(emerald).String()
	tilde     = gloss.NewStyle().SetString("~").Foreground(emerald).String()
)
