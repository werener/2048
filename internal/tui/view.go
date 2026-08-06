package tui

import (
	"fmt"
	"image/color"

	_ "github.com/floatpane/bubble-overlay"
	overlay "github.com/floatpane/bubble-overlay"

	tea "charm.land/bubbletea/v2"
	styling "charm.land/lipgloss/v2"
	core "github.com/werener/2048.git/internal/core/models"
)

func (m model) View() tea.View {
	margin := 2

	window := styling.NewStyle().
		Width(m.width).
		Height(m.height).
		Margin(margin).
		Align(styling.Center)

	availableSquare := min(m.width/2, m.height)

	// account for help menu, score and margin
	gridSize := availableSquare - 2 - 2 - (2 * margin)

	grid := m.gridView(gridSize)

	help := styling.NewStyle().
		Height(2).
		Render(m.help.View(m.keys))

	score := styling.NewStyle().
		Height(2).
		Render(fmt.Sprintf("Your score: %d", m.game.Score()))

	view := window.Render(grid + "\n" + score + "\n" + help)

	gameState := m.game.State()
	if m.showEndlessPopup {
		view = overlay.Center(view, m.endlessPopup(), m.width, m.height)
	} else if gameState == core.DEFEAT {
		view = overlay.Center(view, m.defeatScreen(), m.width, m.height)
	} else if gameState == core.WIN {
		view = overlay.Center(view, m.winScreen(), m.width, m.height)
	}

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
		rows[i] = styling.JoinHorizontal(styling.Top, tiles...)
	}

	grid := styling.JoinVertical(styling.Left, rows...)

	return styling.NewStyle().
		Height(size).
		Width(size * 2).
		AlignVertical(styling.Center).AlignHorizontal(styling.Center).
		BorderStyle(styling.NormalBorder()).
		Render(grid)
}

func (m model) tileView(tile core.Tile, size int) string {
	repr := tile.Repr()

	bg, fg, border, newlySpawnedBorder := colorTile(tile)
	style := styling.NewStyle().
		// alignment
		Height(size).Width(size*2 - 1).
		AlignHorizontal(styling.Center).AlignVertical(styling.Center).
		// colors
		Background(bg).Foreground(fg).
		BorderStyle(styling.NormalBorder()).
		BorderForeground(border).BorderBackground(bg)
	// show if a tile has just been spawned
	if tile.NewlySpawned {
		style = style.BorderStyle(styling.ASCIIBorder()).BorderForeground(newlySpawnedBorder)
	}
	// if a tile is empty, don't show a border
	if tile.IsEmpty() {
		style = style.UnsetBorderStyle()
	}

	return style.Render(repr)
}

func colorTile(tile core.Tile) (
	bg color.Color, fg color.Color,
	border color.Color, newlySpawnedBorder color.Color,
) {
	switch tile.Value {
	case 1 << 0:
		return
	case 1 << 1:
		bg = styling.Color("#2c2c79")
	case 1 << 2:
		bg = styling.Color("#286994")
	case 1 << 3:
		bg = styling.Color("#438594")
	case 1 << 4:
		bg = styling.Color("#328570")
	case 1 << 5:
		bg = styling.Color("#236d39")
	case 1 << 6:
		bg = styling.Color("#7f9726")
	case 1 << 7:
		bg = styling.Color("#bd9100")
	case 1 << 8:
		bg = styling.Color("#c25700")
	case 1 << 9:
		bg = styling.Color("#8f0c0c")
	case 1 << 10:
		bg = styling.Color("#770059")
	case 1 << 11:
		bg = styling.Color("#55007c")
	case 1 << 12:
		bg = styling.Color("#393a3a")
		fg = styling.Lighten(bg, 0.5)
		border = styling.Color("#ff3a5b")
		return
	case 1 << 13:
		bg = styling.Color("#393a3a")
		fg = styling.Lighten(bg, 0.5)
		border = styling.Color("#6b0380")
		return
	case 1 << 14:
		bg = styling.Color("#393a3a")
		fg = styling.Lighten(bg, 0.5)
		border = styling.Color("#00179b")
		return
	case 1 << 15:
		bg = styling.Color("#393a3a")
		fg = styling.Lighten(bg, 0.5)
		border = styling.Color("#00c41a")
		return
	case 1 << 16:
		bg = styling.Color("#393a3a")
		fg = styling.Lighten(bg, 0.5)
		border = styling.Color("#fffb00")
		return
	case 1 << 17:
		bg = styling.Color("#393a3a")
		fg = styling.Lighten(bg, 0.5)
		border = styling.Color("#ff9100")
		return
	}

	fg = styling.Lighten(bg, 0.5)
	border = styling.Darken(bg, 0.15)
	newlySpawnedBorder = styling.Lighten(border, 0.3)
	return bg, fg, border, newlySpawnedBorder
}
