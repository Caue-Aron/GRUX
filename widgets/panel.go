package widgets

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	StylePanelRegular = iota
	StylePanelBorderless
	StylePanelSunken
)

type Panel struct {
	BaseWidget
	container Container
}

const (
	PanelStyleNormal = 1 << iota
	PanelStyleSunken
	PanelStyleRaised
	PanelStyleDoubleBordered
	PanelStyleThickBordered
	PanelStyleBorderless
)

func NewPanel() *Panel {
	panel := &Panel{
		container: nil,
		BaseWidget: BaseWidget{
			Bitmask: PanelStyleNormal,
			State:   DisplayStateActive,
			Rectangle: rl.Rectangle{
				X: 100, Y: 100,
				Width: 600, Height: 250},
		},
	}

	return panel
}

func (this *Panel) SetContainer(container Container, distribuition ...uint8) bool {
	if this.container != nil {
		return false
	}

	this.container = container
	container.CalculateBounds(&this.Rectangle, distribuition...)

	return true
}

func (this *Panel) GetBounds() rl.Rectangle {
	return this.Rectangle
}

func (this *Panel) Update() {
	this.container.Update()
}

func (this *Panel) Draw() {
	rl.DrawRectangleLinesEx(this.Rectangle, 1, ColorPanelBorder)
	rl.DrawRectangleLines(int32(this.X)+1, int32(this.Y)+1, int32(this.Width)-2, int32(this.Height)-2, ColorPanelInactive)
}

func (this *Panel) CalculateBounds(limits *rl.Rectangle) {

}

func (this *Panel) Activate() {

}

func (this *Panel) Deactivate() {

}
