package widgets

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	PanelStyleNormal = Bitmask(1) << iota
	PanelStyleSunken
	PanelStyleRaised
	PanelStyleDoubleBordered
	// PanelStyleThickBordered
	PanelStyleBorderless
	PanelStyleShadowless
)

type Panel struct {
	BaseWidget
	container Container
}

func NewPanel(style Bitmask) *Panel {
	panel := &Panel{
		container: nil,
		BaseWidget: BaseWidget{
			Bitmask: style,
			State:   DisplayStateActive,
		},
	}

	return panel
}

func (this *Panel) SetContainer(container Container) bool {
	if this.container != nil {
		return false
	}

	this.container = container
	this.container.CalculateBounds(this.GetWorkingArea())

	return true
}

func (this *Panel) Update() {
	if this.container != nil {
		this.container.Update()
	}
}

func (this *Panel) Draw() {
	borderOffset := float32(1)
	shadowOffset := float32(0)

	if this.HasFlag(PanelStyleSunken) {
		rl.DrawRectangleRec(this.Rectangle, ColorPanelSunken)
		borderOffset = 0
		shadowOffset = 1
	} else {
		rl.DrawRectangleRec(this.Rectangle, ColorBackground)
	}

	if !this.HasFlag(PanelStyleBorderless) {
		rl.DrawRectangleLinesEx(*ScaleRect(&this.Rectangle, borderOffset), 1, ColorPanelBorder)

		if this.HasFlag(PanelStyleDoubleBordered) {
			rl.DrawRectangleLinesEx(*ScaleRect(&this.Rectangle, DoubleBorderOffset), 1, ColorPanelBorder)
		}
	}
	if !this.HasFlag(PanelStyleShadowless) {
		rl.DrawRectangleLinesEx(*ScaleRect(&this.Rectangle, shadowOffset), 1, ColorPanelShadow)
	}

	// Non panel specific
	if this.container != nil {
		this.container.Draw()
	}
}

func (this *Panel) Activate() {
	this.BaseWidget.Activate()
	if this.container != nil {
		this.container.Activate()
	}
}

func (this *Panel) Deactivate() {
	this.BaseWidget.Deactivate()
	if this.container != nil {
		this.container.Deactivate()
	}
}

func (this *Panel) CalculateBounds(limits *rl.Rectangle) {
	this.BaseWidget.CalculateBounds(limits)
	if this.container != nil {
		this.container.CalculateBounds(limits)
	}
}

func (this *Panel) GetWorkingArea() *rl.Rectangle {
	if this.HasFlag(PanelStyleDoubleBordered) {
		return ScaleRect(&this.Rectangle, NaturalPadding*2)
	}

	return ScaleRect(&this.Rectangle, NaturalPadding)
}
