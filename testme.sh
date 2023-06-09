#!/bin/bash

pushd 0.27 >/dev/null 2>&1 ; go build ; popd >/dev/null 2>&1 
pushd 0.29 >/dev/null 2>&1 ; go build ; popd >/dev/null 2>&1

export AWS_CONFIG_FILE=$(pwd)/awsconfig

PROFILES="standalone usesession"
BKT="s3://duanetestfoo"
FILE=test.txt
SDKVER="v1 v2"

for prof in $PROFILES; do
	for ver in $SDKVER; do
		for f in 0.27 0.29; do
			BKTVER="$BKT?awssdk=$ver"
			echo "===== testing profile $prof gocloud-version $f awssdkver $ver"
			AWS_PROFILE=$prof ./$f/gocloud-test $BKTVER $FILE
			echo ""
		done
	done
done
