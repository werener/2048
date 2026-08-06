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

	// account for help menu, score and margin
	gridSize := availableSquare - 2 - 2 - (2 * margin)

	// for adding gaps and padding
	grid := m.gridView(gridSize)

	help := lipgloss.NewStyle().
		Height(2).
		Render(m.help.View(m.keys))

	score := lipgloss.NewStyle().
		Height(2).
		Render(fmt.Sprintf("Your score: %d", m.game.Score()))
	view := window.Render(grid + "\n" + score + "\n" + help)

	v := tea.NewView(view)
	v.AltScreen = true
	return v
}

func (m model) gridView(size int) string {
	numTiles := int(m.game.Grid().Size)
	rows := make([]string, numTiles)

	// fit to the amount of Tiles
	size = size / 8 * 8

	tileSize := size / numTiles

	for i := range numTiles {
		tiles := make([]string, numTiles)
		for j := range numTiles {
			tiles[j] = m.tileView(m.game.Grid().Tiles[i][j], tileSize)
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

func (m model) tileView(tile core.Tile, size int) string {
	repr := tile.Repr()

	bg, fg, border, newlySpawnedBorder := colorTile(tile)
	style := lipgloss.NewStyle().
		// alignment
		Height(size).Width(size*2 - 1).
		AlignHorizontal(lipgloss.Center).AlignVertical(lipgloss.Center).
		// colors
		Background(bg).Foreground(fg).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(border).BorderBackground(bg)
	// show if a tile has just been spawned
	if tile.NewlySpawned {
		style = style.BorderStyle(lipgloss.ASCIIBorder()).BorderForeground(newlySpawnedBorder)
	}
	// if a tile is empty, don't show a border
	if tile.IsEmpty() {
		style = style.UnsetBorderStyle()
	}

	return style.Render(repr)
}

func colorTile(tile core.Tile) (bg color.Color, fg color.Color, border color.Color, newlySpawnedBorder color.Color) {
	switch tile.Value {
	case 2:
		bg = lipgloss.Color("#2c2c79")
	case 4:
		bg = lipgloss.Color("#286994")
	case 8:
		bg = lipgloss.Color("#438594")
	case 16:
		bg = lipgloss.Color("#328570")
	case 32:
		bg = lipgloss.Color("#236d39")
	case 64:
		bg = lipgloss.Color("#7f9726")
	case 128:
		bg = lipgloss.Color("#bd9100")
	case 256:
		bg = lipgloss.Color("#c25700")
	case 512:
		bg = lipgloss.Color("#8f0c0c")
	case 1024:
		bg = lipgloss.Color("#770059")
	case 2048:
		bg = lipgloss.Color("#55007c")
	}

	fg = lipgloss.Lighten(bg, 0.5)
	border = lipgloss.Darken(bg, 0.15)
	newlySpawnedBorder = lipgloss.Lighten(border, 0.3)
	newlySpawnedBorder = lipgloss.Lighten(border, 0.3)
	return
}
