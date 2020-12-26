package gotprint

import "fmt"

func init() {
	SetDefaultFormat()
}

func Sprint(s interface{}) (result string) {
	sm := ToStringMatrix(s)
	result = sm.String()
	return
}

func Print(s interface{}) {
	sm := ToStringMatrix(s)
	fmt.Print(sm.String())
}
