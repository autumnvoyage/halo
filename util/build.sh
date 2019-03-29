#!/bin/sh
#

if [ "${BUILD}" = '' ]; then
	BUILD="$(dirname "$(readlink -f "$0")")/../build";
	echo "Build directory: ./build";
else
	echo "Build directory: ${BUILD}";
fi

echo '';
echo 'HALO build script';
echo 'Copyright (C) 2018-2019 HALO Contributors';
echo '';

if [ "$1" = '' ]; then
	echo 'Valid targets: arbiter bastion browser clean cli electron';
	echo '';
	unset BUILD;
	exit 3;
fi

test='0'

for arg in $@; do
	[ "${arg}" = 'test' ] && test='1';
done
unset arg;

for arg in $@; do
	echo "Executing target ${arg}..."
	if [ "${arg}" = 'arbiter' ]; then
		. util/target/arbiter.sh "$test";
	elif [ "${arg}" = 'bastion' ]; then
		. util/target/bastion.sh "$test";
	elif [ "${arg}" = 'electron' ]; then
		. util/target/electron.sh "$test";
	elif [ "${arg}" = 'browser' ]; then
		. util/target/browser.sh "$test";
	elif [ "${arg}" = 'cli' ]; then
		. util/target/cli.sh "$test";
	elif [ "${arg}" = 'clean' ]; then
		. util/target/clean.sh;
	elif [ "${arg}" != 'test' ]; then
		echo 'Must specify a valid build target as an argument. Exiting...';
		unset arg;
		unset BUILD CWD;
		exit 2;
	fi
done
unset arg;

unset BUILD;
