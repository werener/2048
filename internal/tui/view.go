package tui

import (
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

	// account for help menu, score and margin
	gridSize := availableSquare - 2 - 1 - (2 * margin)

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
	numTiles := int(m.grid.Size)
	rows := make([]string, numTiles)

	// fit to the amount of Tiles
	size = size / 8 * 8

	tileSize := size / numTiles

	for i := range numTiles {
		tiles := make([]string, numTiles)
		for j := range numTiles {
			tiles[j] = m.TileView(m.grid.Tiles[i][j], tileSize)
		}
		rows[i] = lipgloss.JoinHorizontal(lipgloss.Top, tiles...)
	}

	grid := lipgloss.JoinVertical(lipgloss.Left, rows...)

	return lipgloss.NewStyle().
		Height(size).
		Width(size * 2).
		Align(lipgloss.Center).
		BorderStyle(lipgloss.NormalBorder()).
		Render(grid)
}

func (m model) TileView(tile core.Tile, size int) string {
	repr := tile.Repr()

	bg, fg := colorTile(tile)
	style := lipgloss.NewStyle().
		Height(size).Width(size*2 - 1).
		AlignHorizontal(lipgloss.Center).AlignVertical(lipgloss.Center).
		Background(bg).Foreground(fg)

	return style.Render(repr)
}

func colorTile(tile core.Tile) (bg color.Color, fg color.Color) {
	switch tile.Value {
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
	case 32:
		fg = lipgloss.Color("#aeffc0")
		bg = lipgloss.Color("#236d39")
	case 64:
		fg = lipgloss.Color("#d8f5a3")
		bg = lipgloss.Color("#7f9726")
	case 128:
		fg = lipgloss.Color("#ffd5a6")
		bg = lipgloss.Color("#bd9100")
	case 256:
		fg = lipgloss.Color("#ffc582")
		bg = lipgloss.Color("#c25700")
	case 512:
		fg = lipgloss.Color("#ff9182")
		bg = lipgloss.Color("#8f0c0c")
	case 1024:
		fg = lipgloss.Color("#ffadc6")
		bg = lipgloss.Color("#770059")
	case 2048:
		fg = lipgloss.Color("#f9ffc2")
		bg = lipgloss.Color("#55007c")
	}
	return
}
