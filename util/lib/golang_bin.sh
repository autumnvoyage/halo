#!/bin/sh
#

mkdir -p "${BUILD}/bin" "${BUILD}/lib";
cd "${MODULEDIR}";

go build ./...;
mv "${OUTNAME}" "${BUILD}/bin";

for file in ${COPYLIB}; do
	cp "${file}" "${BUILD}/lib/";
done
unset file;

cd "${CWD}";
