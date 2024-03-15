package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "is dirty")
		}
	}
}
func main() {
	rooms := [...]Room{
		{name: "office"},
		{name: "ware house"},
		{name: "Reception"},
		{name: "Ops"},
	}
	checkCleanness(rooms)
	fmt.Println("performing cleaning")

	rooms[2].cleaned = true
	rooms[3].cleaned = true
	checkCleanness(rooms)

}
