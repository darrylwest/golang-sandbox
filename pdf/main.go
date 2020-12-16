/**
 * https://github.com/jung-kurt/gofpdf
 *
 */
package main

import (
    "fmt"
    "github.com/jung-kurt/gofpdf"
)

func main() {
    // P = portrait, L = landscape; inch, mm; letter
    pdf := gofpdf.New("P", "in", "letter", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "", 16)
    pdf.Cell(4, 1, "Hello PDF world")
    err := pdf.OutputFileAndClose("hello.pdf")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("created pdf...")
}
