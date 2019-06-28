package main

import (
	"fmt"
	"github.com/yamakiller/mgolua/mlua"
)


func main() {
	L := mlua.NewState()
	L.OpenLibs()
	defer L.Close()

	L.GetGlobal("_G")

	if L.IsNil(-1) {
		fmt.Printf("nil\n")
	} else {
		if L.IsTable(-1) {
			fmt.Print("table\n");
		} else {
			fmt.Printf("no table\n")
		}
	}

	var ispass bool
	fmt.Scanln(&ispass)
}
