#!/bin/sh
#

if [ "${BUILD}" = '' ]; then
	if [ "${BASH_SOURCE}" != '' ]; then
		BUILD="$(dirname "${BASH_SOURCE}")/../build";
	else
		BUILD="$(dirname "$(readlink -f "$0")")/../build";
	fi
	echo "Build directory: ${BUILD}"
fi

if [ "$1" = '' ]; then
	echo '';
	echo 'HALO build script';
	echo 'Copyright (C) 2018-2019 HALO Contributors';
	echo 'Valid targets: arbiter bastion browser clean cli electron';
	echo '';
	exit 3;
fi

for arg in $@; do
	if [ "${arg}" = 'arbiter' ]; then
		. util/target/arbiter.sh;
	elif [ "${arg}" = 'bastion' ]; then
		. util/target/bastion.sh;
	elif [ "${arg}" = 'electron' ]; then
		. util/target/electron.sh;
	elif [ "${arg}" = 'browser' ]; then
		. util/target/browser.sh;
	elif [ "${arg}" = 'cli' ]; then
		. util/target/cli.sh;
	elif [ "${arg}" = 'clean' ]; then
		. util/target/clean.sh;
	else
		echo 'Must specify a valid build target as an argument. Exiting...';
		exit 2;
	fi
done
unset arg;

unset BUILD;
