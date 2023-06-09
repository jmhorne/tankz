package projectile

import (
	"tankz/internal/collision"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	gravity = 9.8
)

type Projectile interface {
	Update() error
	Draw(screen *ebiten.Image)
	GetCollisionArea() collision.CollisionArea
	GetDamageDealt() int
	X() float64
	Y() float64
}

type properties struct {
	x, y      float64
	startX    float64
	startY    float64
	angle     float64
	velocity  float64
	velocityX float64
	velocityY float64
	time      float64
	image     *ebiten.Image
}
