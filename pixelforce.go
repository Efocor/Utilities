//...............................................................................................................
/*
Simple pixel art editor using Ebiten, only for educational purposes.
I used this myself to make some sprites for my games, but i jumped to a better editor made in C++.

This editor allows you to draw, delete, save and clear the canvas by pressing one of the sizes buttons.
You can also change the color by clicking the color button and selecting one of the colors in the color picker.

You use the left mouse button to draw or delete pixels on the canvas.

Made by @FECORO
*/
//...............................................................................................................

// ...... Stack
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// ...... CODE:
// constants for canvas sizes
const (
	canvasSize8x8     = 8
	canvasSize16x16   = 16
	canvasSize32x32   = 32
	canvasSize64x64   = 64
	canvasSize128x128 = 128
)

// pixelArtEditor represents the main application
type pixelArtEditor struct {
	canvas          *ebiten.Image // the canvas where the pixel art is drawn
	currentColor    color.Color   // the current selected color
	buttons         []*button     // list of buttons in the ui
	canvasWidth     int           // current canvas width
	canvasHeight    int           // current canvas height
	selectedTool    string        // currently selected tool (paint, delete)
	isPainting      bool          // flag to check if the user is painting
	colorPickerOpen bool          // flag to check if the color picker is open
	colorPicker     *colorPicker  // color picker instance
}

// button represents a ui button
type button struct {
	x, y, width, height int
	label               string
	action              func()
	color               color.Color
}

// newButton creates a new button instance
func newButton(x, y, width, height int, label string, action func(), color color.Color) *button {
	return &button{
		x:      x,
		y:      y,
		width:  width,
		height: height,
		label:  label,
		action: action,
		color:  color,
	}
}

// update handles the button click event
func (b *button) update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if x >= b.x && x <= b.x+b.width && y >= b.y && y <= b.y+b.height {
			b.action()
		}
	}
}

// draw renders the button on the screen
func (b *button) draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, float64(b.x), float64(b.y), float64(b.width), float64(b.height), b.color)
	ebitenutil.DebugPrintAt(screen, b.label, b.x+10, b.y+10)
}

// colorPicker represents a color picker ui
type colorPicker struct {
	x, y, width, height int
	colors              []color.Color
	selectedColor       color.Color
}

// newColorPicker creates a new color picker instance
func newColorPicker(x, y int) *colorPicker {
	return &colorPicker{
		x:      x,
		y:      y,
		width:  200,
		height: 100,
		colors: []color.Color{
			color.White,
			color.Black,
			color.RGBA{255, 0, 0, 255},   // red
			color.RGBA{0, 255, 0, 255},   // green
			color.RGBA{0, 0, 255, 255},   // blue
			color.RGBA{255, 255, 0, 255}, // yellow
			color.RGBA{255, 165, 0, 255}, // orange
			color.RGBA{128, 0, 128, 255}, // purple
		},
		selectedColor: color.White,
	}
}

// update handles the color picker interaction
func (cp *colorPicker) update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if x >= cp.x && x <= cp.x+cp.width && y >= cp.y && y <= cp.y+cp.height {
			index := (x - cp.x) / 25 // each color box is 25x25 pixels
			if index >= 0 && index < len(cp.colors) {
				cp.selectedColor = cp.colors[index]
			}
		}
	}
}

// draw renders the color picker on the screen
func (cp *colorPicker) draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, float64(cp.x), float64(cp.y), float64(cp.width), float64(cp.height), color.RGBA{200, 200, 200, 255})
	for i, col := range cp.colors {
		ebitenutil.DrawRect(screen, float64(cp.x+i*25), float64(cp.y), 25, 25, col)
	}
}

// newPixelArtEditor initializes a new pixel art editor
func newPixelArtEditor() *pixelArtEditor {
	p := &pixelArtEditor{
		canvasWidth:  canvasSize64x64,
		canvasHeight: canvasSize64x64,
		currentColor: color.Black, // default color is black
		selectedTool: "paint",
	}
	p.canvas = ebiten.NewImage(p.canvasWidth, p.canvasHeight)
	p.canvas.Fill(color.White) // initialize canvas with white background

	// initialize buttons
	p.buttons = []*button{
		newButton(10, 10, 80, 30, "paint", func() { p.selectedTool = "paint" }, color.RGBA{0, 255, 0, 255}),
		newButton(100, 10, 80, 30, "delete", func() { p.selectedTool = "delete" }, color.RGBA{255, 0, 0, 255}),
		newButton(190, 10, 80, 30, "save", p.saveImage, color.RGBA{0, 0, 255, 255}),
		newButton(280, 10, 80, 30, "clear", p.clearCanvas, color.RGBA{255, 255, 0, 255}),
		newButton(370, 10, 80, 30, "8x8", func() { p.changeCanvasSize(canvasSize8x8, canvasSize8x8) }, color.RGBA{255, 165, 0, 255}),
		newButton(460, 10, 80, 30, "16x16", func() { p.changeCanvasSize(canvasSize16x16, canvasSize16x16) }, color.RGBA{255, 192, 203, 255}),
		newButton(550, 10, 80, 30, "32x32", func() { p.changeCanvasSize(canvasSize32x32, canvasSize32x32) }, color.RGBA{128, 0, 128, 255}),
		newButton(640, 10, 80, 30, "64x64", func() { p.changeCanvasSize(canvasSize64x64, canvasSize64x64) }, color.RGBA{0, 255, 255, 255}),
		newButton(730, 10, 80, 30, "128x128", func() { p.changeCanvasSize(canvasSize128x128, canvasSize128x128) }, color.RGBA{128, 128, 128, 255}),
		newButton(820, 10, 80, 30, "color", func() { p.colorPickerOpen = !p.colorPickerOpen }, color.RGBA{200, 200, 200, 255}),
	}

	// initialize color picker
	p.colorPicker = newColorPicker(10, 50)

	return p
}

// changeCanvasSize changes the canvas size and resets the canvas
func (p *pixelArtEditor) changeCanvasSize(width, height int) {
	p.canvasWidth = width
	p.canvasHeight = height
	p.canvas = ebiten.NewImage(p.canvasWidth, p.canvasHeight)
	p.canvas.Fill(color.White) // reset canvas to white
}

// clearCanvas clears the canvas to white
func (p *pixelArtEditor) clearCanvas() {
	p.canvas.Fill(color.White)
}

// saveImage saves the current canvas as a png file
func (p *pixelArtEditor) saveImage() {
	img := image.NewRGBA(image.Rect(0, 0, p.canvasWidth, p.canvasHeight))
	for y := 0; y < p.canvasHeight; y++ {
		for x := 0; x < p.canvasWidth; x++ {
			img.Set(x, y, p.canvas.At(x, y))
		}
	}

	file, err := os.Create("pixelart.png")
	if err != nil {
		log.Println("error creating file:", err)
		return
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		log.Println("error encoding png:", err)
		return
	}
	log.Println("image saved successfully")
}

// Update handles the game logic
func (p *pixelArtEditor) Update() error {
	// handle button clicks
	for _, button := range p.buttons {
		button.update()
	}

	// handle color picker
	if p.colorPickerOpen {
		p.colorPicker.update()
		p.currentColor = p.colorPicker.selectedColor
	}

	// handle painting or deleting
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		p.isPainting = true
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		p.isPainting = false
	}

	if p.isPainting {
		x, y := ebiten.CursorPosition()
		// adjust for canvas position (centered)
		canvasX := (x - (1024/2 - p.canvasWidth*2)) / 4
		canvasY := (y - 100) / 4
		if canvasX >= 0 && canvasX < p.canvasWidth && canvasY >= 0 && canvasY < p.canvasHeight {
			if p.selectedTool == "paint" {
				p.canvas.Set(canvasX, canvasY, p.currentColor)
			} else if p.selectedTool == "delete" {
				p.canvas.Set(canvasX, canvasY, color.White) // delete by setting pixel to white
			}
		}
	}

	return nil
}

// draw renders the game
func (p *pixelArtEditor) Draw(screen *ebiten.Image) {
	// fill the entire screen with grey
	grey := color.RGBA{50, 50, 50, 255} // light grey color
	ebitenutil.DrawRect(screen, 0, 0, 1024, 768, grey)

	// draw canvas (centered)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(4, 4)                                     // scale canvas for better visibility
	op.GeoM.Translate(float64(1024/2-p.canvasWidth*2), 100) // center the canvas
	screen.DrawImage(p.canvas, op)

	// draw buttons
	for _, button := range p.buttons {
		button.draw(screen)
	}

	// draw color picker if open
	if p.colorPickerOpen {
		p.colorPicker.draw(screen)
	}
}

// Layout defines the game layout
func (p *pixelArtEditor) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1024, 768
}

func main() {
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Pixel FORCE | Pixel Art Editor @FECORO")

	p := newPixelArtEditor()
	if err := ebiten.RunGame(p); err != nil {
		log.Fatal(err)
	}
}

// ...... END
