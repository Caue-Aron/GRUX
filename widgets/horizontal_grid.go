package widgets

import rl "github.com/gen2brain/raylib-go/raylib"

type HorizontalGrid struct {
	MultiContainer
}

func NewHorizontalGrid(children []WidgetDisplayer, distribuition ...uint8) *HorizontalGrid {
	container := HorizontalGrid{
		MultiContainer: MultiContainer{
			children: children,
			BaseContainer: BaseContainer{
				BaseWidget: BaseWidget{
					State:   DisplayStateActive,
					Bitmask: 0,
				},
			},
		},
	}

	return &container
}

func (this *HorizontalGrid) CalculateBounds(limits *rl.Rectangle) {
	for _, childWidget := range this.children {
		_ = childWidget
	}
}

func (this *HorizontalGrid) AddChild(child WidgetDisplayer, space uint8) {

	this.CalculateBounds(&this.Rectangle)
}

func (this *HorizontalGrid) RemoveChildren(index ...uint8) {

}

func (this *HorizontalGrid) SetDistribuition(distribuition ...uint8) {

}

func (this *HorizontalGrid) SetChildMargin(margin uint) {

}

func (this *HorizontalGrid) SetPadding(padding uint) {

}

func (this *HorizontalGrid) Update() {
	for _, childWidget := range this.children {
		childWidget.Update()
	}
}

func (this *HorizontalGrid) Draw() {
	for _, childWidget := range this.children {
		childWidget.Draw()
	}
}

func (this *HorizontalGrid) Activate() {
	for _, childWidget := range this.children {
		childWidget.Activate()
	}
}

func (this *HorizontalGrid) Deactivate() {
	for _, childWidget := range this.children {
		childWidget.Deactivate()
	}
}

func (this *HorizontalGrid) GetBounds() rl.Rectangle {
	return this.Rectangle
}
