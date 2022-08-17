package cbasepkgs

import (
	cactor "github.com/jurgen-kluft/cactor/package"
	callocator "github.com/jurgen-kluft/callocator/package"
	cbase "github.com/jurgen-kluft/cbase/package"
	cbinmap "github.com/jurgen-kluft/cbinmap/package"
	ccmdline "github.com/jurgen-kluft/ccmdline/package"
	"github.com/jurgen-kluft/ccode/denv"
	ccompress "github.com/jurgen-kluft/ccompress/package"
	ccore "github.com/jurgen-kluft/ccore/package"
	ccrypto "github.com/jurgen-kluft/ccrypto/package"
	cecs "github.com/jurgen-kluft/cecs/package"
	centry "github.com/jurgen-kluft/centry/package"
	cfilesystem "github.com/jurgen-kluft/cfilesystem/package"
	cgenerics "github.com/jurgen-kluft/cgenerics/package"
	chash "github.com/jurgen-kluft/chash/package"
	clang "github.com/jurgen-kluft/clang/package"
	cp2p "github.com/jurgen-kluft/cp2p/package"
	crandom "github.com/jurgen-kluft/crandom/package"
	csocket "github.com/jurgen-kluft/csocket/package"
	cstring "github.com/jurgen-kluft/cstring/package"
	csystem "github.com/jurgen-kluft/csystem/package"
	ctext "github.com/jurgen-kluft/ctext/package"
	cthread "github.com/jurgen-kluft/cthread/package"
	ctime "github.com/jurgen-kluft/ctime/package"
	cunittest "github.com/jurgen-kluft/cunittest/package"
	cuuid "github.com/jurgen-kluft/cuuid/package"
	cvmem "github.com/jurgen-kluft/cvmem/package"
)

// GetPackage returns the package object of 'cbase'
func GetPackage() *denv.Package {

	// Dependencies
	actorpkg := cactor.GetPackage()
	allocatorpkg := callocator.GetPackage()
	basepkg := cbase.GetPackage()
	binmappkg := cbinmap.GetPackage()
	cmdlinepkg := ccmdline.GetPackage()
	compresspkg := ccompress.GetPackage()
	corepkg := ccore.GetPackage()
	cryptopkg := ccrypto.GetPackage()
	ecspkg := cecs.GetPackage()
	entrypkg := centry.GetPackage()
	filesystempkg := cfilesystem.GetPackage()
	genericspkg := cgenerics.GetPackage()
	hashpkg := chash.GetPackage()
	langpkg := clang.GetPackage()
	p2ppkg := cp2p.GetPackage()
	randompkg := crandom.GetPackage()
	socketpkg := csocket.GetPackage()
	stringpkg := cstring.GetPackage()
	systempkg := csystem.GetPackage()
	textpkg := ctext.GetPackage()
	threadpkg := cthread.GetPackage()
	timepkg := ctime.GetPackage()
	unittestpkg := cunittest.GetPackage()
	uuidpkg := cuuid.GetPackage()
	vmempkg := cvmem.GetPackage()

	// The main (cbasepkgs) package
	mainpkg := denv.NewPackage("cbasepkgs")
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

	// 'cbasepkgs' library
	mainlib := denv.SetupDefaultCppLibProject("cbasepkgs", "github.com\\jurgen-kluft\\cbasepkgs")
	mainlib.Dependencies = append(mainlib.Dependencies, unittestpkg.GetMainLib())

	mainpkg.AddMainLib(mainlib)
	return mainpkg
}
