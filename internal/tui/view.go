package tui

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m model) View() tea.View {
	margin := 2
	window := lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Margin(margin).
		Align(lipgloss.Center)

	num_cells := int(m.grid.Size)
	// to evenly fit the cells and square it
	rows := min(
		m.height/num_cells*num_cells,
		m.width/num_cells*num_cells/2,
	)
	// account for help menu and margin
	rows -= 2 + 2*margin
	cols := 2 * rows

	grid := lipgloss.NewStyle().
		Height(rows).
		Width(cols).
		Align(lipgloss.Center).
		BorderStyle(lipgloss.NormalBorder()).
		Render(m.gridView(rows, cols))

	help := lipgloss.NewStyle().
		Height(2).
		Render(m.help.View(m.keys))

	view := window.Render(grid + "\n" + help)

	v := tea.NewView(view)
	v.AltScreen = true
	return v
}

func (m model) gridView(h int, w int) string {
	return m.grid.Repr()
}
