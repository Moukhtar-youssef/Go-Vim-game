package Menus

import (
	"fmt"

	"github.com/Moukhtar-youssef/Go-VIM-games.git/Games"
	"github.com/Moukhtar-youssef/Go-VIM-games.git/utils"
	"github.com/gdamore/tcell/v2"
)

func GamesMainMenu(s tcell.Screen) {
	if s == nil {
		fmt.Println("The tcell screen value is nil")
	}
	s.Clear()
	menuitems := []utils.MenuItems{
		{
			Name: "Delete the Target",
			Function: func() {
				Games.DeleteTheTargetGame(s)
			},
		},
	}
	menu := utils.NewMenu("Game Menu", menuitems, s)
	menu.HandleInput()
}
