package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	width, height int
)

type Game struct {
}

func New(w, h int) (*Game, error) {
	width = w
	height = h

	var err error
	g := new(Game)

	return g, err
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xFF, 0, 0, 0xFF})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}
