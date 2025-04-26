package HJKL_Maze

import (
	"fmt"
	"os"

	"github.com/Moukhtar-youssef/Go-VIM-games.git/utils"
	"github.com/gdamore/tcell/v2"
)

var HJKLLevels = []string{
	"Easy", "Medium", "Hard", "Extreme",
}

func HJKL_MAZE_GAME(s tcell.Screen) {
	selected := 0
	s.Clear()
	for {
		utils.DrawMenu(s, selected, HJKLLevels)
		s.Show()
		switch ev := s.PollEvent().(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyUp:
				if selected > 0 {
					selected--
				}
			case tcell.KeyDown:
				if selected < len(HJKLLevels)-1 {
					selected++
				}
			case tcell.KeyEnter:
				handleSelection(selected, s)
				s.Show()
			case tcell.KeyEscape:
				return
			default:
				utils.ShowTemporaryMessage(s, "Please only use the keyboard keys advised")
			}
		}
	}
}
func handleSelection(index int, s tcell.Screen) {
	switch index {
	case 0:
		s.Clear()
		utils.ShowTemporaryMessage(s, "1")
	case 1:
		s.Clear()
		utils.ShowTemporaryMessage(s, "2")

	case 2:
		s.Clear()
		utils.ShowTemporaryMessage(s, "3")

	case 3:
		s.Clear()
		utils.ShowTemporaryMessage(s, "4")

	default:
		fmt.Println("Bye")
		s.Fini()
		os.Exit(0)
	}
}
