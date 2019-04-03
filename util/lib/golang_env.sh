#!/bin/sh
#

if ! command -v go 1>/dev/null; then
	echo 'Go installation not found. Exiting...';
	exit 4;
elif [ "${Z1_BUILD}" = '' ]; then
	echo 'Build location not specified. Exiting...';
	exit 5;
fi

zgGocc='go tool compile';
zgGoflags="-D ${zgSrcDir}";
if [ "$1" = '1' ]; then
	for _arg in ${zgLibDirs}; do
		zgGoflags="${zgGoflags} -I ${_arg}";
	done
else
	for _arg in ${zgTestLibDirs}; do
		zgGoflags="${zgGoflags} -I ${_arg}";
	done
fi
unset _arg;
zgGolink='go tool link';
zgGolinkflags="-d -L ${Z1_BUILD}/lib";

zgGoflags_L='-pack';

zgGoflags_R='-B -N -race';
zgGolinkflags_R='-s -w';

zgGoflags_D='';
zgGolinkflags_D='-race';

export zgGocc zgGoflags zgGolink zgGolinkflags zgGoflags_L zgGolinkflags_R \
	zgGoflags_R zgGolinkflags_D zgGoflags_D;
