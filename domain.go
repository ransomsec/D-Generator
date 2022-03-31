package main

import (
	"fmt"
	"time"

	gen "github.com/cs-5/rsdga"
	"github.com/fatih/color"
)

func Detail() {

	var detail = `
  This is random Domain Generator
  with help of Present-Time/Year/Month/Date

  by: @ransomsec

  [*] If any thing bad, please ping me, because i'am student & this is only for learning purpose.
  `
	Green.Println(detail)
}

var (
	Red   = color.New(color.FgRed).Add(color.Bold)
	Green = color.New(color.FgGreen).Add(color.Bold)
)

func main() {

	var lenght int

	Detail()
	fmt.Println("How many number of domains do you want to Generate??")
	Green.Print(">> ")
	fmt.Scan(&lenght)

	// Generator Process Start
	t := time.Now()
	gen := gen.New(int(t.Year()), int(t.Month()), int(t.Day()), "com")

	for i := 0; i < lenght; i++ {
		Red.Println(gen.Next())

	}

	switch {
	case lenght <= 0: // Checking value in -
		Red.Println("Oops Nothing Generated")
	default:
		Green.Printf("\n[*] Total %v number of domains are successfully generated! ;)\n", lenght)
	}

}
