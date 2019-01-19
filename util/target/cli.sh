#!/bin/sh
#

if ! command -v go 1>/dev/null; then
	echo 'Go installation not found. Exiting...';
	exit 3;
elif [ "${BUILD}" = '' ]; then
	echo 'Build location not specified. Exiting...';
	exit 4;
fi

CFLAGS='-fPIC';
LDFLAGS='-pie';
CFILES='cli/src/main.c';
LIBS='';
OUTNAME='halo-cli';

# Execute the build
source util/lib/cc_bin.sh;
