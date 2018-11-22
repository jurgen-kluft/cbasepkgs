package xbasepkgs

import (
	"github.com/jurgen-kluft/xactor/package"
	"github.com/jurgen-kluft/xallocator/package"
	"github.com/jurgen-kluft/xbase/package"
	"github.com/jurgen-kluft/xbinmap/package"
	"github.com/jurgen-kluft/xcmdline/package"
	"github.com/jurgen-kluft/xcode/denv"
	"github.com/jurgen-kluft/xcompress/package"
	"github.com/jurgen-kluft/xcore/package"
	"github.com/jurgen-kluft/xcrypto/package"
	"github.com/jurgen-kluft/xentry/package"
	"github.com/jurgen-kluft/xfilesystem/package"
	"github.com/jurgen-kluft/xgenerics/package"
	"github.com/jurgen-kluft/xhash/package"
	"github.com/jurgen-kluft/xlang/package"
	"github.com/jurgen-kluft/xp2p/package"
	"github.com/jurgen-kluft/xrandom/package"
	"github.com/jurgen-kluft/xsocket/package"
	"github.com/jurgen-kluft/xstring/package"
	"github.com/jurgen-kluft/xsystem/package"
	"github.com/jurgen-kluft/xtext/package"
	"github.com/jurgen-kluft/xthread/package"
	"github.com/jurgen-kluft/xtime/package"
	"github.com/jurgen-kluft/xunittest/package"
	"github.com/jurgen-kluft/xuuid/package"
)

// GetPackage returns the package object of 'xbase'
func GetPackage() *denv.Package {

	// Dependencies
	actorpkg := xactor.GetPackage()
	allocatorpkg := xallocator.GetPackage()
	basepkg := xbase.GetPackage()
	binmappkg := xbinmap.GetPackage()
	cmdlinepkg := xcmdline.GetPackage()
	compresspkg := xcompress.GetPackage()
	corepkg := xcore.GetPackage()
	cryptopkg := xcrypto.GetPackage()
	entrypkg := xentry.GetPackage()
	filesystempkg := xfilesystem.GetPackage()
	genericspkg := xgenerics.GetPackage()
	hashpkg := xhash.GetPackage()
	langpkg := xlang.GetPackage()
	p2ppkg := xp2p.GetPackage()
	randompkg := xrandom.GetPackage()
	socketpkg := xsocket.GetPackage()
	stringpkg := xstring.GetPackage()
	systempkg := xsystem.GetPackage()
	textpkg := xtext.GetPackage()
	threadpkg := xthread.GetPackage()
	timepkg := xtime.GetPackage()
	unittestpkg := xunittest.GetPackage()
	uuidpkg := xuuid.GetPackage()

	// The main (xbase) package
	mainpkg := denv.NewPackage("xbasepkgs")
	mainpkg.AddPackage(actorpkg)
	mainpkg.AddPackage(allocatorpkg)
	mainpkg.AddPackage(basepkg)
	mainpkg.AddPackage(binmappkg)
	mainpkg.AddPackage(cmdlinepkg)
	mainpkg.AddPackage(compresspkg)
	mainpkg.AddPackage(corepkg)
	mainpkg.AddPackage(cryptopkg)
	mainpkg.AddPackage(entrypkg)
	mainpkg.AddPackage(filesystempkg)
	mainpkg.AddPackage(genericspkg)
	mainpkg.AddPackage(hashpkg)
	mainpkg.AddPackage(langpkg)
	mainpkg.AddPackage(p2ppkg)
	mainpkg.AddPackage(randompkg)
	mainpkg.AddPackage(socketpkg)
	mainpkg.AddPackage(stringpkg)
	mainpkg.AddPackage(systempkg)
	mainpkg.AddPackage(textpkg)
	mainpkg.AddPackage(threadpkg)
	mainpkg.AddPackage(timepkg)
	mainpkg.AddPackage(unittestpkg)
	mainpkg.AddPackage(uuidpkg)

	// 'xbase' library
	mainlib := denv.SetupDefaultCppLibProject("xbasepkgs", "github.com\\jurgen-kluft\\xbasepkgs")
	mainlib.Dependencies = append(mainlib.Dependencies, unittestpkg.GetMainLib())

	mainpkg.AddMainLib(mainlib)
	return mainpkg
}
