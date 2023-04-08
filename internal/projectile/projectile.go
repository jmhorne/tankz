package projectile

import "github.com/hajimehoshi/ebiten/v2"

type Projectile interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type properties struct {
	x, y  float64
	angle int
	image *ebiten.Image
}
