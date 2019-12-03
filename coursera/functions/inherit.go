package main

import(
	"fmt"
	"strings"
)

type Animal struct {
	food string
	locomotion string
    noise string
}

func (a Animal) Eat()  {
	fmt.Println(a.food)
}

func (a Animal) Move()   {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak()   {
	fmt.Println(a.noise)
}

func main ()  {

	animals := [3]Animal{
		Animal{"grass", "walk", "moo"},
		Animal{"worms", "fly", "peep"},
		Animal{"mice", "slither", "hsss"},
	}

	for {
		var an,ac string
		fmt.Print("Enter query: ")
		fmt.Scanf("%s %s", &an, &ac)

		var a Animal
		switch strings.ToLower(an) {
			case "cow" : a = animals[0]
			case "bird" : a = animals[1]
			case "snake" : a = animals[2]
		}

		switch strings.ToLower(ac) {
			case "eat" : a.Eat()
			case "move" : a.Move()
			case "speak" : a.Speak()
		}
	}
}