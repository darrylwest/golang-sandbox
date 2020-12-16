package main
 
import (
    "fmt"
    "container/list"
    "github.com/Pallinder/go-randomdata"
)

var people *list.List = list.New()
var center *list.Element

// search from front of list to find the correct insertion slot
func insertBefore(name string) *list.Element {
    for p := people.Front(); p != nil; p = p.Next() {
        if name < p.Value.(string) {
            fmt.Printf("insert %s before %s\n", name, p.Value)
            return people.InsertBefore(name, p)
        }
    }

    fmt.Printf("insert %s at back\n", name)
    return people.PushBack( name )
}

func populateList(n int) {
    for i := 0; people.Len() < n; i++ {
        name := randomdata.FirstName(randomdata.RandomGender)
        // name := randomdata.PostalCode("US") 
        if find(name) == nil {
            insertBefore(name)
        }
    }
}

func find(name string) *list.Element {
    for p := people.Front(); p != nil; p = p.Next() {
        if p.Value == name {
            return p
        }
    }

    return nil
}

func walkForward() {
    fmt.Print("Iterate Front/Next -> ")
    for p := people.Front(); p != nil; p = p.Next() {
        fmt.Printf("%s ", p.Value)
    }
    fmt.Println("")
}

func walkBackward() {
    fmt.Print("Iterate Back/Prev -> ")
    for p := people.Back(); p != nil; p = p.Prev() {
        fmt.Printf("%s ", p.Value)
    }
    fmt.Println("")
}

// use this to verify that the center is where it should be
func findCenter() *list.Element {

    target := people.Len() / 2

    p := people.Front()
    for count := 0; count < target; count++ {
        p = p.Next()
    }

    return p
}
 
func main() {
    populateList(5)
    walkForward()

    front := people.Front()
    center := findCenter()
    back := people.Back()

    fmt.Printf("\nfront: %s, center: %s back: %s\n", front.Value, center.Value, back.Value)
}

