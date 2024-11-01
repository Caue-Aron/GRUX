package widgets

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type State uint8
type Bitmask uint

func (f Bitmask) HasFlag(flag Bitmask) bool { return (f & flag) != 0 }
func (f Bitmask) HasAll(flag Bitmask) bool  { return (f & flag) == f }

var defaultFont rl.Font
var ColorBackground rl.Color
var ColorPanelInactive rl.Color
var ColorPanelHighlight rl.Color
var ColorPanelShadow rl.Color
var ColorPanelShadowHighlight rl.Color
var ColorPanelPressed rl.Color
var ColorPanelBorder rl.Color
var ColorPanelInnerBorder rl.Color
var ColorTextInactive rl.Color
var ColorTextActive rl.Color
var ColorPanelSunken rl.Color

const (
	NaturalPadding     = float32(4)
	DoubleBorderOffset = float32(6)
)

const (
	DisplayStateActive State = iota
	DisplayStateInactive
	DisplayStateSelected
	DisplayStatePressed
)

func ScaleRect(inRect *rl.Rectangle, factor float32) *rl.Rectangle {
	rect := *inRect

	rect.X += factor
	rect.Y += factor
	rect.Width -= factor * 2
	rect.Height -= factor * 2

	return &rect
}

func InitWidgetsContext() {
	defaultFont = rl.LoadFontEx("ProggyTiny.ttf", 10, nil)
	ColorBackground = rl.Color{45, 45, 45, 255}
	ColorPanelBorder = rl.Color{71, 71, 71, 255}
	ColorPanelInactive = rl.Color{71, 71, 71, 255}
	ColorPanelHighlight = rl.Color{85, 85, 85, 255}
	ColorPanelShadow = rl.Color{30, 30, 30, 255}
	ColorPanelSunken = rl.Color{40, 40, 40, 255}
	ColorPanelShadowHighlight = rl.Color{59, 59, 59, 255}
	ColorPanelPressed = ColorBackground
	ColorPanelInnerBorder = rl.Color{97, 97, 97, 255}
	ColorTextInactive = rl.Color{124, 124, 124, 255}
	ColorTextActive = rl.Color{226, 226, 226, 255}
}

func DestroyWidgetsContext() {
	rl.UnloadFont(defaultFont)
}

type WidgetDisplayer interface {
	Update()
	Draw()
	CalculateBounds(limits *rl.Rectangle)
	Activate()
	Deactivate()
	GetWorkingArea() *rl.Rectangle
}

type BaseWidget struct {
	rl.Rectangle
	Bitmask
	State
}

func (this *BaseWidget) Activate() {
	this.State = DisplayStateActive

	fmt.Printf("X: %.2f\nY: %.2f\nW: %.2f\nH: %.2f\n", this.X, this.Y, this.Width, this.Height)
}

func (this *BaseWidget) Deactivate() {
	this.State = DisplayStateInactive
}

func (this *BaseWidget) GetWorkingArea() *rl.Rectangle {
	return ScaleRect(&this.Rectangle, NaturalPadding)
}

func (this *BaseWidget) CalculateBounds(limits *rl.Rectangle) {
	this.Rectangle = *limits
}
