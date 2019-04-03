#!/bin/sh
#

echo 'Compiling dependencies...';

if [ "$1" = '1' ]; then
	# Testing
	. util/target/halolib.sh 3;
else
	. util/target/halolib.sh 2;
fi

zgLibs='halo';

zgCWD="${PWD}";
zgProject='arbiter';
zgProjDir="${zgCWD}/${zgProject}";
zgSrcDir="${zgProjDir}/src";
zgTestDir="${zgProjDir}/test";
zgToCopy="${zgProjDir}/lib/halo-${zgProject}.service";
zgOutName="halo-${zgProject}";

# Cleanup
export zgOutName zgToCopy zgTestDir zgSrcDir zgProjDir zgProject zgCWD zgLibs;

# Execute the build
if [ "$1" = '1' ]; then
	echo 'Compiling test binary...';
	. util/lib/golang_testbin.sh;
else
	echo 'Compiling binary...';
	. util/lib/golang_bin.sh;
fi

# Cleanup
unset zgOutName zgToCopy zgTestDir zgSrcDir zgProjDir zgProject zgCWD zgLibs;
