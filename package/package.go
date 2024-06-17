package cbasepkgs

import (
	c3dff "github.com/jurgen-kluft/c3dff/package"
	cactor "github.com/jurgen-kluft/cactor/package"
	callocator "github.com/jurgen-kluft/callocator/package"
	catomic "github.com/jurgen-kluft/catomic/package"
	cbase "github.com/jurgen-kluft/cbase/package"
	cbinmap "github.com/jurgen-kluft/cbinmap/package"
	ccmdline "github.com/jurgen-kluft/ccmdline/package"
	"github.com/jurgen-kluft/ccode/denv"
	ccompress "github.com/jurgen-kluft/ccompress/package"
	ccore "github.com/jurgen-kluft/ccore/package"
	ccrypto "github.com/jurgen-kluft/ccrypto/package"
	cecs "github.com/jurgen-kluft/cecs/package"
	centry "github.com/jurgen-kluft/centry/package"
	cfile "github.com/jurgen-kluft/cfile/package"
	cfilesystem "github.com/jurgen-kluft/cfilesystem/package"
	cgenerics "github.com/jurgen-kluft/cgenerics/package"
	charon "github.com/jurgen-kluft/charon/package"
	chash "github.com/jurgen-kluft/chash/package"
	chydra "github.com/jurgen-kluft/chydra/package"
	cjobs "github.com/jurgen-kluft/cjobs/package"
	cjson "github.com/jurgen-kluft/cjson/package"
	clang "github.com/jurgen-kluft/clang/package"
	cmath "github.com/jurgen-kluft/cmath/package"
	cp2p "github.com/jurgen-kluft/cp2p/package"
	cpair "github.com/jurgen-kluft/cpair/package"
	crandom "github.com/jurgen-kluft/crandom/package"
	csocket "github.com/jurgen-kluft/csocket/package"
	cstring "github.com/jurgen-kluft/cstring/package"
	csuperalloc "github.com/jurgen-kluft/csuperalloc/package"
	csystem "github.com/jurgen-kluft/csystem/package"
	ctext "github.com/jurgen-kluft/ctext/package"
	cthread "github.com/jurgen-kluft/cthread/package"
	ctime "github.com/jurgen-kluft/ctime/package"
	cunittest "github.com/jurgen-kluft/cunittest/package"
	cuuid "github.com/jurgen-kluft/cuuid/package"
	cvmem "github.com/jurgen-kluft/cvmem/package"
	cwindow "github.com/jurgen-kluft/cwindow/package"
)

// GetPackage returns the package object of 'cbase'
func GetPackage() *denv.Package {

	// Dependencies
	actorpkg := cactor.GetPackage()
	allocatorpkg := callocator.GetPackage()
	atomicpkg := catomic.GetPackage()
	basepkg := cbase.GetPackage()
	binmappkg := cbinmap.GetPackage()
	cmdlinepkg := ccmdline.GetPackage()
	compresspkg := ccompress.GetPackage()
	corepkg := ccore.GetPackage()
	cryptopkg := ccrypto.GetPackage()
	ecspkg := cecs.GetPackage()
	entrypkg := centry.GetPackage()
	filepkg := cfile.GetPackage()
	filesystempkg := cfilesystem.GetPackage()
	charonpkg := charon.GetPackage()
	genericspkg := cgenerics.GetPackage()
	hydrapkg := chydra.GetPackage()
	hashpkg := chash.GetPackage()
	jobspkg := cjobs.GetPackage()
	jsonpkg := cjson.GetPackage()
	langpkg := clang.GetPackage()
	mathpkg := cmath.GetPackage()
	pairpkg := cpair.GetPackage()
	p2ppkg := cp2p.GetPackage()
	randompkg := crandom.GetPackage()
	socketpkg := csocket.GetPackage()
	stringpkg := cstring.GetPackage()
	superallocpkg := csuperalloc.GetPackage()
	systempkg := csystem.GetPackage()
	textpkg := ctext.GetPackage()
	threadpkg := cthread.GetPackage()
	timepkg := ctime.GetPackage()
	unittestpkg := cunittest.GetPackage()
	uuidpkg := cuuid.GetPackage()
	vmempkg := cvmem.GetPackage()
	windowpkg := cwindow.GetPackage()
	p3dffpkg := c3dff.GetPackage()

	// The main (cbasepkgs) package
	mainpkg := denv.NewPackage("cbasepkgs")
	mainpkg.AddPackage(actorpkg)
	mainpkg.AddPackage(allocatorpkg)
	mainpkg.AddPackage(atomicpkg)
	mainpkg.AddPackage(basepkg)
	mainpkg.AddPackage(binmappkg)
	mainpkg.AddPackage(cmdlinepkg)
	mainpkg.AddPackage(compresspkg)
	mainpkg.AddPackage(corepkg)
	mainpkg.AddPackage(cryptopkg)
	mainpkg.AddPackage(ecspkg)
	mainpkg.AddPackage(entrypkg)
	mainpkg.AddPackage(filepkg)
	mainpkg.AddPackage(filesystempkg)
	mainpkg.AddPackage(charonpkg)
	mainpkg.AddPackage(genericspkg)
	mainpkg.AddPackage(hydrapkg)
	mainpkg.AddPackage(hashpkg)
	mainpkg.AddPackage(jobspkg)
	mainpkg.AddPackage(jsonpkg)
	mainpkg.AddPackage(langpkg)
	mainpkg.AddPackage(mathpkg)
	mainpkg.AddPackage(pairpkg)
	mainpkg.AddPackage(p2ppkg)
	mainpkg.AddPackage(randompkg)
	mainpkg.AddPackage(socketpkg)
	mainpkg.AddPackage(stringpkg)
	mainpkg.AddPackage(superallocpkg)
	mainpkg.AddPackage(systempkg)
	mainpkg.AddPackage(textpkg)
	mainpkg.AddPackage(threadpkg)
	mainpkg.AddPackage(timepkg)
	mainpkg.AddPackage(unittestpkg)
	mainpkg.AddPackage(uuidpkg)
	mainpkg.AddPackage(vmempkg)
	mainpkg.AddPackage(windowpkg)
	mainpkg.AddPackage(p3dffpkg)

	// 'cbasepkgs' library
	mainlib := denv.SetupDefaultCppLibProject("cbasepkgs", "github.com\\jurgen-kluft\\cbasepkgs")
	mainlib.Dependencies = append(mainlib.Dependencies, unittestpkg.GetMainLib())

	mainpkg.AddMainLib(mainlib)
	return mainpkg
}
