package objectcomposition

func objectcomposition() {
	/* penulisan cara pertama
	emp := enhancedMessagePrinter{}
	emp.message = "foo" */

	emp := enhancedMessagePrinter{messagePrinter{"foo"}}
	emp.printMessage()
}

type messagePrinter struct {
	message string
}

func (mp *messagePrinter) printMessage() {
	println(mp.message)
}

type enhancedMessagePrinter struct {
	messagePrinter
}
