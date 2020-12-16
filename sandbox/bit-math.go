package main

import "fmt"

/*\

Given the final position and score of a game player, return the number of coins and
extra points (xp) based on this...

  * coins: 2000 / 2 ^ pos
  * xp: 2 * score if position == 1, otherwise, the score

  1 <= position <= 20
  0 <= points <= 500
\*/

func codefightsTournament(position, score int) []int {
    // coin calc = 4000 right shifted position times, so 4000 >> 1 = 2000, 4000 >> 2 = 1000, etc
    // coins is left sifted 1 or zero times so 500 << 1 == 1000, 500 << 0 == 500

    return []int{4e3 >> uint(position), score + 1 / position * score }
}

type Data struct {
    position int
    score int
    result []int
}

func main() {
    inputs := []Data{
        Data{ 1, 500, []int{ 2000, 1000 } },
        Data{ 2, 400, []int{ 1000,  400 } },
        Data{ 3, 112, []int{  500,  112 } },
        Data{ 4, 400, []int{  250,  400 } },
        Data{ 5,  10, []int{  125,   10 } },
        Data{ 6, 100, []int{   62,  100 } },
    }

    tests := 0
    passes := 0
    for _, v := range inputs {
        tests++
        var passfail string

        r := codefightsTournament(v.position, v.score)

        if (v.result[0] == r[0] && v.result[1] == r[1]) {
            passfail = "ok"
            passes++
        } else {
            passfail = "FAILED!"
        }

        fmt.Println(v, r, passfail)
    }

    fmt.Printf("\n%d tests, %d passed, %d failed\n", tests, passes, tests - passes)
}
