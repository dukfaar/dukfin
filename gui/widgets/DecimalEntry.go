package widgets

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
	"github.com/shopspring/decimal"
)

type DecimalEntry struct {
	widget.Entry
}

func NewDecimalEntry() *DecimalEntry {
	entry := &DecimalEntry{}
	entry.Validator = func(s string) error {
		s = strings.TrimSpace(s)
		s = strings.ReplaceAll(s, ",", ".")
		_, err := decimal.NewFromString(s)
		return err
	}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *DecimalEntry) GetDecimal() (decimal.Decimal, error) {
	s := strings.TrimSpace(e.Text)
	s = strings.ReplaceAll(s, ",", ".")
	return decimal.NewFromString(s)
}

func (e *DecimalEntry) TypedRune(r rune) {
	if (r >= '0' && r <= '9') || r == '.' || r == ',' {
		e.Entry.TypedRune(r)
	}
}

func (e *DecimalEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := decimal.NewFromString(content); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}

func (e *DecimalEntry) Keyboard() mobile.KeyboardType {
	return mobile.NumberKeyboard
}
