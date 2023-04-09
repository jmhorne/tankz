package game

import (
	"fmt"
	"image/color"
	"math"
	"tankz/internal/collision"
	"tankz/internal/projectile"
	"tankz/internal/tank"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	width, height, groundLevel int
)

type Game struct {
	Running bool

	players          []*tank.Tank
	activePlayer     int
	activeProjectile projectile.Projectile
}

func New(w, h int) (*Game, error) {
	width = w
	height = h
	groundLevel = height - (height / 5)

	var err error
	g := new(Game)
	g.Running = true
	g.activeProjectile = nil
	g.activePlayer = 0

	g.players = make([]*tank.Tank, 2)

	if g.players[0], err = tank.New(float64(width-150), float64(groundLevel-100), "red"); err != nil {
		return nil, err
	}

	if g.players[1], err = tank.New(10, float64(groundLevel-100), "blue"); err != nil {
		return nil, err
	}

	g.players[g.activePlayer].Activate()

	return g, err
}

func (g *Game) Update() error {
	var err error
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.Running = false
	}

	if !g.Running {
		// return fmt.Errorf("done")
		return nil
	}

	if g.activeProjectile != nil {
		g.activeProjectile.Update()

		g.testProjectile()

		return nil
	}

	for _, p := range g.players {
		p.Update()
	}

	ap := g.players[g.activePlayer]

	if ap.Fired {
		if g.activeProjectile, err = ap.Fire(); err != nil {
			return err
		}
		ap.Deactivate()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// sky
	screen.Fill(color.RGBA{3, 252, 207, 0xFF})

	// ground
	vector.DrawFilledRect(screen, 0, float32(groundLevel), float32(width), float32(height-groundLevel), color.RGBA{3, 252, 86, 0xFF}, true)

	// players
	for _, p := range g.players {
		p.Draw(screen)
	}

	// projectile
	if g.activeProjectile != nil {
		g.activeProjectile.Draw(screen)
	}

	if !g.Running {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Player %d WON!!", g.activePlayer+1))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func (g *Game) testProjectile() {
	// check if out of bounds
	x := g.activeProjectile.X()
	y := g.activeProjectile.Y()

	if x < 0 || x > float64(width) || y > float64(groundLevel) {
		g.activeProjectile = nil
		g.switchPlayer()
		return
	}

	opponent := g.players[int(math.Abs(float64(g.activePlayer-1)))]

	if collision.Collides(opponent.GetCollisionArea(), g.activeProjectile.GetCollisionArea()) {
		opponent.TakeDamage(g.activeProjectile.GetDamageDealt())

		if opponent.Health <= 0 {
			g.Running = false
		} else {
			g.switchPlayer()
			g.activeProjectile = nil
		}
	}
}

func (g *Game) switchPlayer() {
	g.activePlayer = int(math.Abs(float64(g.activePlayer - 1)))
	g.players[g.activePlayer].Activate()
}
