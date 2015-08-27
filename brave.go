package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var level = [][]uint8{
	{},
	{255, 0, 4, 8, 12},
	{255, 1, 5, 9, 13},
	{255, 2, 6, 10, 14},
	{255, 3, 7, 11, 15},
}

func main() {
	bio := bufio.NewReader(os.Stdin)

	fmt.Println("Save Mel Gibsons Testicles")
	fmt.Println("By: Jesse Donat, Andrew Gross, Jeff Forshee")
	fmt.Println("Testicles Provided by: Nick Miller")
	fmt.Println("")

	b := 1
	c := 1
	pants := false
	breifs := false

	for {
		d := level[b][c]
		where(d, pants)

		line, _, _ := bio.ReadLine()
		n := len(line)
		a := string(line[:n])

		a = strings.ToLower(a)
		if a == "look" {
			look(d)
		}
		if a == "look at" {
			lookat(d)
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
			fmt.Println("go, look, look at, climb, take, and many more")
		}
		// MISC
		if a == "take kilt" && !pants && d == 4 {
			pants = true
			fmt.Println("huzzah, you have the kilt")
		}
		if a == "overthrow" {
			fmt.Println("You have overthrown england. as searching your new castle, in the treasure")
			fmt.Println("room you find a pair of Joe Nameth sling shot briefs")
			fmt.Println("*YOU HAVE RECIEVED SLING SHOT BRIEFS*")
			breifs = true
		}
		if a == "ride sombrero" {
			fmt.Println("In the process of riding the sombrero you lose your kilt and end up in Paris")
			pants = false
			b = 4
			c = 1
			if !breifs {
				fmt.Println("The space police notice you standing nude in the center of Paris,")
				fmt.Println("rather than cause a sceen they choose to proceed to shoot you in the")
				fmt.Println("testicles. You die from loss of blood.")
				dead(bio)
			}
		}
		// CLIMB
		if a == "climb fence" && !pants && d == 1 {
			fmt.Println("As you begin climbing the fence you feel a sharp pull.")
			fmt.Println("You continue to climb the fence as a loud noice comes from your crotch,")
			fmt.Println("then another from your mouth.")
			fmt.Println("As you look down you see your testicles have failed to make the climb with")
			fmt.Println("the rest of your body. They are hanging on the fence. ")
			fmt.Println("You slowly die from an agonizing pain and a severe amount of blood loss.")
			dead(bio)
		}
		if a == "climb fence" && pants && d == 1 {
			b = b + 1
		}
		// BOUNDRIES
		if a == "go east" {
			if d == 4 || d == 3 || d == 12 {
				fmt.Println("That is ocean, you fool")
				b = b - 1
			}
		}
		if a == "go west" {
			if d == 0 || d == 4 || d == 12 {
				fmt.Println("That is the ocean, you fool")
				b = b + 1
			}
		}
		if a == "go south" {
			if d == 4 || d == 1 || d == 2 || d == 3 || d == 12 {
				fmt.Println("That is the ocean, you fool")
				c = c - 1
			}
		}
		if a == "go north" {
			if d == 0 || d == 1 || d == 2 || d == 3 || d == 12 {
				fmt.Println("That is ocean, you fool")
				c = c + 1
			}
		}
		if a == "go east" && d == 1 {
			b = b - 1
			fmt.Println("their is a fence in the way, you must climb the fence")
		}
	}

}

func dead(bio *bufio.Reader) {
	for {
		bio.ReadLine()
		fmt.Println("You are dead, you jack ass")
	}
}

func look(d uint8) {
	if d == 0 {
		fmt.Println("east: fence  north: ocean  south: corpse")
	}
	if d == 1 {
		fmt.Println("east: england  north: ocean  south: ocean  west: where you were nude")
	}
	if d == 2 {
		fmt.Println("east: sombrero  north: ocean  south: exotic pelvic dancing  west: fence")
	}
	if d == 4 {
		fmt.Println("east: ocean  west: ocean  north: where you were naked  south: ocean")
	}

}

func lookat(d uint8) {
	if d == 0 {
		fmt.Println("It is a field")
	}
	if d == 1 {
		fmt.Println("You see a fence on which you could easily catch your testicles")
	}
	if d == 2 {
		fmt.Println("You see england in need of overthrowing")
	}
	if d == 3 {
		fmt.Println("You see a giant sombrero in need of riding")
	}
	if d == 4 {
		fmt.Println("Corpse is wearing YOUR kilt! It is the theiving corpse! You must take back")
		fmt.Println("what is rightfully yours!")
	}
}

func where(d uint8, pants bool) {
	if !pants {
		if d == 0 {
			fmt.Println("You are standing naked in a field")
		}
		if d == 1 {
			fmt.Println("You are standing naked at fence")
		}
		if d == 4 {
			fmt.Println("You are standing naked at a corpse")
		}
		if d == 12 {
			fmt.Println("You are standing in breifs in Paris")
		}
	}

	if pants {

		if d == 0 {
			fmt.Println("You are wearing a kilt in a field")
		}
		if d == 1 {
			fmt.Println("You are wearing a kilt at fence")
		}
		if d == 2 {
			fmt.Println("You are wearing a kilt in england")
		}
		if d == 3 {
			fmt.Println("You are wearing a kilt near a giant sombrero")
		}
		if d == 4 {
			fmt.Println("You are wearing a kilt at a naked corpse")
		}
	}
}
