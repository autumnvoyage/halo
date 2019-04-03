#!/bin/sh
#

zgCWD="${PWD}";
zgProject='halo';
zgProjDir="${zgCWD}/${zgProject}lib";
zgSrcDir="${zgProjDir}/src";
zgTestDir="${zgProjDir}/test";
zgOutName="${zgProject}";

# Cleanup
export zgOutName zgToCopy zgTestDir zgSrcDir zgProjDir zgProject zgCWD zgLibs;

# Execute the build
if [ "$1" = '3' ]; then
	# External program needs library for testing
	. util/lib/golang_lib.sh 1;
elif [ "$1" = '2' ]; then
	# External project needs library
	. util/lib/golang_lib.sh 0;
elif [ "$1" = '1' ]; then
	# Testing library
	. util/lib/golang_testlib.sh 0;
else
	# Just the library
	. util/lib/golang_lib.sh 0;
fi

# Cleanup
unset zgOutName zgToCopy zgTestDir zgSrcDir zgProjDir zgProject zgCWD zgLibs;
