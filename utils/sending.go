package lemin

import "fmt"

type Ant struct {
	path   []string
	roomId int
	prev   string
	pass   bool
}

func Send(g [][]string, ant int, c Colony, text []string) {
	levels := ant / len(g)
	if ant%len(g) != 0 {
		levels++
	}

	var allAnts = make([]Ant, ant+1)

	allAnts[0].pass = true
	id := 0

	for j := 0; j < levels; j++ {
		for _, path := range g {
			id++
			allAnts[id].path = path
			allAnts[id].roomId = 0
			allAnts[id].pass = false
			if id == ant {
				break
			}
		}
	}

	for _, v := range text {
		fmt.Println(v)
	}
	fmt.Println()

	exit := false
	var taken = make(map[string]bool)
	for !exit {
		for id, curAnt := range allAnts {
			if curAnt.pass {
				continue
			}
			room := curAnt.path[curAnt.roomId]
			if taken[room] {
				fmt.Println()
				break
			}
			fmt.Print("L", id, "-", room, " ")
			if id == ant {
				fmt.Println()
				if room == c.End {
					exit = true
				}
			}
			allAnts[id].roomId++
			taken[allAnts[id].prev] = false
			if room != c.End {
				taken[room] = true
				allAnts[id].prev = room
			}
			if room == c.End {
				allAnts[id].pass = true
			}
		}
	}
}
