package main

import (
    // "time"
    ui "github.com/gizak/termui"
)

func main() {
    err := ui.Init()
    if err != nil {
        panic(err)
    }
    defer ui.Close()

    p := ui.NewPar(":Press q to Quit demo")
    p.Height = 3
    p.Width = 50
    p.Y = 5
    p.TextFgColor = ui.ColorWhite
    p.BorderLabel = "Text Box"
    p.BorderFg = ui.ColorCyan

    g1 := ui.NewGauge()
    g1.Percent = 50
    g1.Width = 50
    g1.Height = 3
    g1.Y = 11
    g1.BorderLabel = "Aspects"
    g1.BarColor = ui.ColorRed
    g1.BorderFg = ui.ColorWhite
    g1.BorderLabelFg = ui.ColorCyan

    g2 := ui.NewGauge()
    g2.Percent = 80
    g2.Width = 50
    g2.Height = 3
    g2.Y = 41
    g2.BorderLabel = "Aspects"
    g2.BarColor = ui.ColorRed
    g2.BorderFg = ui.ColorWhite
    g2.BorderLabelFg = ui.ColorCyan

    ui.Render(p, g1, g2)

    ui.Handle("/sys/kbd/q", func(ui.Event) {
        ui.StopLoop()
    })

    ui.Loop()
}
