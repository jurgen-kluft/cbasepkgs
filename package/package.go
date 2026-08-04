package cbasepkgs

import (
	c3dff "github.com/jurgen-kluft/c3dff/package"
	cactor "github.com/jurgen-kluft/cactor/package"
	callocator "github.com/jurgen-kluft/callocator/package"
	catomic "github.com/jurgen-kluft/catomic/package"
	cbase "github.com/jurgen-kluft/cbase/package"
	ccmdline "github.com/jurgen-kluft/ccmdline/package"
	ccompress "github.com/jurgen-kluft/ccompress/package"
	cconartist "github.com/jurgen-kluft/cconartist/package"
	ccore "github.com/jurgen-kluft/ccore/package"
	ccrypto "github.com/jurgen-kluft/ccrypto/package"
	cecs "github.com/jurgen-kluft/cecs/package"
	centry "github.com/jurgen-kluft/centry/package"
	cfile "github.com/jurgen-kluft/cfile/package"
	cfilesystem "github.com/jurgen-kluft/cfilesystem/package"
	cgenerics "github.com/jurgen-kluft/cgenerics/package"
	cgx2 "github.com/jurgen-kluft/cgx2/package"
	charon "github.com/jurgen-kluft/charon/package"
	chash "github.com/jurgen-kluft/chash/package"
	chydra "github.com/jurgen-kluft/chydra/package"
	cjobs "github.com/jurgen-kluft/cjobs/package"
	cjson "github.com/jurgen-kluft/cjson/package"
	ckalman "github.com/jurgen-kluft/ckalman/package"
	cmath "github.com/jurgen-kluft/cmath/package"
	cmmio "github.com/jurgen-kluft/cmmio/package"
	cmsg "github.com/jurgen-kluft/cmsg/package"
	cp2p "github.com/jurgen-kluft/cp2p/package"
	cpair "github.com/jurgen-kluft/cpair/package"
	crandom "github.com/jurgen-kluft/crandom/package"
	craylib "github.com/jurgen-kluft/craylib/package"
	csocket "github.com/jurgen-kluft/csocket/package"
	cstring "github.com/jurgen-kluft/cstring/package"
	csuperalloc "github.com/jurgen-kluft/csuperalloc/package"
	csystem "github.com/jurgen-kluft/csystem/package"
	ctext "github.com/jurgen-kluft/ctext/package"
	cthread "github.com/jurgen-kluft/cthread/package"
	ctime "github.com/jurgen-kluft/ctime/package"
	cunittest "github.com/jurgen-kluft/cunittest/package"
	cuuid "github.com/jurgen-kluft/cuuid/package"
	cwindow "github.com/jurgen-kluft/cwindow/package"
	"github.com/jurgen-kluft/gide/denv"
)

// GetPackage returns the package object of 'cbase'
func GetPackage() *denv.Package {

	// Dependencies
	actorpkg := cactor.GetPackage()
	allocatorpkg := callocator.GetPackage()
	atomicpkg := catomic.GetPackage()
	basepkg := cbase.GetPackage()
	cmdlinepkg := ccmdline.GetPackage()
	compresspkg := ccompress.GetPackage()
	conartistpkg := cconartist.GetPackage()
	corepkg := ccore.GetPackage()
	cryptopkg := ccrypto.GetPackage()
	ecspkg := cecs.GetPackage()
	entrypkg := centry.GetPackage()
	filepkg := cfile.GetPackage()
	filesystempkg := cfilesystem.GetPackage()
	charonpkg := charon.GetPackage()
	genericspkg := cgenerics.GetPackage()
	gx2pkg := cgx2.GetPackage()
	hydrapkg := chydra.GetPackage()
	hashpkg := chash.GetPackage()
	jobspkg := cjobs.GetPackage()
	jsonpkg := cjson.GetPackage()
	kalmanpkg := ckalman.GetPackage()
	mathpkg := cmath.GetPackage()
	mmiopkg := cmmio.GetPackage()
	msgpkg := cmsg.GetPackage()
	pairpkg := cpair.GetPackage()
	p2ppkg := cp2p.GetPackage()
	randompkg := crandom.GetPackage()
	raylibpkg := craylib.GetPackage()
	socketpkg := csocket.GetPackage()
	stringpkg := cstring.GetPackage()
	superallocpkg := csuperalloc.GetPackage()
	systempkg := csystem.GetPackage()
	textpkg := ctext.GetPackage()
	threadpkg := cthread.GetPackage()
	timepkg := ctime.GetPackage()
	unittestpkg := cunittest.GetPackage()
	uuidpkg := cuuid.GetPackage()
	windowpkg := cwindow.GetPackage()
	c3dffpkg := c3dff.GetPackage()

	// The main (cbasepkgs) package
	mainpkg := denv.NewPackage("github.com\\jurgen-kluft", "cbasepkgs")
	mainpkg.AddPackage(actorpkg)
	mainpkg.AddPackage(allocatorpkg)
	mainpkg.AddPackage(atomicpkg)
	mainpkg.AddPackage(basepkg)
	mainpkg.AddPackage(cmdlinepkg)
	mainpkg.AddPackage(compresspkg)
	mainpkg.AddPackage(conartistpkg)
	mainpkg.AddPackage(corepkg)
	mainpkg.AddPackage(cryptopkg)
	mainpkg.AddPackage(ecspkg)
	mainpkg.AddPackage(entrypkg)
	mainpkg.AddPackage(filepkg)
	mainpkg.AddPackage(filesystempkg)
	mainpkg.AddPackage(charonpkg)
	mainpkg.AddPackage(genericspkg)
	mainpkg.AddPackage(gx2pkg)
	mainpkg.AddPackage(hydrapkg)
	mainpkg.AddPackage(hashpkg)
	mainpkg.AddPackage(jobspkg)
	mainpkg.AddPackage(jsonpkg)
	mainpkg.AddPackage(kalmanpkg)
	mainpkg.AddPackage(mathpkg)
	mainpkg.AddPackage(mmiopkg)
	mainpkg.AddPackage(msgpkg)
	mainpkg.AddPackage(pairpkg)
	mainpkg.AddPackage(p2ppkg)
	mainpkg.AddPackage(randompkg)
	mainpkg.AddPackage(raylibpkg)
	mainpkg.AddPackage(socketpkg)
	mainpkg.AddPackage(stringpkg)
	mainpkg.AddPackage(superallocpkg)
	mainpkg.AddPackage(systempkg)
	mainpkg.AddPackage(textpkg)
	mainpkg.AddPackage(threadpkg)
	mainpkg.AddPackage(timepkg)
	mainpkg.AddPackage(unittestpkg)
	mainpkg.AddPackage(uuidpkg)
	mainpkg.AddPackage(windowpkg)
	mainpkg.AddPackage(c3dffpkg)

	// 'cbasepkgs' library
	mainlib := denv.SetupCppLibProject(mainpkg, "cbasepkgs")
	mainlib.AddDependencies(unittestpkg.GetMainLib())

	mainpkg.AddMainLib(mainlib)

	return mainpkg
}
