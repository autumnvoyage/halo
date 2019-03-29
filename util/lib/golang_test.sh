#!/bin/sh
#

cd "${MODULEDIR}";

go test ./...;

cd "${CWD}";
