package main

import (
	"log"

	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{}

type Point struct{ x, y float64 }

type Body struct {
	pos   Point
	color color.Color
	mass  float64
}

var mousePos Point
var mouseBody Body = Body{Point{}, c["white"], 10.0}
var mouseClick, pMouseClick bool

var c map[string]color.Color = map[string]color.Color{
	"red":   color.RGBA{255, 0, 0, 255},
	"black": color.RGBA{0, 0, 0, 255},
	"green": color.RGBA{0, 255, 0, 255},
	"blue":  color.RGBA{0, 0, 255, 255},
	"white": color.RGBA{255, 255, 255, 255},
}

var bodies []Body = []Body{
	{Point{80, 80}, c["white"], 10.0},
	{Point{50, 50}, c["white"], 30.0},
	{Point{100, 20}, c["white"], 5.0},
}

func (b Body) Draw(screen *ebiten.Image) {
	r := float32(b.mass)
	vector.DrawFilledCircle(screen, float32(b.pos.x), float32(b.pos.y), r, b.color, true)
}

func (b Body) DrawOutline(screen *ebiten.Image) {
	r := float32(b.mass)
	vector.StrokeCircle(screen, float32(b.pos.x), float32(b.pos.y), r, 1, b.color, false)
}

func (g *Game) MouseClick() error {
	bodies = append(bodies, mouseBody)
	return nil
}

func (g *Game) Update() error {
	pMouseClick = mouseClick
	mouseClick = inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0)

	if mouseClick && !pMouseClick {
		g.MouseClick()
	}

	x, y := ebiten.CursorPosition()
	mousePos = Point{float64(x), float64(y)}
	mouseBody.pos = mousePos

	// Body resizing
	if ebiten.IsKeyPressed(ebiten.KeyX) {
		mouseBody.mass += 0.2
	}

	// Body resizing
	if ebiten.IsKeyPressed(ebiten.KeyZ) {
		mouseBody.mass -= 0.2
	}

	return nil
}

// ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.2f, TPS: %.2f", ebiten.ActualFPS(), ebiten.ActualTPS()))
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw Bodies
	for _, val := range bodies {
		val.Draw(screen)
	}

	// Draw Cursor Body
	mouseBody.DrawOutline(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetTPS(100)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
