#!/usr/bin/env sh
# This script will bundle all the API files into a single file.
# This is useful for serving the API as a single file.
set -e

SERVICES=(
    "users"
    "products"
    "orders"
    # "payments"
)

for SERVICE_NAME in ${SERVICES[@]}; do
    swagger-cli bundle api/${SERVICE_NAME}/main.yml --type yaml > api/${SERVICE_NAME}/main.all.yml
    REVISION_DATE=$(date +%Y%m%d)
    if [[ -z "$CI_COMMIT_SHORT_SHA" ]]; then
        REVISION_HASH=$(git rev-parse --short HEAD)
    else
        REVISION_HASH=$CI_COMMIT_SHORT_SHA
    fi
    REVISION=$REVISION_HASH

    # awk and sed have different behavior between Mac and Linux, so we
    # use javascript
    if [ "$(uname)" == "Darwin" ]; then
        gsed --version > /dev/null || echo "No gnu-sed installed"
        SEDCMD=gsed       
    elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
        SEDCMD=sed       
    fi
    $SEDCMD "s/\$REVISION/$REVISION/" api/${SERVICE_NAME}/main.all.yml > docs/${SERVICE_NAME}/api.yml
    swagger-cli validate docs/${SERVICE_NAME}/api.yml
done