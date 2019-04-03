#!/bin/sh
#

. util/lib/golang_env.sh 0;

if [ "$1" = '1' ]; then
	# External program is testing
	_outfile="${Z1_BUILD}/test/lib/${zgOutName}.a";
else
	_outfile="${Z1_BUILD}/lib/${zgOutName}.a";
fi

if [ "${zgRelease}" = '1' ]; then
	${zgGocc} ${zgGoflags} ${zgGoflags_L} ${zgGoflags_R} -o \
	"${_outfile}" $(find "${zgSrcDir}" -type f -name '*.go')
else
	${zgGocc} ${zgGoflags} ${zgGoflags_L} ${zgGoflags_D} -o \
	"${_outfile}" $(find "${zgSrcDir}" -type f -name '*.go')
fi

if [ ! -f "${_outfile}" ]; then
	echo 'Build failed! Exiting...';
	unset _outfile;
	. util/lib/golang_unenv.sh;
	exit 4;
fi

unset _outfile;
. util/lib/golang_unenv.sh;
