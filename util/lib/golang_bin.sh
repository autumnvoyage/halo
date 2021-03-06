#!/bin/sh
#

. util/lib/golang_env.sh 0;

_outfile="${Z1_BUILD}/bin/${zgOutName}";

if [ "${zgRelease}" = '1' ]; then
	_srcfiles=$(find "${zgSrcDir}" -type f -name '*.go');
	echo "Source files: ${_srcfiles}";
	for _file in ${_srcfiles}; do
		_out="${Z1_BUILD}/obj/$(basename ${_file}).o";
		${zgGocc} ${zgGoflags} ${zgGoflags_R} -o "${_out}" "${_file}";
	done
	unset _file _out _srcfiles;
	if [ "${zgLibs}" = '' ]; then
		_libs='';
	else
		_libs="${zgLibs}";
	fi
	${zgGolink} ${zgGolinkflags} ${zgGolinkflags_R} -o "${_outfile}" \
	            $(find "${Z1_BUILD}/obj" -type f -name '*.go.o') ${_libs};
	unset _libs;
else
	_srcfiles=$(find "${zgSrcDir}" -type f -name '*.go');
	echo "Source files: ${_srcfiles}";
	for _file in ${_srcfiles}; do
		_out="${Z1_BUILD}/obj/$(basename ${_file}).o";
		${zgGocc} ${zgGoflags} ${zgGoflags_D} -o "${_out}" "${_file}";
	done
	unset _file _out _srcfiles;
	if [ "${zgLibs}" = '' ]; then
		_libs='';
	else
		_libs="${zgLibs}";
	fi
	${zgGolink} ${zgGolinkflags} ${zgGolinkflags_D} -o "${_outfile}" \
	            $(find "${Z1_BUILD}/obj" -type f -name '*.go.o') ${_libs};
	unset _libs;
fi

rm -rf "${Z1_BUILD}/obj";
mkdir "${Z1_BUILD}/obj";

if [ ! -f "${_outfile}" ]; then
	echo 'Build failed! Exiting...';
	unset _outfile;
	. util/lib/golang_unenv.sh;
	exit 4;
fi

unset _outfile;
. util/lib/golang_unenv.sh;
