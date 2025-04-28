package utils

import (
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"
)

func ShowTemporaryMessage(s tcell.Screen, message string) {
	w, h := s.Size()
	width := utf8.RuneCountInString(message) + 4
	height := 5
	startX := (w - width) / 2
	startY := (h - height) / 2
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite)
	s.SetContent(startX, startY, ' ', nil, style)
	PrintCentered(s, startX+2, startY+1, message, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	buttonText := "[Press Enter to continue]"
	PrintCentered(s, startX+2, startY+3, buttonText, tcell.StyleDefault.Foreground(tcell.ColorWhite))

	s.Show()

	for {
		event := s.PollEvent()
		switch ev := event.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEnter {
				ClearOverlay(s)
				return
			}
		}
	}
}

func ClearOverlay(s tcell.Screen) {
	s.Clear()
	s.Show()
}
