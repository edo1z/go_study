package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func update(screen *ebiten.Image) error {
	_ = ebitenutil.DebugPrint(screen, "Hello World")
	return nil
}

func main() {
	_ = ebiten.Run(update, 320, 240, 2, "Hello")
}
