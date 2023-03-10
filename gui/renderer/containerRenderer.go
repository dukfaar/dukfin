package renderer

import "fyne.io/fyne/v2"

type ContainerRenderer struct {
	*fyne.Container
}

func (r *ContainerRenderer) Destroy() {
}

func (r *ContainerRenderer) Objects() []fyne.CanvasObject {
	return r.Container.Objects
}

func (r *ContainerRenderer) Layout(s fyne.Size) {
	if r.Container.Layout != nil {
		r.Container.Layout.Layout(r.Objects(), s)
	}
}
