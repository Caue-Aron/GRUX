package widgets

type Container interface {
	WidgetDisplayer
	AddChild(child WidgetDisplayer, space uint8)
	RemoveChildren(index ...uint8)
	SetDistribuition(distribuition ...uint8)
	SetChildMargin(margin uint)
	SetPadding(padding uint)
}

type BaseContainer struct {
	BaseWidget
	padding     uint
	childMargin uint
}

type SingleContainer struct {
	BaseContainer
	child WidgetDisplayer
}

type MultiContainer struct {
	BaseContainer
	children []WidgetDisplayer
}
