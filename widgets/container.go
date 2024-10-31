package widgets

import rl "github.com/gen2brain/raylib-go/raylib"

type Container interface {
	CalculateSpace(limits rl.Rectangle) rl.Rectangle
	AddChild(child *WidgetDisplayer)
	RemoveChildren(index uint8)
	SetSpaceDistribuition(distribuition ...uint8)
	SetChildrenMargin(margin uint)
	SetPadding(padding uint)
}

type BaseContainer struct {
	padding     uint
	childMargin uint
}
