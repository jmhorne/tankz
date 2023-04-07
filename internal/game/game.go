package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	width, height, groundLevel int
)

type Game struct {
	Running bool
}

func New(w, h int) (*Game, error) {
	width = w
	height = h
	groundLevel = height - (height / 5)

	var err error
	g := new(Game)

	g.Running = true

	return g, err
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.Running = false
	}

	if !g.Running {
		return fmt.Errorf("done")
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// sky
	screen.Fill(color.RGBA{3, 252, 207, 0xFF})

	// ground
	vector.DrawFilledRect(screen, 0, float32(groundLevel), float32(width), float32(height - groundLevel), color.RGBA{3, 252, 86, 0xFF}, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}
