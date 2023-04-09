package tank

import (
	"fmt"
	"image"
	"math"
	"tankz/internal/collision"
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

		if t.power > 150 {
			t.power = 150
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
	// create projectile in order to get its dimensions
	radians := float64(t.turretAngle - 90) * (math.Pi/180.0)
	p, err := projectile.NewCannonball(0, 0, float64(t.power), radians)
	ps := p.Size()

	// get tank body and turret image sizes
	tis := t.turretImage.Bounds().Size()
	tbs := t.bodyImage.Bounds().Size()

	// r is distance from middle of tank(origins below) to tip of turret
	r := float64(tis.Y)
	originX := -float64(ps.X/2) + t.x + float64(tbs.X/2)
	originY := -float64(ps.Y/2) + t.y + float64(tbs.Y/2)

	// convert turret angle to radians and calculate x,y of where cannonball
	// should be so it starts at tip of turret
	posX := r * math.Cos(float64(radians))
	posY := r * math.Sin(float64(radians))
	
	// reset projectile position
	p.SetX(posX + originX)
	p.SetY(posY + originY)

	return p, err
}

func (t *Tank) drawTurret(screen *ebiten.Image) {
	ti := t.turretImage.Bounds().Size()
	tb := t.bodyImage.Bounds().Size()
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-float64(ti.X)/2, -float64(ti.Y))
	op.GeoM.Rotate(float64(t.turretAngle%360) * 2 * math.Pi / 360)

	op.GeoM.Translate(t.x+float64(tb.X/2), t.y+float64(tb.Y/2))
	screen.DrawImage(t.turretImage, op)

	//TODO add HUD here
	if t.activeTank {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("angle: %d, power: %d", t.turretAngle, t.power))
	}
}

func (t *Tank) GetCollisionArea() collision.CollisionArea {
	tb := t.bodyImage.Bounds().Size()
	return collision.Rectangle(t.x+50, t.y+float64(tb.X/2), float64(tb.X)-100, float64(tb.Y)-float64(tb.X/2))
}