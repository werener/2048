package tui

import (
	"fmt"
	"image/color"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/werener/2048.git/internal/core"
)

func (m model) View() tea.View {
	margin := 2

	window := lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Margin(margin).
		Align(lipgloss.Center)

	availableSquare := min(m.width/2, m.height)

	// account for help menu and margin
	gridSize := availableSquare - 2 - (2 * margin)

	// for adding gaps and padding
	grid := m.gridView(gridSize)

	help := lipgloss.NewStyle().
		Height(2).
		Render(m.help.View(m.keys))

	view := window.Render(grid + "\n" + help)

	v := tea.NewView(view)
	v.AltScreen = true
	return v
}

func (m model) gridView(size int) string {
	numCells := int(m.grid.Size)
	rows := make([]string, numCells)

	// fit to the amount of Cells
	size = size / 8 * 8

	cellSize := size / numCells

	for i := range numCells {
		cells := make([]string, numCells)
		for j := range numCells {
			cells[j] = m.CellView(m.grid.Cells[i][j], cellSize)
		}
		rows[i] = lipgloss.JoinHorizontal(lipgloss.Top, cells...)
	}

	grid := lipgloss.JoinVertical(lipgloss.Left, rows...)

	return lipgloss.NewStyle().
		Height(size).
		Width(size * 2).
		Align(lipgloss.Center).
		BorderStyle(lipgloss.NormalBorder()).
		Render(grid)
}

func (m model) CellView(cell core.Cell, size int) string {

	repr, fg, bg := colorCell(cell)

	style := lipgloss.NewStyle().Height(size).Width(size*2 - 1)

	style = style.BorderStyle(lipgloss.NormalBorder()).
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Center).
		BorderForeground(fg).
		BorderBackground(bg).
		Background(bg)
	return style.Render(repr)

}

func colorCell(cell core.Cell) (repr string, fg color.Color, bg color.Color) {
	if cell != core.EMPTY {
		repr = fmt.Sprintf("%d", cell)
	}
	switch cell {
	case 2:
		fg = lipgloss.Color("#9f9fc7")
		bg = lipgloss.Color("#2c2c79")
	case 4:
		fg = lipgloss.Color("#9d9dff")
		bg = lipgloss.Color("#444483")
	case 8:
		fg = lipgloss.Color("#7fafe7")
		bg = lipgloss.Color("#438594")
	case 16:
		fg = lipgloss.Color("#b4f8e7")
		bg = lipgloss.Color("#328570")
	}
	return
}
