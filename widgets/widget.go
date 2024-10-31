package widgets

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type State uint8
type Bitmask uint32

func (f Bitmask) HasFlag(flag Bitmask) bool { return (f & flag) != f }

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

const (
	DisplayStateActive State = iota
	DisplayStateInactive
	DisplayStateSelected
	DisplayStatePressed
)

func InitWidgetsContext() {
	defaultFont = rl.LoadFontEx("ProggyTiny.ttf", 10, nil)
	ColorBackground = rl.Color{45, 45, 45, 255}
	ColorPanelInactive = rl.Color{71, 71, 71, 255}
	ColorPanelHighlight = rl.Color{85, 85, 85, 255}
	ColorPanelShadow = ColorBackground
	ColorPanelShadowHighlight = rl.Color{59, 59, 59, 255}
	ColorPanelPressed = ColorBackground
	ColorPanelBorder = rl.Color{30, 30, 30, 255}
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
	GetBounds() rl.Rectangle
}

type BaseWidget struct {
	rl.Rectangle
	Bitmask
	State
}
