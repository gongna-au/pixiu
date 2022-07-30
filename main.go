package main

import (
	"fmt"

	//"github.com/pixiu/global"
	"gopkg.in/alecthomas/kingpin.v2"
)

var ()

func main() {
	fmt.Print("test")
}
func CollectorFlagAction(s string) func(ctx *kingpin.ParseContext) error {
	return func(ctx *kingpin.ParseContext) error {
		fmt.Print(s)
		return nil
	}

}
