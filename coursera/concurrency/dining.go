package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type stick struct {
	sync.Mutex
}

type phi struct {
	ate    int
	n      int
	lC, rC *stick
}

var emptyP phi

func (p phi) eat(c chan *phi) {
	if p.ate < 3 {
		p.lC.Lock()
		p.rC.Lock()
		fmt.Printf("Starting to eat : %d\n", p.n)
		p.ate++
		fmt.Printf("Finished eating : %d, for %d time.\n", p.n, p.ate)
		p.lC.Unlock()
		p.rC.Unlock()
		c <- &p
	} else {
		fmt.Printf("Philosopher %d done eating\n", p.n)
		c <- &emptyP
	}
}

func getLen(phis map[int]*phi) int {
	var len int
	for _, v := range phis {
		if v != &emptyP {
			len++
		}
	}
	return len
}

func getNextRandomPhi(phis map[int]*phi) *phi {
	i := rand.Intn(len(phis))
	for {
		if phis[i] != &emptyP {
			curP := phis[i]
			phis[i] = &emptyP
			return curP
		} else {
			i = (i + 1) % len(phis)
		}
	}
}

func host(phis map[int]*phi, wg *sync.WaitGroup) {
	ch := make(chan *phi, 2)
	for getLen(phis) > 0 {
		curP1 := getNextRandomPhi(phis)
		go curP1.eat(ch)
		// fmt.Printf("Current phis :")
		// for _, v := range phis {
		// 	fmt.Printf(" %d ", v.n)
		// }
		// fmt.Println()

		curP2 := getNextRandomPhi(phis)
		go curP2.eat(ch)
		// fmt.Printf("Current phis :")
		// for _, v := range phis {
		// 	fmt.Printf(" %d ", v.n)
		// }
		// fmt.Println()

		var add *phi
		add = <-ch
		// fmt.Printf("Added : %d\n", add.n)
		if add != &emptyP {
			phis[add.n-1] = add
		}
		// fmt.Printf("Current phis :")
		// for _, v := range phis {
		// 	fmt.Printf(" %d ", v.n)
		// }
		// fmt.Println()
		add = <-ch
		// fmt.Printf("Added : %d\n", add.n)
		if add != &emptyP {
			phis[add.n-1] = add
		}
		// fmt.Printf("Current phis :")
		// for _, v := range phis {
		// 	fmt.Printf(" %d ", v.n)
		// }
		// fmt.Println()
	}
	wg.Done()
}

func main() {

	sticks := make([]*stick, 5)
	for i := 0; i < 5; i++ {
		sticks[i] = new(stick)
	}

	phis := make(map[int]*phi, 5)
	for i := 0; i < 5; i++ {
		phis[i] = &phi{
			ate: 0,
			n:   i + 1,
			lC:  sticks[i],
			rC:  sticks[(i+1)%5],
		}
	}
	fmt.Printf("Init phis :")
	for _, v := range phis {
		fmt.Printf(" %d ", v.n)
	}
	fmt.Println()

	var wg sync.WaitGroup
	wg.Add(1)
	go host(phis, &wg)
	wg.Wait()
}
