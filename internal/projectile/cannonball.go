package projectile

import (
	"image"
	"math"
	"tankz/internal/collision"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Cannonball struct {
	props properties
}

func NewCannonball(x, y, velocity, angle float64) (*Cannonball, error) {
	c := new(Cannonball)
	c.props.startX = x
	c.props.startY = y
	c.props.x, c.props.y = x, y
	c.props.angle = angle
	c.props.velocity = velocity
	c.props.time = 0

	c.props.velocityX = c.props.velocity * math.Cos(c.props.angle)
	c.props.velocityY = c.props.velocity * math.Sin(c.props.angle)

	var err error
	c.props.image, _, err = ebitenutil.NewImageFromFile("internal/assets/projectiles/cannonball.png")
	return c, err
}

func (c *Cannonball) Update() error {
	c.props.x = (c.props.velocityX * c.props.time) + c.props.startX
	c.props.y = c.props.startY + (c.props.velocityY * c.props.time) + (gravity * (math.Pow(c.props.time, 2)) / 2)
	c.props.time += .2

	return nil
}

func (c *Cannonball) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(c.props.x, c.props.y)
	screen.DrawImage(c.props.image, op)
}

func (c *Cannonball) Size() image.Point {
	return c.props.image.Bounds().Size()
}

func (c *Cannonball) SetX(x float64) {
	c.props.startX = x
}

func (c *Cannonball) SetY(y float64) {
	c.props.startY = y
}

func (c *Cannonball) X() float64 {
	return c.props.x
}
func (c *Cannonball) Y() float64 {
	return c.props.y
}

func (c *Cannonball) GetCollisionArea() collision.CollisionArea {
	s := c.Size()
	half := float64(s.X)/2
	return collision.Circle(c.props.x+half, c.props.y+half, half)
}
