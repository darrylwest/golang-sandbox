package main

import (
	"fmt"
    "strings"
)

/*
    replace the switch
*/

const (
    yes = "YES"
    no = "NO"
)

var refmap map[string]string = make(map[string]string)

type StackType interface{}
type Stack []StackType

func (s Stack) Empty() bool {
	return len(s) == 0
}

func (s *Stack) Push(v StackType) {
	(*s) = append(*s, v)
}

func (s *Stack) Pop() StackType {
	idx := len(*s) - 1
	v := (*s)[idx]

	(*s) = (*s)[0:idx]

	return v
}

func (s Stack) Peek() StackType {
    idx := len(s) - 1
    return s[idx]
}

func testBracketBalance(s string) string {
    if len(s) % 2 != 0 {
        return no
    }

    stack := Stack{}
    chars := strings.Split(s, "")


    for _, x := range chars {
        v, ok := refmap[x]
        if ok { // must be a closer
            // fmt.Printf("poped %s for %s\n", v, x)
            if stack.Empty() || v != stack.Pop() {
                return no
            }
        } else {
            stack.Push(x)
        }
    }

    if stack.Empty() {
        return yes
    } else {
        return no
    }
}

func initRefMap() {
    refmap["}"] = "{"
    refmap[")"] = "("
    refmap["]"] = "["
    // refmap[">"] = "<"
}

func testit() {
    n := 79
    ss := strings.Split(`[()][{}()][](){}([{}(())([[{}]])][])[]([][])(){}{{}{[](){}}}()[]({})[{}{{}([{}][])}]
    [()][{}[{}[{}]]][]{}[]{}[]{{}({}(){({{}{}[([[]][[]])()]})({}{{}})})}
    (])[{{{][)[)])(]){(}))[{(})][[{)(}){[(]})[[{}(])}({)(}[[()}{}}]{}{}}()}{({}](]{{[}}(([{]
    ){[]()})}}]{}[}}})}{]{](]](()][{))])(}]}))(}[}{{)}{[[}[]
    }(]}){
    ((]()(]([({]}({[)){}}[}({[{])(]{()[]}}{)}}]]{({)[}{(
    {}{({{}})}[][{{}}]{}{}(){{}[]}{}([[][{}]]())
    (){}[()[][]]{}(())()[[([])][()]{}{}(({}[]()))()[()[{()}]][]]
    ()([]({}[]){}){}{()}[]{}[]()(()([[]]()))()()()[]()(){{}}()({[{}][]}[[{{}({({({})})})}]])
    []([{][][)(])}()([}[}(})}])}))]](}{}})[]({{}}))[])(}}[[{]{}]()[(][])}({]{}[[))[[}[}{(]})()){{(]]){][
    {()({}[[{}]]()(){[{{}{[[{}]{}((({[]}{}()[])))]((()()))}(()[[[]]])((()[[](({([])()}))[]]))}]})}
    ()(){{}}[()()]{}{}
    {}()([[]])({}){({[][[][[()]]{{}[[]()]}]})}[](())((())[{{}}])
    {}(((){}){[]{{()()}}()})[]{{()}{(){()(){}}}}{()}({()(()({}{}()((()((([])){[][{()}{}]})))))})
    ][[{)())))}[)}}}}[{){}()]([][]){{{{{[)}]]{([{)()][({}[){]({{
    {{}(
    {[{((({}{({({()})()})[]({()[[][][]]}){}}))){}}]}{}{({((){{}[][]{}[][]{}}[{}])(())}[][])}
    ()[[][()[]][]()](([[[(){()[[]](([]))}]]]))
    ()[]({}{})(()){{{}}()()}({[]()}())[](){}(({()}[{}[{({{}}){({}){({})((({()})))}}}]]))
    }[{){({}({)})]([}{[}}{[(([])[(}){[]])([]]}(]]]]{][
    [{]{[{(){[}{}(([(]}])(){[[}(]){(})))}}{{)}}{}][({(}))]}({)
    )})[(]{][[())]{[]{{}}[)[)}[]){}](}({](}}}[}{({()]]
    [[[({[]}({[][[[[][[{(()[][])}()[][]][]{}]]]]}))][(()){}]]]()[{}([]{}){}{{}}]
    ({[]({[]})}())[][{}[{{(({{{([{}])}}}))}}]]
    ([((()))()])[][][]{}()(([]))[]()[]((){}[]){}(){{}[]}[[{[]}]]
    [[(((({}{[]{}()}){}{{}}){({[]{[{}]{(){}(((){()}))}()}}[[]]()()[()])[[{}{}]()]}))]]{}[]{}({({{}})})
    (]{()}((
    [][(())[({{{()[]}}{[[][[][[[]{{{[()]{{{{}{[]}[][]}}}}}}]]]]}})]]
    }[})})}[)]{}{)
    ({(}{})))}(}[)[}{)}}[)[{][{(}{{}]({}{[(})[{[({{[}{(]]})}
    ]}})[]))]{][])[}(])]({[]}[]([)
    [{{}{[{{[}[[}([]
    [([]){}][({})({[(([])[][])][[{}{([{{}{(()){{{({}{{}}())}}[]}}()[()[{{{([](()){[[[]]]})}}}]]}])}]]})]
    ]{}{(}))}](})[{]]()(]([}]([}][}{]{[])}{{{]([][()){{})[{({{{[}{}](]}}
    {[{}}){(}[][)(}[}][)({[[{]}[(()[}}){}{)([)]}(()))]{)(}}}][
    (]{}{(}}}[)[
    []{}{[[]]}([{}]{}[]){{(())}}
    [)([{(][(){)[)}{)]]}}([((][[}}(]{}]]}]][(({{{))[[){}{]][))[]{]][)[{{}{()]){)])))){{{[(]}[}}{}]
    {({(){[[[][]{}[[([]{})]{}]][[]()()]]}})}[{}{{}}]
    )}][(})){))[{}[}
    {[]{({]}[}}[{([([)([){{}{(}}[]}}[[{[}[[()(])[}[]
    ()()()[]
    ((){}])][]][}{]{)]]}][{]}[)(])[}[({(
    )[((])(]]]]((]){{{{())]}]}(}{([}(({}]])[[{){[}]{{}})[){(
    }][[{[((}{[]){}}[[[)({[)}]]}(]]{[)[]}{}(){}}][{()]))})]][(((}}
    ([]){}{{}{}}()([([{}{[[]()([(([]()))()[[]]])]}])])
    [()[[]{{[]}()([])}[]][][]][]()[]{}{}[][]{}{}[()(){}]
    {[{){]({(((({](]{([])([{{([])[}(){(]](]{[{[]}}())[){})}))[{})))[
    {}[()[]][]{}{}[[{{[[({})]()[[()]]]}}]]
    {[{}[][]]}[((()))][]({})[]{}{()}
    (){[{({})}]}
    ([]])][{)]({)[]))}]())[}]))][}{(}}})){]}]{[)}(][})[[
    ((({{}(([{}(())]))[()]{[[[]()]]}})))
    }()))}(}]]{{})}][{](]][{]{[[]]]}]]}([)({([))[[(]}])}[}(([{)[)]]([[](]}]}{]{{})[]){]}{])(
    {}{}{}{[[()]][]}
    )]}]({{})[[[{]{{{}}][))]{{
    ))){({}])}])}}]{)()(}(]}([
    ([[]][])[[]()][]()(([[]]{[()[]{[][{}]}[()]]{}{[]}}{{}()}(()[([][]{})[[{}][]]{}[]])))
    (]{[({}[){)))}]{[{}][({[({[]))}[}]}{()(([]{]()}})}[]{[)](((]]])([]}}]){)(([]]}[[}[
    ([[]])({}(([(){{}[{}]}]){[{}]}))[][{}{}](){}
    [][][][][][([])][]{({()}[[()()]{([(){[]{}}{(())}{[](){}()({}())}[({}[[]()])][]])}])}
    }[{{(}})}}(((())()({]{([]((][(({)[({[]]}[])}]{][{{}]{)][}(])}}}))}}}
    []({})()[]{}{}[]({}{})[]{([])()[()][{()({})[{}{[[()]{}[]][]}(({{[]{()()()}{}[]()}[]}){{}{}})]}]}
    {{(([{)]{}({][{](){({([[[][)}[)})(
    [{}]{[()({[{}]})]}
    [[{}]]
    ]{{({[{]}[[)]]}{}))}{){({]]}{]([)({{[]){)]{}){){}()})(]]{{])(])[]}][[()()}
    {[([}[[{{(]]][}()())]{){(){)]]){})}]{][][(}[]())[}[)})})[][{[)[})()][]))}[[}
    ]()])}[}}}{]]{)[}(}]]])])}{(}{([{]({)]}])(})[{}[)]])]}[]{{)){}{()}]}((}}{({])[}])[]}
    (]}[{}{{][}))){{{([)([[])([]{[
    {(()[]){}}(){[]}({{}(()())})([]){}{}(())()[()]{}()
    {{}[{}[{}[]]]}{}({{[]}})[[(){}][]]{}(([]{[][]()()}{{{()()}{[]}({}[]{()})}{()}[[]][()]}))
    {[][]}[{}[](){}]{{}{[][{}]}}
    ()(){}(){((){}[])([[]]())}
    {}[[{[((}[(}[[]{{]([(}]][[
    {}[([{[{{}()}]{}}([[{}[]]({}{{()}[][][]})])])]
    {[](}([)(])[]]})()]){[({]}{{{)({}(][{{[}}(]{
    [][]{{}[](())}{}({[()]}())[][[][({}([{}]))]]
    ((()))[]{[(()({[()({[]}{})]}))]}{[]}{{({}{})[{}{}]{()([()])[{()}()[[]{}()]{}{}[]()]}[[]{[]}([])]}}`,"\n")

    expect := strings.Split(`YES
    YES
    NO
    NO
    NO
    NO
    YES
    YES
    YES
    NO
    YES
    YES
    YES
    YES
    NO
    NO
    YES
    YES
    YES
    NO
    NO
    NO
    YES
    YES
    YES
    YES
    NO
    YES
    NO
    NO
    NO
    NO
    YES
    NO
    NO
    NO
    YES
    NO
    YES
    NO
    NO
    YES
    NO
    NO
    NO
    YES
    YES
    NO
    YES
    YES
    YES
    NO
    YES
    NO
    YES
    NO
    NO
    YES
    NO
    YES
    YES
    NO
    YES
    NO
    YES
    YES
    NO
    NO
    NO
    NO
    YES
    YES
    YES
    YES
    NO
    YES
    NO
    YES
    YES`, "\n")

    if len(ss) != n || len(expect) != n {
        panic("bracket array length is incorrect")
    }

    errors := 0
    for i, v := range ss {
        s := strings.Trim(v, " ")
        ex := strings.Trim(expect[i], " ")
        r := testBracketBalance(s)
        if r != ex {
            fmt.Printf("ERROR item %d %s expected %s got %s\n", i, s, expect[i], r)
            errors++
        }
    }
    if errors == 0 {
        fmt.Printf("Tests passed for %d samples...\n", len(ss))
    } else {
        fmt.Printf("%d Tests, %d FAILED!\n", len(ss), errors)
    }
}

func main() {
    initRefMap()
    A := []string{ "{[()]}", "{[(])}", "{{[[(())]]}}" }
    for _,s := range A {
        fmt.Println(testBracketBalance(s))
    }

    testit()
}
