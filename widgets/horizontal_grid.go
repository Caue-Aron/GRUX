package widgets

import rl "github.com/gen2brain/raylib-go/raylib"

type HorizontalGrid struct {
	MultiContainer
}

func NewHorizontalGrid(children []WidgetDisplayer, limits *rl.Rectangle, distribuition ...uint8) *HorizontalGrid {
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

	container.CalculateBounds(limits)

	return &container
}

func (this *HorizontalGrid) CalculateBounds(limits *rl.Rectangle) {
	// set space for every children

	for _, childWidget := range this.children {
		_ = childWidget
	}
}

func (this *HorizontalGrid) AddChild(child WidgetDisplayer, space uint8) {
	this.CalculateBounds(&this.Rectangle)
}

func (this *HorizontalGrid) RemoveChild(index uint8) {

}

func (this *HorizontalGrid) SetDistribuition(distribuition ...uint8) {

}

func (this *HorizontalGrid) SetChildMargin(margin uint) {

}

func (this *HorizontalGrid) SetPadding(padding uint) {

}
