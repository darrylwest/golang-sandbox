#

## go doc <package>

## declare then use, e.g.

type Greeter struct {
    Format string
}
func (g Greeter) Greet(name string) {
    if g.Format == "" {
        g.Format = "Hi %s"
    }
    return fmt.Sprintf(g.Format, name)
}

func main() {
    g := Greeter{
        Format: "Howdy %s",
    }
    fmt.Println(g.Greet())
}

## avoid constructors if possible

## avoid writing interfaces--let the user do that

## reduce the number of dependencies (3rd party packages)

## third party

* go meta linter
* go report card
* go fmt


###### darryl.west@raincitysoftware.com | Version 2017.02.02
