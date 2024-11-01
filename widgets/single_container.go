package widgets

import rl "github.com/gen2brain/raylib-go/raylib"

type SingleContainer struct {
	BaseContainer
	child WidgetDisplayer
}

func NewSingleContainer(child WidgetDisplayer, padding uint) *SingleContainer {
	container := &SingleContainer{
		child: child,
		BaseContainer: BaseContainer{
			padding: padding,
			BaseWidget: BaseWidget{
				State:   DisplayStateActive,
				Bitmask: 0,
			},
		},
	}

	return container
}

func (this *SingleContainer) SetPadding(padding uint) {
	this.BaseContainer.SetPadding(padding)
	this.child.CalculateBounds(this.GetWorkingArea())
}

func (this *SingleContainer) GetWorkingArea() *rl.Rectangle {
	return ScaleRect(&this.Rectangle, float32(this.padding))
}

func (this *SingleContainer) CalculateBounds(limits *rl.Rectangle) {
	this.BaseWidget.CalculateBounds(limits)
	this.child.CalculateBounds(this.GetWorkingArea())
}

func (this *SingleContainer) Update() {
	this.child.Update()
}

func (this *SingleContainer) Draw() {
	this.child.Draw()
}

func (this *SingleContainer) Activate() {
	this.child.Activate()
}

func (this *SingleContainer) Deactivate() {
	this.child.Deactivate()
}
