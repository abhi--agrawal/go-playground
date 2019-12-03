package main

import(
	"fmt"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

func (c Cow) Eat()  {
	fmt.Println( "grass")
}

func (c Cow) Move()  {
	fmt.Println(  "walk")
}

func (c Cow) Speak()  {
	fmt.Println(  "moo")
}

type Cow struct {
}


func (b Bird) Eat()  {
	fmt.Println(  "worms")
}

func (b Bird) Move()  {
	fmt.Println(  "fly")
}

func (b Bird) Speak()  {
	fmt.Println(  "peep")
}

type Bird struct {
}

func (s Snake) Eat()  {
	fmt.Println(  "mice")
}

func (s Snake) Move()  {
	fmt.Println(  "slither")
}

func (s Snake) Speak()  {
	fmt.Println(  "hsss")
}

type Snake struct {
}

func main ()  {

	animals :=  map[string]Animal{}
	for {
		var ty, n,ac string
		fmt.Print("Enter command: ")
		fmt.Scanf("%s %s %s",&ty, &n, &ac)

		switch strings.ToLower(ty) {
			case "newanimal": 
				var a Animal
				switch strings.ToLower(ac) {
					case "cow" : a = Cow{}
					case "bird" : a = Bird{}
					case "snake" : a = Snake{}
				}
				animals[n] = a
				fmt.Println("Created it!")
			
			case "query" :
				a := animals[n]
				switch strings.ToLower(ac) {
					case "eat" : a.Eat()
					case "move" : a.Move()
					case "speak" : a.Speak()
				}
		}
	}
}