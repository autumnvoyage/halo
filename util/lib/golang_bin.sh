#!/bin/sh
#

mkdir -p "${BUILD}/bin" "${BUILD}/lib";

for dep in ${GODEPS}; do
	go get ${dep};
	go install ${dep};
done
unset dep;

go build -v -o "${BUILD}/bin/${OUTNAME}" ${GOFILES};

for file in ${COPYLIB}; do
	cp "${file}" "${BUILD}/lib/";
done
unset file;
