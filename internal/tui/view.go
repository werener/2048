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
		Render(fmt.Sprintf("Your score: %d", m.score))
	view := window.Render(grid + "\n" + score + "\n" + help)

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
			tiles[j] = m.tileView(m.grid.Tiles[i][j], tileSize)
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
		fg = lipgloss.Color("#9f9fc7")
		bg = lipgloss.Color("#2c2c79")
		border = lipgloss.Color("#4545ad")
		newlySpawnedBorder = lipgloss.Color("#5c5caa")
	case 4:
		fg = lipgloss.Color("#9d9dff")
		bg = lipgloss.Color("#426292")
		border = lipgloss.Color("#4f4fa1")
		newlySpawnedBorder = lipgloss.Color("#7373d4")
	case 8:
		fg = lipgloss.Color("#7fafe7")
		bg = lipgloss.Color("#438594")
		border = lipgloss.Color("#4496a8")
	case 16:
		fg = lipgloss.Color("#b4f8e7")
		bg = lipgloss.Color("#328570")
		border = lipgloss.Color("#3dad91")
	case 32:
		fg = lipgloss.Color("#aeffc0")
		bg = lipgloss.Color("#236d39")
		border = lipgloss.Color("#237e3d")
	case 64:
		fg = lipgloss.Color("#d8f5a3")
		bg = lipgloss.Color("#7f9726")
		border = lipgloss.Color("#8fac25")
	case 128:
		fg = lipgloss.Color("#ffd5a6")
		bg = lipgloss.Color("#bd9100")
		border = lipgloss.Color("#dda900")
	case 256:
		fg = lipgloss.Color("#ffc582")
		bg = lipgloss.Color("#c25700")
		border = lipgloss.Color("#be5601")
	case 512:
		fg = lipgloss.Color("#ff9182")
		bg = lipgloss.Color("#8f0c0c")
		border = lipgloss.Color("#a10c0c")
	case 1024:
		fg = lipgloss.Color("#ffadc6")
		bg = lipgloss.Color("#770059")
		border = lipgloss.Color("#91006c")
	case 2048:
		fg = lipgloss.Color("#f9ffc2")
		bg = lipgloss.Color("#55007c")
		border = lipgloss.Color("#6a009b")
		// default:
		// 	fg = lipgloss.Color
	}
	return
}
