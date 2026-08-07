package tui

import (
	"fmt"

	styling "charm.land/lipgloss/v2"
)

func (m model) winScreen() string {
	msg := "YOU WON!\n" +
		fmt.Sprintf("Press %s to continue endless\n\n", m.keys.Continue.Help().Key) +

		fmt.Sprintf("%s Restart\n", m.keys.Restart.Help().Key) +
		fmt.Sprintf("%s Quit", m.keys.Quit.Help().Key)

	popupStyle := styling.NewStyle().
		Background(styling.Color("#d0df00")).
		Foreground(styling.Color("#200f24")).
		BorderForeground(styling.Color("#d0df00")).
		Border(styling.RoundedBorder()).
		Padding(1, 4)
	return popupStyle.Render(msg)
}

func (m model) defeatScreen() string {
	msg := "YOU LOST :(\n\n" +
		fmt.Sprintf("%s Restart\n", m.keys.Restart.Help().Key) +
		fmt.Sprintf("%s Quit", m.keys.Quit.Help().Key)

	popupStyle := styling.NewStyle().
		Background(styling.Color("#ce493f")).
		Foreground(styling.Color("#caffe4")).
		BorderForeground(styling.Color("#ce493f")).
		Border(styling.RoundedBorder()).
		Padding(1, 4)
	return popupStyle.Render(msg)
}

func (m model) endlessPopup() string {
	msg := "You are now entering Endless mode. \n" +
		" - Score gain is doubled.\n" +
		" - 'Undo' is not available\n" +
		" - Higher value tiles spawn.\n\n" +

		fmt.Sprintf("%s to continue", m.keys.Continue.Help().Key)

	popupStyle := styling.NewStyle().
		Background(styling.Color("#414141")).
		Foreground(styling.Color("#ffe7e7")).
		BorderForeground(styling.Color("#ac3847")).
		Border(styling.RoundedBorder()).
		Padding(1, 4)
	return popupStyle.Render(msg)
}
