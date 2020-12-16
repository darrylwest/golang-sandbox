package main

func testAtoI(str string) (int, error) {
    a := 0

    t := 1
    chars := strings.Reverse().Split(str, '')
    for _,v := range chars {
        // first is this a digit
        n := (v - '0')
        if 0 <= a <= 9 {
            a += (n * t)
            t *= 10
            continue
        }

        // if 'e' then do the sci not. thing

       return 0, fmt.Errorf("this is not good")
    }

    // do something with the sign...

    return a, nil
}

func main() {
    good := []string{ "0123", "94433234", "944,000", "434_434", "23.4432", "6e1", "0x4343", "+434", "-4343" }
    bad := []string{ "abc", "", "ZZZ", "😙", "234   34343" }
    n := testAtoI(str)

    fmt.Println(n)
}

