package widgets

type MultiContainer struct {
	BaseContainer
	children []WidgetDisplayer
}

func (this *MultiContainer) Update() {
	for _, childWidget := range this.children {
		childWidget.Update()
	}
}

func (this *MultiContainer) Draw() {
	for _, childWidget := range this.children {
		childWidget.Draw()
	}
}

func (this *MultiContainer) Activate() {
	for _, childWidget := range this.children {
		childWidget.Activate()
	}
}

func (this *MultiContainer) Deactivate() {
	for _, childWidget := range this.children {
		childWidget.Deactivate()
	}
}
