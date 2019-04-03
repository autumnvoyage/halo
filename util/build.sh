#!/bin/sh
#

# Startup splash
echo '';
echo 'Z1 build system, for HALO';
echo 'Copyright (C) 2018-2019 HALO Contributors';
echo '';

if [ "$1" = '' ] || [ "$1" = '-h' ] || [ "$1" = '--help' ]; then
	echo 'Usage: util/build.sh [options] <target ...>';
	echo '';
	echo 'Options:';
	echo '    -r  Do a release build';
	echo '    -i  Install selected targets';
	echo '';
	echo 'Valid targets: arbiter bastion halolib browser cli electron clean';
	echo '               test';
	echo '';
	unset z1fn_helptext;
	exit 2;
fi
unset z1fn_helptext;

if [ "${Z1_BUILD}" = '' ]; then
	if [ "${BASH_SOURCE}" != '' ]; then
		Z1_BUILD="$(dirname "${BASH_SOURCE}")/../build";
	else
		Z1_BUILD="$(dirname "$(readlink -f "$0")")/../build";
	fi
	echo "Build directory: ./build";
else
	echo "Build directory: ${Z1_BUILD}";
fi

mkdir -p "${Z1_BUILD}/bin" "${Z1_BUILD}/lib" "${Z1_BUILD}/obj";
mkdir -p "${Z1_BUILD}/test/bin" "${Z1_BUILD}/test/lib";
zgLibDirs="${Z1_BUILD}/lib";
zgTestLibDirs="${Z1_BUILD}/test/lib";

zgTest=0;
zgRelease=0;
zgInstall=0;

export zgInstall zgRelease zgTest zgTestLibDirs zgLibDirs Z1_BUILD;

for _arg in $@; do
	[ "${_arg}" = 'test' ] && zgTest=1;
	[ "${_arg}" = '-r' ] && zgRelease=1;
	[ "${_arg}" = '-i' ] && zgInstall=1;
done
unset _arg;

for _arg in $@; do
	echo "Executing target ${_arg}..."
	if [ "${_arg}" = 'arbiter' ]; then
		. util/target/arbiter.sh "${zgTest}";
	elif [ "${_arg}" = 'bastion' ]; then
		. util/target/bastion.sh "${zgTest}";
	elif [ "${_arg}" = 'electron' ]; then
		. util/target/electron.sh "${zgTest}";
	elif [ "${_arg}" = 'browser' ]; then
		. util/target/browser.sh "${zgTest}";
	elif [ "${_arg}" = 'cli' ]; then
		. util/target/cli.sh "${zgTest}";
	elif [ "${_arg}" = 'halolib' ]; then
		. util/target/halolib.sh "${zgTest}";
	elif [ "${_arg}" = 'clean' ]; then
		. util/target/clean.sh;
	elif [ "${_arg}" != 'test' ]; then
		echo 'Must specify a valid build target as an argument. Exiting...';
		unset _arg;
		unset ZG_BUILD zgLibDirs zgTest zgRelease zgInstall;
		exit 3;
	fi
done
unset _arg;

[ "${zgInstall}" = '1' ] && . util/lib/install.sh;

unset Z1_BUILD zgLibDirs zgTest zgRelease zgInstall;
exit 0;
