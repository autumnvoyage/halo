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

for arg in $@; do
	if [ "${arg}" = 'arbiter' ]; then
		source util/target/arbiter.sh;
	elif [ "${arg}" = 'bastion' ]; then
		source util/target/bastion.sh;
	elif [ "${arg}" = 'electron' ]; then
		source util/target/electron.sh;
	elif [ "${arg}" = 'browser' ]; then
		source util/target/browser.sh;
	elif [ "${arg}" = 'cli' ]; then
		source util/target/cli.sh;
	elif [ "${arg}" = 'clean' ]; then
		source util/target/clean.sh;
	else
		echo 'Must specify a valid build target as an argument. Exiting...';
		exit 2;
	fi
done
unset arg;

unset BUILD;
