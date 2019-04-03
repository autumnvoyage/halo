#!/bin/sh
#

. util/lib/golang_env.sh 1;

_outfile="${Z1_BUILD}/test/bin/${zgOutName}";

# Testing the library itself
if [ "${zgRelease}" = '1' ]; then
	${zgGocc} ${zgGoflags} ${zgGoflags_R} -o "${_outfile}" \
	          $(find "${zgSrcDir}" -type f -name '*.go') \
	          $(find "${zgTestDir}" -type f -name '*.go')
else
	${zgGocc} ${zgGoflags} ${zgGoflags_D} -o "${_outfile}" \
	          $(find "${zgSrcDir}" -type f -name '*.go') \
	          $(find "${zgTestDir}" -type f -name '*.go')
fi
	if [ ! -f "${_outfile}" ]; then
	echo 'Build failed! Exiting...';
	unset _outfile;
	. util/lib/golang_unenv.sh;
	exit 4;
fi

unset _outfile;
. util/lib/golang_unenv.sh;
