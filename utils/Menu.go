package utils

import "github.com/gdamore/tcell/v2"

type MenuItems struct {
	ID       string
	Name     string
	CheckBox bool
	Function func()
}

type Menu struct {
	Title     string
	MenuItems []MenuItems
	s         tcell.Screen
	selected  int
}

func NewMenu(title string, Items []MenuItems, s tcell.Screen) *Menu {
	return &Menu{
		Title:     title,
		MenuItems: Items,
		s:         s,
		selected:  0,
	}
}

func (m *Menu) HandleInput() {
	selected := m.selected
	for {
		m.s.Clear()
		m.DrawMenu(selected)
		m.s.Show()
		event := m.s.PollEvent()
		switch ev := event.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyUp:
				if selected > 0 {
					selected--
				}
			case tcell.KeyDown:
				if selected < len(m.MenuItems)-1 {
					selected++
				}
			case tcell.KeyEnter:
				m.HandleSelection(selected)
			case tcell.KeyEsc:
				return
			case tcell.KeyRune:
				switch ev.Rune() {
				case 'k':
					if selected > 0 {
						selected--
					}
				case 'j':
					if selected < len(m.MenuItems)-1 {
						selected++
					}
				case 'q':
					return
				default:
					m.s.Clear()
					ShowTemporaryMessage(m.s, "plz")
				}
			default:
				m.s.Clear()
				ShowTemporaryMessage(m.s, "plz")

			}
		}
	}
}

func (m *Menu) HandleSelection(index int) {
	m.MenuItems[index].Function()
}

func (m *Menu) DrawMenu(selected int) {
	w, h := m.s.Size()
	menuHeight := len(m.MenuItems)
	startY := (h - menuHeight) / 2
	for i, option := range m.MenuItems {
		style := tcell.StyleDefault
		if i == selected {
			style = style.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite)
		}
		displayText := option.Name
		if option.CheckBox {
			displayText = "[ ] " + option.Name
		}
		PrintCentered(m.s, w/2, startY+i, displayText, style)
	}
}
