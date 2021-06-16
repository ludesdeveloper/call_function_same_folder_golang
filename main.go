package main

import (
	"call_function/func_print"
	"call_function/func_print/func_print_capital"
	"fmt"
)

func main() {
	fmt.Println("This is from Main")
	func_print.Func_Print()
	func_print.Question_Mark()
	func_print_capital.Print_Capital()
}
