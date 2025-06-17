package main

import (
	"fmt"
	"os"

	"github.com/Moukhtar-youssef/Go-VIM-games.git/Menus"
	"github.com/Moukhtar-youssef/Go-VIM-games.git/utils"
	"github.com/gdamore/tcell/v2"
)

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
	items := []utils.MenuItems{
		{
			Name: "Games Menu",
			Function: func() {
				Menus.GamesMainMenu(screen)
			},
		},
		{
			Name: "Leader Board",
			Function: func() {
				fmt.Println("Starting the leader board")
			},
		},
		{
			Name: "Options",
			Function: func() {
				Menus.OptionsMenu(screen)
			},
		},
		{
			Name: "Exit",
			Function: func() {
				screen.Fini()
				os.Exit(0)
			},
		},
	}

	screen.Clear()
	menu := utils.NewMenu("Start menu", items, screen)
	menu.HandleInput()
}
