package main

import (
    "fmt"
    "strings"
    "strconv"
)

/*
    (see 1.6)
    Implement a method to perform basic string compression using the counts of repeated characters.  For example the
    string "aabcccccaaa" => "a2b1c5a3".  If the compressed string's length is >= the string's length return the string.

    what is the big O time
*/

func compress(s string) string {
    // can't compress 2 or less characters
    if len(s) < 3 {
        return s
    }

    chars := strings.Split(s, "")
    buf := []string{ chars[0] }
    count := 1

    fmt.Println(chars, buf)

    for i := 1; i < len(chars); i++ {
        if chars[i-1] == chars[i] {
            count++
        } else {
            buf = append(buf, strconv.Itoa(count))
            count = 1
            buf = append(buf, chars[i])

            // quick exit...
            if len(buf) >= len(s) {
                return s
            }
        }

    }

    // the final count...
    buf = append(buf, strconv.Itoa(count))

    if len(buf) < len(s) {
        return strings.Join(buf, "")
    } else {
        return s
    }
}

func main() {
    s := "aabcccccaaa"
    expect := "a2b1c5a3"

    r := compress(s)
    fmt.Println(expect, r, expect == r)
}
