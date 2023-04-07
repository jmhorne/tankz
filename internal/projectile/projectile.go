package projectile

import "github.com/hajimehoshi/ebiten/v2"

type Projectile interface {
	Update() error
	Draw(screen *ebiten.Image)
}