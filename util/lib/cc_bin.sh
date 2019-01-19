#!/bin/sh
#

if [ "$CC" = '' ]; then
	if command -v clang 1>/dev/null; then
		CC=clang
	elif command -v gcc 1>/dev/null; then
		CC=gcc
	fi
fi

if [ "$LD" = '' ]; then
	LD="${CC}"
fi

mkdir -p "${BUILD}/bin" "${BUILD}/lib" "${BUILD}/obj";

OFILES=''
for file in ${CFILES}; do
	bfile="$(basename "${file}")"
	"${CC}" ${CFLAGS} -c -o "${BUILD}/obj/${bfile}.o" ${file};
	OFILES+=" ${BUILD}/obj/${bfile}.o";
	unset bfile;
done
unset file;

"${LD}" -o "${BUILD}/bin/${OUTNAME}" ${LIBDIRS} ${LDFLAGS} ${OFILES} ${LIBS};
unset OFILES;

for file in ${COPYLIB}; do
	cp "${file}" "${BUILD}/lib/";
done
unset file;
