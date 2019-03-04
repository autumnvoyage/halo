#!/bin/sh
#

if ! command -v npm 1>/dev/null; then
	echo 'npm installation not found. Exiting...';
	exit 3;
elif [ "${BUILD}" = '' ]; then
	echo 'Build location not specified. Exiting...';
	exit 4;
fi

CWD="${PWD}";
cd "${CWD}/frontend";
#npm run build:browser;
cd "${CWD}";
