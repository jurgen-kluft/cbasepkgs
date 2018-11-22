package main

import (
	"github.com/jurgen-kluft/xbasepkgs/package"
	"github.com/jurgen-kluft/xcode"
)

func main() {
	xcode.Init()
	xcode.Generate(xbasepkgs.GetPackage())
}
