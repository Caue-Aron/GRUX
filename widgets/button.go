/***************************
WARNING: this file is a test. You can create a new Button in the main file,
call the Update() and Draw() functions at the right time and it should work just fine.
***************************/

package widgets

import (
	"main/widgets/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	DisplayStateActive = iota
	DisplayStateSelected
	DisplayStatePressed
	DisplayStateInactive
)

type Button struct {
	rl.Rectangle
	*rl.Font
	text    string
	textPos rl.Vector2
	state   uint8
}

func (btn *Button) Deactivate() {
	btn.state = DisplayStateActive
}

func (btn *Button) Draw() {
	switch btn.state {

	case DisplayStateActive:
		rl.DrawRectangleGradientV(int32(btn.X), int32(btn.Y), int32(btn.Width), int32(btn.Height), color.PanelHighlight, color.PanelShadow)
	case DisplayStatePressed:
		rl.DrawRectangleRec(btn.Rectangle, color.PanelPressed)
	case DisplayStateSelected:
		rl.DrawRectangleGradientV(int32(btn.X), int32(btn.Y), int32(btn.Width), int32(btn.Height), color.PanelHighlight, color.PanelShadowHighlight)
	case DisplayStateInactive:
		rl.DrawRectangleRec(btn.Rectangle, color.PanelInactive)
		rl.DrawTextEx(*btn.Font, btn.text, btn.textPos, float32(btn.Font.BaseSize), float32(0.45), color.TextInactive)
		rl.DrawRectangleLines(int32(btn.X-1), int32(btn.Y-1), int32(btn.Width+2), int32(btn.Height+2), color.PanelBorder)
		return
	}

	rl.DrawTextEx(*btn.Font, btn.text, btn.textPos, float32(btn.Font.BaseSize), float32(0.45), color.TextActive)
	rl.DrawRectangleLinesEx(btn.Rectangle, 1, color.PanelInnerBorder)
	rl.DrawRectangleLines(int32(btn.X-1), int32(btn.Y-1), int32(btn.Width)+3, int32(btn.Height+2), color.PanelBorder)
}

func (btn *Button) Update() {
	if btn.state != DisplayStateInactive {
		mousePos := rl.GetMousePosition()
		mouseInside := mousePos.X >= btn.X && mousePos.X <= btn.Width+btn.X && mousePos.Y >= btn.Y && mousePos.Y <= btn.Height+btn.Y

		if mouseInside {
			if (rl.IsMouseButtonPressed(rl.MouseButtonLeft) || btn.state == DisplayStatePressed) && rl.IsMouseButtonDown(rl.MouseButtonLeft) {
				btn.state = DisplayStatePressed
			} else {
				btn.state = DisplayStateSelected
			}
		} else {
			btn.state = DisplayStateActive
		}
	}
}

func NewButton(position rl.Vector2, text string) *Button {
	fontSpacing := float32(0.45)
	fontDimensions := rl.MeasureTextEx(defaultFont, text, float32(defaultFont.BaseSize), fontSpacing)

	padding := float32(10)
	buttonBounds := rl.Rectangle{
		X: float32(int(position.X)), Y: float32(int(position.Y)),
		Width:  float32(int(fontDimensions.X + padding*2)),
		Height: float32(int(fontDimensions.Y + padding*2))}

	textPos := rl.Vector2{
		X: float32(int(buttonBounds.X + padding)),
		Y: float32(int(buttonBounds.Y + padding)),
	}

	return &Button{
		Rectangle: buttonBounds,
		text:      text,
		textPos:   textPos,
		Font:      &defaultFont,
		state:     DisplayStateActive,
	}
}
