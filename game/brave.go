package game

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var level = [][]uint8{
	{},
	{255, 0, 4, 8, 12},
	{255, 1, 5, 9, 13},
	{255, 2, 6, 10, 14},
	{255, 3, 7, 11, 15},
}

type Game struct {
	bio          *bufio.Reader
	out          io.Writer
	bufferSize   int
	afterHeading func(in io.Reader, out io.Writer)
}

type Option func(*Game)

func WithBufferSize(size int) Option {
	return func(g *Game) {
		g.bufferSize = size
	}
}

func WithAfterHeading(fn func(in io.Reader, out io.Writer)) Option {
	return func(g *Game) {
		g.afterHeading = fn
	}
}

func NewGame(in io.Reader, out io.Writer, opts ...Option) *Game {
	g := &Game{
		out:        out,
		bufferSize: 4096, // default buffer size
	}

	// Apply options
	for _, opt := range opts {
		opt(g)
	}

	// Create reader with configured buffer size
	g.bio = bufio.NewReaderSize(in, g.bufferSize)

	return g
}

func (g *Game) Run() {

	fmt.Fprintln(g.out, "Save Mel Gibsons Testicles")
	fmt.Fprintln(g.out, "By: Jesse Donat, Andrew Gross, Jeff Forshee")
	fmt.Fprintln(g.out, "Testicles Provided by: Nick Miller")
	fmt.Fprintln(g.out, "")

	// Call afterHeading hook if provided
	if g.afterHeading != nil {
		g.afterHeading(g.bio, g.out)
	}

	b := 1
	c := 1
	pants := false
	breifs := false

	for {
		d := level[b][c]
		g.where(d, pants)

		line, _, err := g.bio.ReadLine()
		if err != nil {
			return
		}
		n := len(line)
		a := string(line[:n])

		a = strings.ToLower(a)
		if a == "look" {
			g.look(d)
		}
		if a == "look at" {
			g.lookat(d)
		}

		// WALK
		if a == "go east" {
			b = b + 1
		}
		if a == "go west" {
			b = b - 1
		}
		if a == "go south" {
			c = c + 1
		}
		if a == "go north" {
			c = c - 1
		}
		if a == "help" {
			fmt.Fprintln(g.out, "go, look, look at, climb, take, and many more")
		}
		// MISC
		if a == "take kilt" && !pants && d == 4 {
			pants = true
			fmt.Fprintln(g.out, "huzzah, you have the kilt")
		}
		if a == "overthrow" {
			fmt.Fprintln(g.out, "You have overthrown england. as searching your new castle, in the treasure")
			fmt.Fprintln(g.out, "room you find a pair of Joe Nameth sling shot briefs")
			fmt.Fprintln(g.out, "*YOU HAVE RECIEVED SLING SHOT BRIEFS*")
			breifs = true
		}
		if a == "ride sombrero" {
			fmt.Fprintln(g.out, "In the process of riding the sombrero you lose your kilt and end up in Paris")
			pants = false
			b = 4
			c = 1
			if !breifs {
				fmt.Fprintln(g.out, "The space police notice you standing nude in the center of Paris,")
				fmt.Fprintln(g.out, "rather than cause a sceen they choose to proceed to shoot you in the")
				fmt.Fprintln(g.out, "testicles. You die from loss of blood.")
				g.dead()
			}
		}
		// CLIMB
		if a == "climb fence" && !pants && d == 1 {
			fmt.Fprintln(g.out, "As you begin climbing the fence you feel a sharp pull.")
			fmt.Fprintln(g.out, "You continue to climb the fence as a loud noice comes from your crotch,")
			fmt.Fprintln(g.out, "then another from your mouth.")
			fmt.Fprintln(g.out, "As you look down you see your testicles have failed to make the climb with")
			fmt.Fprintln(g.out, "the rest of your body. They are hanging on the fence. ")
			fmt.Fprintln(g.out, "You slowly die from an agonizing pain and a severe amount of blood loss.")
			g.dead()
		}
		if a == "climb fence" && pants && d == 1 {
			b = b + 1
		}
		// BOUNDRIES
		if a == "go east" {
			if d == 4 || d == 3 || d == 12 {
				fmt.Fprintln(g.out, "That is ocean, you fool")
				b = b - 1
			}
		}
		if a == "go west" {
			if d == 0 || d == 4 || d == 12 {
				fmt.Fprintln(g.out, "That is the ocean, you fool")
				b = b + 1
			}
		}
		if a == "go south" {
			if d == 4 || d == 1 || d == 2 || d == 3 || d == 12 {
				fmt.Fprintln(g.out, "That is the ocean, you fool")
				c = c - 1
			}
		}
		if a == "go north" {
			if d == 0 || d == 1 || d == 2 || d == 3 || d == 12 {
				fmt.Fprintln(g.out, "That is ocean, you fool")
				c = c + 1
			}
		}
		if a == "go east" && d == 1 {
			b = b - 1
			fmt.Fprintln(g.out, "their is a fence in the way, you must climb the fence")
		}
	}

}

func (g *Game) dead() {
	for {
		_, _, err := g.bio.ReadLine()
		if err != nil {
			return
		}
		fmt.Fprintln(g.out, "You are dead, you jack ass")
	}
}

func (g *Game) look(d uint8) {
	if d == 0 {
		fmt.Fprintln(g.out, "east: fence  north: ocean  south: corpse")
	}
	if d == 1 {
		fmt.Fprintln(g.out, "east: england  north: ocean  south: ocean  west: where you were nude")
	}
	if d == 2 {
		fmt.Fprintln(g.out, "east: sombrero  north: ocean  south: exotic pelvic dancing  west: fence")
	}
	if d == 4 {
		fmt.Fprintln(g.out, "east: ocean  west: ocean  north: where you were naked  south: ocean")
	}

}

func (g *Game) lookat(d uint8) {
	if d == 0 {
		fmt.Fprintln(g.out, "It is a field")
	}
	if d == 1 {
		fmt.Fprintln(g.out, "You see a fence on which you could easily catch your testicles")
	}
	if d == 2 {
		fmt.Fprintln(g.out, "You see england in need of overthrowing")
	}
	if d == 3 {
		fmt.Fprintln(g.out, "You see a giant sombrero in need of riding")
	}
	if d == 4 {
		fmt.Fprintln(g.out, "Corpse is wearing YOUR kilt! It is the theiving corpse! You must take back")
		fmt.Fprintln(g.out, "what is rightfully yours!")
	}
}

func (g *Game) where(d uint8, pants bool) {
	if !pants {
		if d == 0 {
			fmt.Fprintln(g.out, "You are standing naked in a field")
		}
		if d == 1 {
			fmt.Fprintln(g.out, "You are standing naked at fence")
		}
		if d == 4 {
			fmt.Fprintln(g.out, "You are standing naked at a corpse")
		}
		if d == 12 {
			fmt.Fprintln(g.out, "You are standing in breifs in Paris")
		}
	}

	if pants {

		if d == 0 {
			fmt.Fprintln(g.out, "You are wearing a kilt in a field")
		}
		if d == 1 {
			fmt.Fprintln(g.out, "You are wearing a kilt at fence")
		}
		if d == 2 {
			fmt.Fprintln(g.out, "You are wearing a kilt in england")
		}
		if d == 3 {
			fmt.Fprintln(g.out, "You are wearing a kilt near a giant sombrero")
		}
		if d == 4 {
			fmt.Fprintln(g.out, "You are wearing a kilt at a naked corpse")
		}
	}
}
