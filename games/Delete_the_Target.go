package Games

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
	"time"

	"github.com/Moukhtar-youssef/Go-VIM-games.git/utils"
	"github.com/gdamore/tcell/v2"
)

var HJKLLevels = []string{
	"Easy", "Medium", "Hard", "Extreme",
}

type level struct {
	Name     string
	Finished bool
}

var levels []level

type Point struct {
	x, y int
}

type Game struct {
	screen         tcell.Screen
	playerX        int
	playerY        int
	gameOver       bool
	level          int
	targets        []Point
	DeletedTargets int
	targetGoal     int
	dPressedOnce   bool
}

const (
	width  = 50
	height = 20
)

func InitializingLevels() {
	levels = []level{
		{Name: "Easy", Finished: false},
		{Name: "Medium", Finished: false},
		{Name: "Hard", Finished: false},
		{Name: "Extreme", Finished: false},
	}
}
func NewGame(level int, s tcell.Screen) *Game {
	game := &Game{
		screen:         s,
		playerX:        width / 2,
		playerY:        height / 2,
		level:          level,
		targets:        []Point{},
		DeletedTargets: 0,
		targetGoal:     (level + 1) * 10,
		gameOver:       false,
	}
	game.AddTarget()
	return game
}
func (g *Game) Draw() {
	g.screen.Clear()
	for y := range height {
		for x := range width {
			if x == 0 || x == width-1 || y == 0 || y == height-1 {
				g.screen.SetContent(x, y, '#', nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
			}
		}
	}
	for _, target := range g.targets {
		g.screen.SetContent(target.x, target.y, 'X', nil, tcell.StyleDefault.Foreground(tcell.ColorRed))
	}
	g.screen.SetContent(g.playerX, g.playerY, '@', nil, tcell.StyleDefault.Foreground(tcell.ColorGreen))
	g.DisplayCounter()
	g.screen.Show()
}

func (g *Game) AddTarget() {
	numtarget := g.level + 1
	for range numtarget {
		for {
			x := rand.Intn(width-2) + 1
			y := rand.Intn(height-2) + 1
			distance := int(math.Abs(float64(g.playerX-x)) + math.Abs(float64(g.playerY-y)))
			valid := false
			switch g.level {
			case 0:
				valid = (distance <= 5)
			case 1:
				valid = (distance >= 6 && distance <= 10)
			case 2:
				valid = (distance >= 11 && distance <= 15)
			case 3:
				valid = (distance >= 15)
			}
			if valid {
				g.targets = append(g.targets, Point{x, y})
				break
			}
		}
	}
}
func (g *Game) DeleteTarget() {
	for i, target := range g.targets {
		if target.x == g.playerX && target.y == g.playerY {
			// Remove the target from the list and create a new one
			g.targets = slices.Delete(g.targets, i, i+1)
			g.DeletedTargets++
			break
		}
	}
}

func (g *Game) HandleInput() {
	ev := g.screen.PollEvent()
	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyESC:
			g.gameOver = true
		case tcell.KeyRune:
			switch ev.Rune() {
			case 'h':
				if g.playerX > 1 {
					g.playerX--
					g.dPressedOnce = false
				}
			case 'j':
				if g.playerY < height-2 {
					g.playerY++
					g.dPressedOnce = false
				}
			case 'k':
				if g.playerY > 1 {
					g.playerY--
					g.dPressedOnce = false
				}
			case 'l':
				if g.playerX < width-2 {
					g.playerX++
					g.dPressedOnce = false
				}
			case 'd':
				if g.dPressedOnce {
					g.DeleteTarget()
					g.dPressedOnce = false
				} else {
					g.dPressedOnce = true
				}
			}
		default:
			g.dPressedOnce = false
		}
	}
}

func (g *Game) DisplayCounter() {
	message := fmt.Sprintf("Targets Deleted: %d / %d", g.DeletedTargets, g.targetGoal)
	startx := len(message)
	starty := height + 1
	utils.PrintCentered(g.screen, startx, starty, message, tcell.StyleDefault.Foreground(tcell.ColorWhite))
}
func (g *Game) Run() {
	for !g.gameOver {
		g.Draw()
		g.HandleInput()
		if g.DeletedTargets >= g.targetGoal {
			g.gameOver = true
			g.screen.Clear()
			utils.ShowTemporaryMessage(g.screen, "yay the {}  level is done")
		}
		if len(g.targets) == 0 && !g.gameOver {
			g.AddTarget()
		}
	}
}
func DeleteTheTargetGame(s tcell.Screen) {
	InitializingLevels()
	items := []utils.MenuItems{
		{
			Name: "Easy",
			Function: func() {
				handleSelection(0, s)
			},
		},
		{
			Name: "Medium",
			Function: func() {
				handleSelection(1, s)
			},
		},
		{
			Name: "Hard",
			Function: func() {
				handleSelection(2, s)
			},
		},
		{
			Name: "Extreme",
			Function: func() {
				handleSelection(3, s)
			},
		},
	}
	Menu := utils.NewMenu("Choose level", items, s)
	s.Clear()
	Menu.HandleInput()
}

func handleSelection(index int, s tcell.Screen) {
	level := levels[index]
	s.Clear()
	utils.PrintCentered(s, 10, 5, fmt.Sprintf("Starting level: %s", level.Name), tcell.StyleDefault)
	time.After(2 * time.Second)
	game := NewGame(index, s)
	game.Run()

}
