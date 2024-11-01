package widgets

type Container interface {
	WidgetDisplayer
	SetPadding(padding uint)
}

type BaseContainer struct {
	BaseWidget
	padding uint
}

func (this *BaseContainer) SetPadding(padding uint) {
	this.padding = padding
}
