package xbasepkgs

import (
	"github.com/jurgen-kluft/ccode/denv"
	xactor "github.com/jurgen-kluft/xactor/package"
	xallocator "github.com/jurgen-kluft/xallocator/package"
	xbase "github.com/jurgen-kluft/xbase/package"
	xbinmap "github.com/jurgen-kluft/xbinmap/package"
	xcmdline "github.com/jurgen-kluft/xcmdline/package"
	xcompress "github.com/jurgen-kluft/xcompress/package"
	xcore "github.com/jurgen-kluft/xcore/package"
	xcrypto "github.com/jurgen-kluft/xcrypto/package"
	xecs "github.com/jurgen-kluft/xecs/package"
	xentry "github.com/jurgen-kluft/xentry/package"
	xfilesystem "github.com/jurgen-kluft/xfilesystem/package"
	xgenerics "github.com/jurgen-kluft/xgenerics/package"
	xhash "github.com/jurgen-kluft/xhash/package"
	xlang "github.com/jurgen-kluft/xlang/package"
	xp2p "github.com/jurgen-kluft/xp2p/package"
	xrandom "github.com/jurgen-kluft/xrandom/package"
	xsocket "github.com/jurgen-kluft/xsocket/package"
	xstring "github.com/jurgen-kluft/xstring/package"
	xsystem "github.com/jurgen-kluft/xsystem/package"
	xtext "github.com/jurgen-kluft/xtext/package"
	xthread "github.com/jurgen-kluft/xthread/package"
	xtime "github.com/jurgen-kluft/xtime/package"
	xunittest "github.com/jurgen-kluft/xunittest/package"
	xuuid "github.com/jurgen-kluft/xuuid/package"
	xvmem "github.com/jurgen-kluft/xvmem/package"
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
	ecspkg := xecs.GetPackage()
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
	vmempkg := xvmem.GetPackage()

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
	mainpkg.AddPackage(ecspkg)
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
	mainpkg.AddPackage(vmempkg)

	// 'xbase' library
	mainlib := denv.SetupDefaultCppLibProject("xbasepkgs", "github.com\\jurgen-kluft\\xbasepkgs")
	mainlib.Dependencies = append(mainlib.Dependencies, unittestpkg.GetMainLib())

	mainpkg.AddMainLib(mainlib)
	return mainpkg
}
