package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ChopStick struct {
	sync.Mutex
}

type Philosopher struct {
	ID              int
	LeftCs, RightCs *ChopStick
}

// - There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
// - Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
// - The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
// - In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
// - The host allows no more than 2 philosophers to eat concurrently.
// - Each philosopher is numbered, 1 through 5.
// - When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>”
//		on a line by itself, where <number> is the number of the philosopher.
// - When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>”
//		on a line by itself, where <number> is the number of the philosopher.
func main() {

	philosophers := make([]*Philosopher, 5)
	chopsticks := make([]*ChopStick, 5)

	for i := 0; i < 5; i++ {
		chopsticks[i] = new(ChopStick)
	}

	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i, chopsticks[i], chopsticks[(i+1)%5]}
	}

	var wg sync.WaitGroup
	philoToHostChannel := make(chan string)

	var hostToPhiloChannels [5]chan bool
	for i := range hostToPhiloChannels {
		hostToPhiloChannels[i] = make(chan bool)
	}

	go hostRoutine(&wg, &philoToHostChannel, &hostToPhiloChannels)
	for id := 0; id < 5; id++ {
		wg.Add(1)
		go philoRoutine(&wg, philosophers[id], &philoToHostChannel, &hostToPhiloChannels)
	}
	wg.Wait()
	close(philoToHostChannel)
}

func hostRoutine(wg *sync.WaitGroup, pToHostChannel *chan string, hostToPhiloChannels *[5]chan bool) {
	defer wg.Done()
	philoDinning := 0

	for idAction := range *pToHostChannel {
		parts := strings.Split(idAction, " ")
		philoID, _ := strconv.Atoi(parts[0])
		action := parts[1] // s->start or f->finished
		if action == "start" {
			if philoDinning < 2 {
				philoDinning++
				// can eat
				hostToPhiloChannels[philoID] <- true
			} else {
				fmt.Printf("--- P.%d needs to wait, %d P dining ---\n", philoID+1, philoDinning)
				hostToPhiloChannels[philoID] <- false
			}
		} else {
			philoDinning--
		}
	}
}

func philoRoutine(wg *sync.WaitGroup, p *Philosopher, pToHostChannel *chan string, hostToPhiloChannels *[5]chan bool) {
	defer wg.Done()
	timesEaten := 0
	for timesEaten < 3 {
		*pToHostChannel <- strconv.Itoa(p.ID) + " start"
		canEat := <-hostToPhiloChannels[p.ID]
		if canEat {
			p.LeftCs.Lock()
			p.RightCs.Lock()
			fmt.Printf("P.%d starting to eat\n", p.ID+1)
			time.Sleep(time.Second)
			timesEaten++
			fmt.Printf("p.%d finishing eating, eaten %d times\n", p.ID+1, timesEaten)
			p.LeftCs.Unlock()
			p.RightCs.Unlock()
			*pToHostChannel <- strconv.Itoa(p.ID) + " finish"
		} else {
			time.Sleep(1000 * time.Millisecond)
		}
	}
}
