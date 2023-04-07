package tank

import (
	"fmt"
	"image"
	"math"
	"tankz/internal/projectile"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Tank struct {
	bodyImage   *ebiten.Image
	turretImage *ebiten.Image
	turretAngle int
	x, y, speed float64
	activeTank  bool
	power       int
	Fired       bool
}

func New(x, y float64, color string) (*Tank, error) {
	var err error
	t := new(Tank)
	t.x = x
	t.y = y
	t.speed = 0.5
	t.turretAngle = 0
	t.Fired = false
	t.power = 0
	t.activeTank = false

	if t.bodyImage, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("internal/assets/tanks/%s/body.png", color)); err != nil {
		return nil, err
	}

	if t.turretImage, _, err = ebitenutil.NewImageFromFile(fmt.Sprintf("internal/assets/tanks/%s/turret.png", color)); err != nil {
		return nil, err
	}

	return t, err
}

func (t *Tank) Update() error {
	if !t.activeTank {
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		t.turretAngle--

		if t.turretAngle < -90 {
			t.turretAngle = -90
		}

	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		t.turretAngle++

		if t.turretAngle > 90 {
			t.turretAngle = 90
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		t.power++

		if t.power > 100 {
			t.power = 100
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		t.power--

		if t.power < 0 {
			t.power = 0
		}

	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		t.Fired = true
	}

	return nil
}

func (t *Tank) Draw(screen *ebiten.Image) {
	t.drawTurret(screen)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(t.x, t.y)
	screen.DrawImage(t.bodyImage, op)
}

func (t *Tank) Bounds() image.Rectangle {
	return t.bodyImage.Bounds()
}

func (t *Tank) Activate() {
	t.activeTank = true
	t.Fired = false
}

func (t *Tank) Deactivate() {
	t.activeTank = false
}

func (t *Tank) Fire() (projectile.Projectile, error) {
	return projectile.NewCannonball(50, 50)
}

func (t *Tank) drawTurret(screen *ebiten.Image) {
	ti := t.turretImage.Bounds().Size()
	tb := t.bodyImage.Bounds().Size()
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-float64(ti.X)/2, -float64(ti.Y))
	op.GeoM.Rotate(float64(t.turretAngle%360) * 2 * math.Pi / 360)

	op.GeoM.Translate(t.x+float64(tb.X/2), t.y+float64(tb.Y/2))
	screen.DrawImage(t.turretImage, op)

	if t.activeTank {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("angle: %d, power: %d", t.turretAngle, t.power))
	}
}
