#!/bin/sh
#

if ! command -v go 1>/dev/null; then
	echo 'Go installation not found. Exiting...';
	exit 3;
elif [ "${BUILD}" = '' ]; then
	echo 'Build location not specified. Exiting...';
	exit 4;
fi

GODEPS='github.com/graphql-go/graphql';
GOFILES='arbiter/src/main.go';
COPYLIB='arbiter/lib/halo-arbiter.service';
OUTNAME='halo-arbiter';

# Execute the build
. util/lib/golang_bin.sh;
