package main

import "fmt"

type NewFamily struct {
	familyName string
	familySize int
	children   []NewChild
}

type NewChild struct {
	name string
	age  int
}

func main() {
	norouzieh := NewFamily{}
	norouzieh.familyName = "norouziehfamily"
	norouzieh.familySize = 5
	norouzieh.children = append(norouzieh.children, NewChild{"behrouz", 50})
	norouzieh.children = append(norouzieh.children, NewChild{"nahid", 46})
	norouzieh.children = append(norouzieh.children, NewChild{})
	norouzieh.children[2].name = "behzad"
	norouzieh.children[2].age = 39
	four := NewChild{"behnam", 35}
	norouzieh.children = append(norouzieh.children, four)

	fmt.Println(len(norouzieh.children), cap(norouzieh.children))
	norouzieh.children[0].age = 55
	fmt.Println(norouzieh)
}
