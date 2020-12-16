package main

import (
    "fmt"
    "strings"
)

/*
    1 <= |a|,|b| <= 1e4
*/

func countChars(a []string, counts []int) {
    acode := []rune("a")[0]
    for _, v := range a {
        idx := []rune(v)[0] - acode
        counts[idx]++
    }
}

func abs(a int) int {
    if a < 0 {
        return a * -1
    }
    return a
}
func countCommonChars(a, b []string) int {
    A := make([]int, 26)
    B := make([]int, 26)

    countChars(a, A)
    countChars(b, B)

    count := 0
    for i := 0; i < 26; i++ {
        count += abs(A[i] - B[i])
    }

    // return the number of chars to delete...
    return count
}

func showResults(expect, count int) {
    var ok string
    if count == expect {
        ok = "ok"
    } else {
        ok = "FAILED!"
    }

    fmt.Printf("expect: %d got %d %s\n", expect, count, ok)
}

func main() {
    a := []string{ "c", "d", "e" }
    b := []string{ "a", "b", "c" }
    expect := 4

    fmt.Println("inputs:", a, b)
    showResults(expect, countCommonChars(a, b))


    s1 := "imkhnpqnhlvaxlmrsskbyyrhwfvgteubrelgubvdmrdmesfxkpykprunzpustowmvhupkqsyjxmnptkcilmzcinbzjwvxshubeln"
    s2 := "wfnfdassvfugqjfuruwrdumdmvxpbjcxorettxmpcivurcolxmeagsdundjronoehtyaskpwumqmpgzmtdmbvsykxhblxspgnpgfzydukvizbhlwmaajuytrhxeepvmcltjmroibjsdkbqjnqjwmhsfopjvehhiuctgthrxqjaclqnyjwxxfpdueorkvaspdnywupvmy"

    a = strings.Split(s1, "")
    b = strings.Split(s2, "")
    expect = 102

    showResults(expect, countCommonChars(a, b))
}
