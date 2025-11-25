package main

import (
	"fmt"

	"github.com/lfhy/flag"
)

var (
	IntData    int
	StringData string
)

func init() {
	flag.Var(&flag.FlagVar{
		Value: &IntData,
		Name:  "i",
	})
	flag.Var(&flag.FlagVar{
		Value: &StringData,
		Name:  "s",
	})
}

func main() {
	flag.Parse()
	fmt.Printf("IntData: %v\n", IntData)
	fmt.Printf("StringData: %v\n", StringData)
}
