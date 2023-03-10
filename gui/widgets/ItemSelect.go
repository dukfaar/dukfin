package widgets

import "fyne.io/fyne/v2/widget"

type ItemSelect[t interface{}] struct {
	widget.Select

	strF         func(t) string
	selectedItem t
}

func NewItemSelect[t interface{}](list []t, strF func(selectedItem t) string) *ItemSelect[t] {
	result := &ItemSelect[t]{
		Select: widget.Select{
			Options:     make([]string, 0, len(list)),
			PlaceHolder: "Select One",
		},
		strF: strF,
	}
	itemMap := make(map[string]t)
	for _, i := range list {
		displayName := strF(i)
		itemMap[displayName] = i
		result.Options = append(result.Options, displayName)
	}
	result.OnChanged = func(s string) {
		i, ok := itemMap[s]
		if ok {
			result.selectedItem = i
		}
	}
	result.ExtendBaseWidget(result)
	return result
}

func (s *ItemSelect[t]) Get() t {
	return s.selectedItem
}

func (s *ItemSelect[t]) Set(i t) {
	s.SetSelected(s.strF(i))
}
