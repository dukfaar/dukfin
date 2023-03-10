package widgets

import (
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
	"github.com/shopspring/decimal"
)

type IntEntry struct {
	widget.Entry
}

func NewIntEntry() *IntEntry {
	entry := &IntEntry{}
	entry.Validator = func(s string) error {
		s = strings.TrimSpace(s)
		_, err := decimal.NewFromString(s)
		return err
	}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *IntEntry) GetInt() (int, error) {
	s := strings.TrimSpace(e.Text)
	i, err := strconv.ParseInt(s, 10, 32)
	return int(i), err
}

func (e *IntEntry) TypedRune(r rune) {
	if r >= '0' && r <= '9' {
		e.Entry.TypedRune(r)
	}
}

func (e *IntEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := strconv.ParseFloat(content, 64); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}

func (e *IntEntry) Keyboard() mobile.KeyboardType {
	return mobile.NumberKeyboard
}
