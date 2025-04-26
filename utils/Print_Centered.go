package utils

import "github.com/gdamore/tcell/v2"

func PrintCentered(s tcell.Screen, x, y int, text string, style tcell.Style) {
	x -= len(text) / 2
	for i, r := range text {
		s.SetContent(x+i, y, r, nil, style)
	}
}
