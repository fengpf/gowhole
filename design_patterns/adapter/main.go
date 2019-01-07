package main

import "fmt"

func main() {
	msg := "Hello World!"
	adapter := PrinterAdaper{
		OldPrinter: &MyLegacyPrinter{},
		Msg:        msg,
	}

	newMsg := adapter.PrintStored()
	if newMsg != "Legacy Printer:Adapter:Hello World!\n" {
		fmt.Printf("Msg didn't match %s\n", newMsg)
	}

	adapter2 := PrinterAdaper{
		OldPrinter: nil,
		Msg:        msg,
	}

	newMsg2 := adapter2.PrintStored()
	if newMsg2 != "Hello World!" {
		fmt.Printf("Msg didn't match %s\n", newMsg2)
	}
}

type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPrinter struct{}

func (l *MyLegacyPrinter) Print(msg string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy Printer:%s\n", msg)
	println(newMsg)
	return
}

type ModernPrinter interface {
	PrintStored() string
}

type PrinterAdaper struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrinterAdaper) PrintStored() (newMsg string) {
	//implement blow can make something about that
	//use the old LegacyPrinter interface by using this Adapter while
	//we use the ModernPrinter interface for future implementations.
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter:%s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
	} else {
		newMsg = p.Msg
	}
	return
}
