package Menus

import (
	"fmt"

	"github.com/Moukhtar-youssef/Go-VIM-games.git/utils"
	"github.com/gdamore/tcell/v2"
)

func OptionsMenu(s tcell.Screen) {
	if s == nil {
		fmt.Println("the screen passed down is nil")
	}
	s.Clear()
	menuitems := []utils.MenuItems{
		{
			Name: "Graphic",
		},
		{
			Name: "User Data",
		},
	}
	menu := utils.NewMenu("Option Menu",menuitems,s)
	menu.HandleInput()
}
