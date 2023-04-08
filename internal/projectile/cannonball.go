package projectile

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Cannonball struct {
	props properties
}

func NewCannonball(x, y float64, angle int) (*Cannonball, error) {
	c := new(Cannonball)
	c.props.x = x
	c.props.y = y
	c.props.angle = angle

	var err error
	c.props.image, _, err = ebitenutil.NewImageFromFile("internal/assets/projectiles/cannonball.png")
	return c, err
}

func (c *Cannonball) Update() error {
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
	c.props.x = x
}

func (c *Cannonball) SetY(y float64) {
	c.props.y = y
}
