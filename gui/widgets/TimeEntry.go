package widgets

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TimeEntry struct {
	widget.Entry
}

func NewTimeEntry() *TimeEntry {
	entry := &TimeEntry{}
	entry.Text = time.Now().Format(time.RFC3339)
	entry.Validator = func(s string) error {
		_, err := time.Parse(time.RFC3339, s)
		return err
	}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *TimeEntry) GetTime() (time.Time, error) {
	return time.Parse(time.RFC3339, e.Text)
}

func (e *TimeEntry) SetTime(t time.Time) {
	e.SetText(t.Format(time.RFC3339))
}

func (e *TimeEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := time.Parse(time.RFC3339, content); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}
