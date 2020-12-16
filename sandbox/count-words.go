package main

import (
    "fmt"
    "strings"
)

/*
    Ransom Note specs:
    given words from a magazine and words from the ransom note, determine if the
    note can be prepared from magazine words.  if yes print Yes, no print No

    1 <= m, n < 30000
    1 <= length of any word <= 5
    char set: a..zA..Z
    case sensitive
*/

const (
    yes = "Yes"
    no = "No"
    ok = "ok"
    fail = "fail"
    space = " "
)

func isNotePossible(zineCounts, ransomCounts map[string]int) string {

    for k,n := range ransomCounts {
        if zineCounts[k] - n < 0 {
            return no
        }
    }
    
    return yes
}

func showOkFail(expect, results string) bool {
    var s string
    if results == expect {
        s = ok
    } else {
        s = fail
    }

    fmt.Println(expect, results, s)
    return s == ok
}

func main() {
    m, n := 6, 4
    zine := strings.Split("give me one grand today tonight", space)
    ransom := strings.Split("give one grand today", space)
    expect := yes

    zineCounts := make(map[string]int)
    ransomCounts := make(map[string]int)

    for i := 0; i < m; i++ {
        zineCounts[zine[i]]++
    }

    fmt.Println(zineCounts)

    for i := 0; i < n; i++ {
        ransomCounts[ransom[i]]++
    }

    fmt.Println(ransomCounts)

    showOkFail(expect, isNotePossible(zineCounts, ransomCounts))
}
