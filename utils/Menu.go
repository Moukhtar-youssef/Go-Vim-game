package utils

import "github.com/gdamore/tcell/v2"

type MenuItems struct {
	ID string
	Name     string
	function func()
}

type Menu struct {
	Title     string
	MenuItems []MenuItems
	s         tcell.Screen
	selected int
}

func New(title string, Items []MenuItems,s tcell.Screen) *Menu {
	return &Menu{
		Title: title,
		MenuItems: Items,
		selected: 0,
		s: s,
	}
}



func DrawMenu(s tcell.Screen, selected int, TextArray []string) {
	w, h := s.Size()
	menuHeight := len(TextArray) // Number of menu options
	// Calculate the starting Y position for vertical centering
	startY := (h - menuHeight) / 2
	for i, option := range TextArray {
		style := tcell.StyleDefault
		if i == selected {
			style = style.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite)
		}

		displayText := option
		if option != "Exit" {
			displayText = "[ ] " + option
		}
		PrintCentered(s, w/2, startY+i, displayText, style)
	}
}
