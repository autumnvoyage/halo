#!/bin/sh
#

if [ "${BUILD}" = '' ]; then
	echo 'Build location not specified. Exiting...';
	exit 4;
fi

rm -rf "${BUILD}";
