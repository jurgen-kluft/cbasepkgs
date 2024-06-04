#!/usr/bin/env bash

yell() { echo "$0: $*" >&2; }
die() { yell "$*"; exit 111; }
try() { "$@" || cd cbasepkgs || yell "cannot $*"; }
verify() { "$@" || cd cbasepkgs || die "cannot $*"; }

cd ..
try git clone https://github.com/jurgen-kluft/cbase.git
try git clone https://github.com/jurgen-kluft/cactor.git
try git clone https://github.com/jurgen-kluft/callocator.git
try git clone https://github.com/jurgen-kluft/csuperalloc.git
try git clone https://github.com/jurgen-kluft/catomic.git
try git clone https://github.com/jurgen-kluft/cbasepkgs.git
try git clone https://github.com/jurgen-kluft/cbinmap.git
try git clone https://github.com/jurgen-kluft/ccmdline.git
try git clone https://github.com/jurgen-kluft/ccompress.git
try git clone https://github.com/jurgen-kluft/ccore.git
try git clone https://github.com/jurgen-kluft/ccrypto.git
try git clone https://github.com/jurgen-kluft/cecs.git
try git clone https://github.com/jurgen-kluft/centry.git
try git clone https://github.com/jurgen-kluft/cfibers.git
try git clone https://github.com/jurgen-kluft/cfile.git
try git clone https://github.com/jurgen-kluft/cfilesystem.git
try git clone https://github.com/jurgen-kluft/cgenerics.git
try git clone https://github.com/jurgen-kluft/chash.git
try git clone https://github.com/jurgen-kluft/cjson.git
try git clone https://github.com/jurgen-kluft/crandom.git
try git clone https://github.com/jurgen-kluft/csocket.git
try git clone https://github.com/jurgen-kluft/cstring.git
try git clone https://github.com/jurgen-kluft/csystem.git
try git clone https://github.com/jurgen-kluft/ctext.git
try git clone https://github.com/jurgen-kluft/cthread.git
try git clone https://github.com/jurgen-kluft/ctime.git
try git clone https://github.com/jurgen-kluft/cuuid.git
try git clone https://github.com/jurgen-kluft/cpair.git
try git clone https://github.com/jurgen-kluft/c3dff.git
try git clone https://github.com/jurgen-kluft/cmath.git
try git clone https://github.com/jurgen-kluft/cgfx.git
try git clone https://github.com/jurgen-kluft/cwindow.git
try git clone https://github.com/jurgen-kluft/cunittest.git
try git clone https://github.com/jurgen-kluft/chistogram.git
try git clone https://github.com/jurgen-kluft/cbenchmark.git
try git clone https://github.com/jurgen-kluft/ccode.git
try git clone https://github.com/jurgen-kluft/cfort.git
try git clone https://github.com/jurgen-kluft/cvmem.git
try git clone https://github.com/jurgen-kluft/cgamelogic.git
try git clone https://github.com/jurgen-kluft/cjobs.git
try git clone https://github.com/jurgen-kluft/cvulkan.git
try git clone https://github.com/jurgen-kluft/charon.git
try git clone https://github.com/jurgen-kluft/chydra.git
cd cbasepkgs
