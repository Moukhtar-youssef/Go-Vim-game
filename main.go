package main

import (
	"fmt"
	"os"

	Games "github.com/Moukhtar-youssef/Go-VIM-games.git/games"
	"github.com/Moukhtar-youssef/Go-VIM-games.git/utils"
	"github.com/gdamore/tcell/v2"
)

type player struct {
	Name  string
	Score int
}

var GamesArray = []string{
	"HJKL Maze", "Insert Invader", "Delete the Virus", "Exit",
}

func main() {
	screen, err := tcell.NewScreen()
	if err != nil || screen == nil {
		fmt.Fprintf(os.Stderr, "Error creating screen: %v \n", err)
		os.Exit(1)
	}
	if err := screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing screen: %v \n", err)
		os.Exit(1)
	}
	defer func() {
		if screen != nil {
			screen.Fini()
		}
	}()

	selected := 0
	screen.Clear()
	for {
		utils.DrawMenu(screen, selected, GamesArray)
		screen.Show()

		event := screen.PollEvent()
		switch ev := event.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyUp:
				if selected > 0 {
					selected--
				}
			case tcell.KeyDown:
				if selected < len(GamesArray)-1 {
					selected++
				}
			case tcell.KeyEnter:
				handleSelection(selected, screen)
				screen.Clear()
				utils.DrawMenu(screen, selected, GamesArray)
				screen.Show()
			case tcell.KeyEscape:
				return
			default:
				utils.ShowTemporaryMessage(screen, "Please only use the keyboard keys advised")
			}
		}
	}
}

func handleSelection(index int, s tcell.Screen) {
	switch index {
	case 0:
		utils.PrintCentered(s, 40, 10, "Starting HJKL Maze", tcell.StyleDefault)
		Games.DeleteTheTargetGame(s)
	case 1:
		s.Clear()
		utils.ShowTemporaryMessage(s, "Comming soon")

	case 2:
		s.Clear()
		utils.ShowTemporaryMessage(s, "Comming soon")

	default:
		fmt.Println("Bye")
		s.Fini()
		os.Exit(0)
	}
}
