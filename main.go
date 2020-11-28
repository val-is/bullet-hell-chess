package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/val-is/bullet-hell-chess/engine"
)

func main() {
	g, err := engine.NewGameInstance()
	if err != nil {
		log.Fatalf("Error when initializing game: %s", err)
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatalf("Game crashed with %s", err)
	}
}
