package tank

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Tank struct {
	bodyImage   *ebiten.Image
	turretImage *ebiten.Image
	x, y        float64
}

func New(x, y float64) (*Tank, error) {
	var err error
	t := new(Tank)
	t.x = x
	t.y = y

	if t.bodyImage, _, err = ebitenutil.NewImageFromFile("internal/assets/tanks/red/body.png"); err != nil {
		return nil, err
	}

	if t.turretImage, _, err = ebitenutil.NewImageFromFile("internal/assets/tanks/red/turret.png"); err != nil {
		return nil, err
	}

	return t, err
}

func (t *Tank) Update() error {
	return nil
}

func (t *Tank) Draw(screen *ebiten.Image) {
}
