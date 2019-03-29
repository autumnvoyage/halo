#!/bin/sh
#

if ! command -v go 1>/dev/null; then
	echo 'Go installation not found. Exiting...';
	exit 3;
elif [ "${BUILD}" = '' ]; then
	echo 'Build location not specified. Exiting...';
	exit 4;
fi

CWD="${PWD}";
PROJECT='arbiter';
PROJDIR="${PWD}/${PROJECT}";
MODULEDIR="${PROJDIR}/src";
COPYLIB="${PROJDIR}/lib/halo-${PROJECT}.service";
OUTNAME="halo-${PROJECT}";

# Execute the build
. util/lib/golang_bin.sh;
[ "$1" = '1' ] && . util/lib/golang_test.sh;

unset CWD PROJECT PROJDIR MODULEDIR COPYLIB OUTNAME;
