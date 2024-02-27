#!/bin/bash

rm -rf assets/bin

mkdir -p assets/bin/texit assets/bin/texit-activities assets/bin/texit-discord assets/bin/texit-discord-callback

cp ../../dist/api_linux_arm64/texit assets/bin/texit/bootstrap
cp ../../dist/texit-lambda-sfn-activities_linux_arm64/bootstrap assets/bin/texit-activities/bootstrap
cp ../../dist/texit-discord-lambda_linux_arm64/bootstrap assets/bin/texit-discord/bootstrap
cp ../../dist/texit-discord-callback-lambda_linux_arm64/bootstrap assets/bin/texit-discord-callback/bootstrap

echo "Assets copied to assets/bin"
