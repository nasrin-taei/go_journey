package main

import "fmt"

type Whisperer interface {
	Whisper() string
}

type Screamer interface {
	Scream() string
}

type Talker interface {
	Whisperer
	Screamer
	Speak() string
}

type Human struct {
	name string
}

func (h Human) Whisper() string {
	return "The human is whispering..."
}

func (h Human) Scream() string {
	return "The human is Screaming..."
}

func (h Human) Speak() string {
	return "The human is speaking..."
}

func (h Human) getName() string {
	return h.name
}

func main() {
	h1 := Human{}
	fmt.Println(h1.Whisper())
	fmt.Println(h1.Scream())
	fmt.Println(h1.Speak())

	var newTalker Talker
	newTalker = h1

	fmt.Println(newTalker.Speak())

	var newWhisperer Whisperer
	newWhisperer = h1
	fmt.Println(newWhisperer.Whisper())

	newScreamer := newWhisperer.(Screamer) //Type Conversion
	fmt.Println(newScreamer.Scream())

	myTalker, ok := newScreamer.(Talker) //this line asserts and then converts to the target type
	if ok {
		fmt.Println(myTalker.Whisper())
	}

	hum := newScreamer.(Human)
	fmt.Println(hum.getName())
}
