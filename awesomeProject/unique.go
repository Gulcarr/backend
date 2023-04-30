package main

import (
	"awesomeProject/parser"
	"awesomeProject/printer"
	"awesomeProject/reader"
	"awesomeProject/structs"
	"awesomeProject/unification"
	"fmt"
)

func main() {
	fmt.Println(unification.ParsingUnification(structs.FlagSet{"c", 0, 0, true}, []string{
		"Of course Adi loves music.",
		"Of course Adi lovEs musIc.",
		"It's true Bob LOVES music.",
		"C love music.",
		"",
		"I love and music of Kartik.",
		"We love per music of Kartik.",
		"Thanks.",
	}))
	flagSet, err := parser.ParseFlags()
	if err != nil {
		fmt.Println(err)
		return
	}
	inpfile, outfile, err := parser.ParseSources()
	if err != nil {
		fmt.Println(err)
		return
	}
	args := reader.Read(inpfile)
	result := unification.ParsingUnification(flagSet, args)
	printer.Print(result, outfile)
}
