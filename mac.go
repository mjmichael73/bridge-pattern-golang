package main

import "fmt"

// Refined abstraction
type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}