package main

import (
	"fmt"
	"log"
	"tankz/internal/game"

	"github.com/fstanis/screenresolution"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	WIDTH int
	HEIGHT int
)

func main() {
	if err := SetResolution(); err != nil {
		log.Fatal(err)
	}

	game, err := game.New(WIDTH, HEIGHT)
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowTitle("TANKZ!")

	if err = ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func SetResolution() error {
	res := screenresolution.GetPrimary()

	if res.String() == "" {
		return fmt.Errorf("unable to retrieve screen resolution")
	}

	WIDTH = res.Width
	HEIGHT = res.Height

	return nil
}