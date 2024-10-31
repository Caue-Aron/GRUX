package widgets

import (
	"main/widgets/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var defaultFont rl.Font

func InitWidgetsContext() {
	defaultFont = rl.LoadFontEx("ProggyTiny.ttf", 10, nil)
}

func DestroyWidgetsContext() {
	rl.UnloadFont(defaultFont)
}

type WidgetDisplayer interface {
	Update()
	Draw()
	CalculateBounds(limits rl.Rectangle) rl.Rectangle
	Activate()
	Deactivate()
}

type BaseWidget struct {
	bounds *rl.Rectangle
	state  uint8
	style  utils.Bitmask
	parent *BaseWidget
}
