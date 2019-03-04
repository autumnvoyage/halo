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
GOFILES='bastion/src/main.go';
COPYLIB='bastion/lib/halo-bastion.service';
OUTNAME='halo-bastion';

# Execute the build
. util/lib/golang_bin.sh;
