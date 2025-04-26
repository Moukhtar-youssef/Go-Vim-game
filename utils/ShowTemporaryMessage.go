package utils

import (
	"github.com/gdamore/tcell/v2"
)

func ShowTemporaryMessage(s tcell.Screen, message string) {
	w, h := s.Size()

	// defining box dimestions
	width := len(message) + 4
	height := 5
	startX := (w - width) / 2
	startY := (h - height) / 2
	// Draw the message box (with transparent background)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			style := tcell.StyleDefault.Foreground(tcell.ColorWhite)
			if y == 0 || y == height-1 || x == 0 || x == width-1 {
				style = style // Draw border
			}

			// Draw the box border
			s.SetContent(startX+x, startY+y, ' ', nil, style)
		}
	}
	// Draw the message inside the box
	PrintCentered(s, startX+2, startY+1, message, tcell.StyleDefault.Foreground(tcell.ColorWhite))
	// Draw button inside the box
	buttonText := "[Press Enter to continue]"
	PrintCentered(s, startX+2, startY+3, buttonText, tcell.StyleDefault.Foreground(tcell.ColorWhite))

	s.Show()
	// Wait for user input (Enter key) to close the overlay and continue
	for {
		event := s.PollEvent()
		switch ev := event.(type) {
		case *tcell.EventKey:
			// Exit overlay when Enter is pressed
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
