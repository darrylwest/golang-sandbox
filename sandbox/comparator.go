package main

import (
    "fmt"
    "sort"
)

/*
    sort by descending score, ascending name
    score 0 <= score <= 1000
    player names may be the same, but are all lowercase alpha
*/

type Player struct {
    name string
    score int
}

var less func(p1, p2 *Player) bool

func lessByScore(p1, p2 *Player) bool {
    if p1.score == p2.score {
        return p1.name < p2.name
    }
    return p1.score > p2.score
}


type Players []*Player

func (p Players) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p Players) Len() int { return len(p) }

func (p Players) Less(i, j int) bool {
    return less(p[i], p[j])
}

func (p Players) show() {
    for _,v := range p {
        fmt.Println(v.name, v.score)
    }
}

func main() {
    less = lessByScore
    var players Players
    players = []*Player{
        &Player{ "amy", 100 },
        &Player{ "david", 100 },
        &Player{ "heraldo", 50 },
        &Player{ "aakansha", 75 },
        &Player{ "aleksa", 150 },
    }

    players.show()
    sort.Sort(players)
    fmt.Println("sorted...")
    players.show()
}

